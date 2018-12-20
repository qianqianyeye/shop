package smobile

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"git.jiaxianghudong.com/go/logs"
	"net/http"
	"shop/admin/utils"
	."shop/api/handler/model"
	."shop/api/constant"
	. "shop/api/utils"
	"shop/api/mysql"
	"strings"
)

type RspShop struct {

}
type RspShopTypeList struct {
	Code         int        `json:"code"`
	Msg          string     `json:"msg"`
	Data         interface{} `json:"data"`
	PageModel    PageModel  `json:"page"`
}

func GetShopType(c *gin.Context)  {
	rsp := RspShopTypeList{Code: RC_OK, Msg: M(RC_OK)}
	var req ReqShopList
	if err := c.ShouldBindWith(&req, binding.Query); err == nil {
		page := GetPageInfo(req.Page_size, req.Current)
		var ShopTypeList []ShopType
		var count Count
		db.SqlDB.Order("update_at desc").Offset(page.OffSet).Limit(page.PageSize).Preload("Image","img_type=3").Find(&ShopTypeList)
		db.SqlDB.Table("shop_type").Select("count(*) count").Scan(&count)
		rsp.Data = ShopTypeList
		rsp.PageModel.Current = page.Current
		rsp.PageModel.Total = count.Count
		rsp.PageModel.PageSize = page.PageSize
		c.JSON(http.StatusOK, rsp)
	} else {
		rsp = RspShopTypeList{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR),Data:nil}
		c.JSON(http.StatusOK,rsp)
	}
}
type RspShopList struct {
	Code      int             `json:"code"`
	Msg       string          `json:"msg"`
	Data     interface{} 	  `json:"data"`
	//ShopList  []ShopInfo      `json:"shop_list"`
	PageModel utils.PageModel `json:"page"`
}
type RspShopCommon struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
	//ShopInfo ShopInfo `json:"shop_info"`
}
type ReqShopList struct {
	Current   string `form:"current"  binding:"required"`
	Page_size string `form:"page_size"  binding:"required"`
	//FieldName string `form:"field_name" json:"field_name"`
	Condition string `form:"condition" json:"condition"`
	SortField string `form:"sort_field" json:"sort_field"`
	Sort      string `form:"sort" json:"sort"`
	TypeId      int     `form:"type_id" json:"type_id"`
	Tags      string  `form:"tags" json:"tags"`
}
type Count struct {
	Count int
}
func GetShopList(c *gin.Context)  {
	rsp := RspShopList{Code: RC_OK, Msg: M(RC_OK)}
	var req ReqShopList
	if err := c.ShouldBindWith(&req, binding.Query); err == nil {
		page := GetPageInfo(req.Page_size, req.Current)
		var ShopInfoList []ShopInfo
		var img []Image
		var count Count
		sdb:=db.SqlDB.Table("shop_info")
		cdb :=db.SqlDB.Table("shop_info")
		typeTagMap := make(map[int64]int64)
		if len(req.Tags)>0 {
			tagArr := strings.Split(req.Tags,",")
			var typeTags []TagType
			cond, vals, err := WhereBuild(map[string]interface{}{
				"tag_id in": tagArr,
			})
			if err!=nil {
				logs.Error(err)
			}
			err =db.SqlDB.Where(cond, vals...).Find(&typeTags).Error
			if err!=nil {
				logs.Error(err)
			}
			for _,v:=range typeTags {
				typeTagMap[v.ShopInfoId]=v.TagId
			}
			var shopId []int64
			for k,_:=range typeTagMap {
				shopId=append(shopId,k)
			}
			sdb=sdb.Where("id in (?)",shopId)
			cdb=cdb.Where("id in (?)",shopId)
		}

		if req.SortField=="" {
			req.SortField="update_at"
		}
		if req.Sort=="" {
			req.Sort=" desc"
		}
		if req.TypeId !=0 {
			sdb=sdb.Where("type_id=?",req.TypeId)
			cdb=cdb.Where("type_id=?",req.TypeId)
		}

		if  req.Condition != "" {
			conditionArr :=strings.Split(req.Condition," ")
			var parmCondition string
			var parmArr []string
			for _,v:=range conditionArr{
				parmCondition = "(shop_name like '%"+v+"%' or r_shop_name like '%"+v+"%' or shop_describe like '%"+v+"%' or r_shop_describe like '%"+v+"%')"
				parmArr=append(parmArr,parmCondition)
			}
			var whereSql string
			for  i,_:=range parmArr{
				whereSql+=parmArr[i]+" or "
			}
			whereSql=whereSql[:len(whereSql)-3]
			err=sdb.Order(req.SortField+" "+req.Sort).Offset(page.OffSet).Limit(page.PageSize).
				Select("id,shop_name,r_shop_name,market_price,discount_price,r_market_price,r_discount_price,shop_describe,r_shop_describe").
				Where(whereSql).Scan(&ShopInfoList).Error
			var ids []int
			for _,v:=range ShopInfoList{
				ids=append(ids, int(v.ID))
			}
			cond, vals, err :=WhereBuild(map[string]interface{}{
				"img_type":"1",
				"target_id in": ids,
			})
			if err !=nil {
				logs.Error(err)
			}
			db.SqlDB.Table("image").Select("img_url,target_id").Where(cond, vals...).Group("target_id").Scan(&img)

			for i,v:=range ShopInfoList{
				for _,j:=range img{
					if v.ID==j.TargetId {
						var imgs []Image
						imgs=append(imgs,j)
						ShopInfoList[i].Image=imgs
					}
				}
			}
			err=cdb.Select("count(*) count").Where(whereSql).Scan(&count).Error

			if err!=nil {
				rsp.Code = RC_SYS_ERR
				rsp.Msg = M(RC_SYS_ERR)
				rsp.Data=nil
				logs.Error(err)
				c.JSON(http.StatusOK, rsp)
			}

		} else {
			err=sdb.Order(req.SortField+" "+req.Sort).Offset(page.OffSet).Limit(page.PageSize).
				Select("id,shop_name,market_price,discount_price,r_shop_name,r_market_price,r_discount_price,shop_describe,r_shop_describe").Scan(&ShopInfoList).Error
			var ids []int
			for _,v:=range ShopInfoList{
				ids=append(ids, int(v.ID))
			}
			cond, vals, err :=WhereBuild(map[string]interface{}{
				"img_type":"1",
				"target_id in": ids,
			})
			if err !=nil {
				logs.Error(err)
			}
			db.SqlDB.Table("image").Select("img_url,target_id").Where(cond, vals...).Group("target_id").Scan(&img)

			for i,v:=range ShopInfoList{
				for _,j:=range img{
					if v.ID==j.TargetId {
						var imgs []Image
						imgs=append(imgs,j)
						ShopInfoList[i].Image=imgs
					}
				}
			}

		err=cdb.Select("count(*) count").Scan(&count).Error
			if err !=nil {
				rsp.Code = RC_SYS_ERR
				rsp.Msg = M(RC_SYS_ERR)
				rsp.Data=nil
				logs.Error(err)
				c.JSON(http.StatusOK, rsp)
			}
		}
		rsp.Data = ShopInfoList
		rsp.PageModel.Current = page.Current
		rsp.PageModel.Total = count.Count
		rsp.PageModel.PageSize = page.PageSize
		c.JSON(http.StatusOK, rsp)
	} else {
		rsp.Code = RC_PARM_ERR
		rsp.Msg = M(RC_PARM_ERR)
		rsp.Data=nil
		c.JSON(http.StatusOK, rsp)
	}
}

func GetShopDetails(c *gin.Context)  {
	id:=c.Param("id")
	rsp := RspShopCommon{Code:RC_OK,Msg:M(RC_OK)}
	if id!="" {
		var shopInfo ShopInfo
		err:=db.SqlDB.Preload("ShopStyle.Image", "img_type=2").Preload("ShopStyle").Preload("Image", "img_type=1").
			Preload("ShopType.Image","img_type=3").Preload("Tags").Preload("ShopType").Find(&shopInfo,"id=?",id).Error
		if err!=nil {
			rsp.Code=RC_SYS_ERR
			rsp.Msg=M(RC_SYS_ERR)
			rsp.Data=nil
			logs.Error(err)
		}
		rsp.Data=shopInfo
	}else {
		rsp.Code=RC_PARM_ERR
		rsp.Msg=M(RC_PARM_ERR)
	}
	c.JSON(http.StatusOK,rsp)
}

func GetShopSearch(C *gin.Context)  {

}

func GetHotList(c *gin.Context)  {
	var Hots []HotSearch
	db.SqlDB.Order("sort desc").Limit(10).Find(&Hots)
	c.JSON(http.StatusOK,gin.H{"code":RC_OK,"msg":M(RC_OK),"data":Hots})
}

