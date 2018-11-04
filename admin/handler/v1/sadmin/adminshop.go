package sadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	. "shop/admin/constant"
	. "shop/admin/handler/model"
	"shop/admin/mysql"
	"shop/admin/utils"
	. "shop/admin/utils"
	"time"
	"git.jiaxianghudong.com/go/logs"
)

type RspCommon struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type RspShopTypeList struct {
	Code         int        `json:"code"`
	Msg          string     `json:"msg"`
	ShopTypeList []ShopType `json:"shop_type_list"`
	PageModel    PageModel  `json:"page"`
}

func GetShopType(c *gin.Context) {
	rsp := RspShopTypeList{Code: RC_OK, Msg: M(RC_OK)}
	var req ReqShopList
	if err := c.ShouldBindWith(&req, binding.Query); err == nil {
		page := GetPageInfo(req.Page_size, req.Current)
		var ShopTypeList []ShopType
		var count Count
		db.SqlDB.Order("update_at desc").Offset(page.OffSet).Limit(page.PageSize).Preload("Image").Find(&ShopTypeList)
		db.SqlDB.Table("shop_type").Select("count(*) count").Scan(&count)
		rsp.ShopTypeList = ShopTypeList
		rsp.PageModel.Current = page.Current
		rsp.PageModel.Total = count.Count
		rsp.PageModel.PageSize = page.PageSize
		c.JSON(http.StatusOK, rsp)
	} else {
		rsp = RspShopTypeList{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
		c.JSON(http.StatusOK,rsp)
	}
}

func AddShopType(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	var req ShopType
	if err := c.ShouldBind(&req); err == nil {
		//tx:=db.SqlDB.Begin()
		req.UpdateAt=GetStringDateTime(time.Now())
		req.CreatedAt=GetStringDateTime(time.Now())
		for i,_ := range req.Image{
			req.Image[i].CreateAt=GetStringDateTime(time.Now())
		}
		db.SqlDB.Create(&req)
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

func DeleteShopType(c *gin.Context) {
	id := c.Param("id")
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	if id != "" {
		tx := db.SqlDB.Begin()
		//删除图片
		if err := tx.Where("target_id=? and img_type=3", id).Delete(Image{}).Error; err != nil {
			tx.Rollback()
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
		//删除类型
		if err := tx.Where("id=?", id).Delete(ShopType{}).Error; err != nil {
			tx.Rollback()
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
		tx.Commit()
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

func UpdateShopType(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	var req ShopType
	if err := c.ShouldBind(&req); err == nil {
		tx := db.SqlDB.Begin()
		//删除旧的图片连接
		if err := tx.Where("target_id=? and img_type=3", req.ID).Delete(Image{}).Error; err != nil {
			tx.Rollback()
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
			return
		}
		//插入新的图片
		for _, v := range req.Image {
			v.TargetId = req.ID
			v.ImgType = 3
			v.CreateAt=GetStringDateTime(time.Now())
			if err := tx.Create(&v).Error; err != nil {
				tx.Rollback()
				rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
				return
			}
		}
		req.UpdateAt = utils.GetStringDateTime(time.Now())
		//更新商品
		ShopTypeMap := utils.StructToJsonMap(req)
		delete(ShopTypeMap, "id")
		delete(ShopTypeMap, "image")
		delete(ShopTypeMap, "created_at")
		if err := tx.Table("shop_type").Where("id=?", req.ID).Update(ShopTypeMap).Error; err != nil {
			tx.Rollback()
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
			return
		}
		tx.Commit()
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

type RspHotList struct {
	Code         int        `json:"code"`
	Msg          string     `json:"msg"`
	HotSearchList[]HotSearch `json:"hot_search_list"`
	PageModel    PageModel  `json:"page"`
}

func GetHotKey(c *gin.Context) {
	rsp := RspHotList{Code: RC_OK, Msg: M(RC_OK)}
	var req ReqShopList
	if err := c.ShouldBindWith(&req, binding.Query); err == nil {
		page := GetPageInfo(req.Page_size, req.Current)
		var HotList []HotSearch
		var count Count
		db.SqlDB.Order("sort desc").Offset(page.OffSet).Limit(page.PageSize).Find(&HotList)
		db.SqlDB.Table("hot_search").Select("count(*) count").Scan(&count)
		rsp.HotSearchList = HotList
		rsp.PageModel.Current = page.Current
		rsp.PageModel.Total = count.Count
		rsp.PageModel.PageSize = page.PageSize
		c.JSON(http.StatusOK, rsp)
	} else {
		rsp = RspHotList{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
		c.JSON(http.StatusOK,rsp)
	}
}

func AddHotKey(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	var req HotSearch
	if err := c.ShouldBind(&req); err == nil {
		var err error
		req.CreateAt=GetStringDateTime(time.Now())
		req.UpdateAt=GetStringDateTime(time.Now())
		err=db.SqlDB.Create(&req).Error
		err =db.SqlDB.Model(&req).Update("sort", req.ID).Error
		if err !=nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
		c.JSON(http.StatusOK,rsp)
	}else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
		c.JSON(http.StatusOK,rsp)
	}
}

func DeleteHotKey(c *gin.Context) {
	id :=c.Param("id")
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	if id!="" {
		if err :=db.SqlDB.Where("id=?",id).Delete(HotSearch{}).Error;err!=nil{
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
	}else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
	}
	defer func() {
		c.JSON(http.StatusOK,rsp)
	}()
}

func UpdateHotKey(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	var req HotSearch
	if err := c.ShouldBind(&req); err == nil {
		//req.CreateAt=GetStringDateTime(time.Now())
		req.UpdateAt=GetStringDateTime(time.Now())
		hotMap := StructToJsonMap(req)
		var err error
		if req.Action !=0 {
			if req.Action == 1 {
				//上
				var topHot HotSearch
				//获取上一条记录
				db.SqlDB.Table("hot_search").Select("*").Where("sort>?",req.Sort).Order("sort asc").Limit(1).Scan(&topHot)

				//=db.SqlDB.Model(&req).Update("sort", req.ID)
				//tx.Model(&req).Update()
				if topHot.Sort!=0 {
					tx := db.SqlDB.Begin()
					err=tx.Table("hot_search").Where("id=?",req.ID).Update("sort",topHot.Sort).Error
					err=tx.Table("hot_search").Where("id=?",topHot.ID).Update("sort",req.Sort).Error
					tx.Commit()
				}
				rsp = RspCommon{Code: RC_OK, Msg: M(RC_OK)}
			}else if req.Action==2 {
				//下
				var downHot HotSearch
				//获取下一条记录
				db.SqlDB.Table("hot_search").Select("*").Where("sort<?",req.Sort).Order("sort desc").Limit(1).Scan(&downHot)
				if downHot.Sort!=0 {
					tx := db.SqlDB.Begin()
					err =tx.Table("hot_search").Where("id=?",req.ID).Update("sort",downHot.Sort).Error
					err =tx.Table("hot_search").Where("id=?",downHot.ID).Update("sort",req.Sort).Error
					tx.Commit()
				}
				rsp = RspCommon{Code: RC_OK, Msg: M(RC_OK)}
			}else {
				rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
			}
		}else {
			delete(hotMap, "id")
			delete(hotMap, "create_at")
			delete(hotMap, "sort")
			delete(hotMap, "action")
			err=db.SqlDB.Table("hot_search").Where("id=?",req.ID).Update(hotMap).Error
		}
		if err !=nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
	}else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
	}
	defer func() {
		c.JSON(http.StatusOK,rsp)
	}()
}

type ReqAddShopStyle struct {
	ShopStyle ShopStyle `json:"shop_style"`
	ImgStyle  []Image   `json:"img_style"`
}
type ReqAddShopInfo struct {
	ShopInfo ShopInfo `json:"shop_info"`
	ImgShop  []Image  `json:"img_shop"`
}
type ReqShopInfo struct {
	Shop       ReqAddShopInfo    `json:"shop"`
	ShopStyles []ReqAddShopStyle `json:"shop_styles"`
}
type RspShopList struct {
	Code      int             `json:"code"`
	Msg       string          `json:"msg"`
	ShopList  []ShopInfo      `json:"shop_list"`
	PageModel utils.PageModel `json:"page"`
}

type ReqShopList struct {
	Current   string `form:"current"  binding:"required"`
	Page_size string `form:"page_size"  binding:"required"`
	FieldName string `form:"field_name" json:"field_name"`
	Condition string `form:"condition" json:"condition"`
}
type Count struct {
	Count int
}

func GetShopList(c *gin.Context) {
	rsp := RspShopList{Code: RC_OK, Msg: M(RC_OK)}
	var req ReqShopList
	if err := c.ShouldBindWith(&req, binding.Query); err == nil {
		page := GetPageInfo(req.Page_size, req.Current)
		var err error
		var ShopInfoList []ShopInfo
		var count Count
		if req.FieldName != "" && req.Condition != "" {
			err=db.SqlDB.Order("update_at desc").Offset(page.OffSet).Limit(page.PageSize).Preload("ShopStyle.Image", "img_type=2").Preload("ShopStyle").Preload("Image", "img_type=1").
				Preload("ShopType").Find(&ShopInfoList, req.FieldName+" like ?", "%"+req.Condition+"%").Error
			err=db.SqlDB.Table("shop_info").Select("count(*) count").Where(req.FieldName+" like ?", "%"+req.Condition+"%").Scan(&count).Error
			if err!=nil {
				rsp.Code = RC_SYS_ERR
				rsp.Msg = M(RC_SYS_ERR)
				logs.Error(err)
				c.JSON(http.StatusOK, rsp)
			}
		} else {
			err=db.SqlDB.Order("update_at desc").Offset(page.OffSet).Limit(page.PageSize).Preload("ShopStyle.Image", "img_type=2").Preload("ShopStyle").Preload("Image", "img_type=1").
				Preload("ShopType").Find(&ShopInfoList).Error
			err=db.SqlDB.Table("shop_info").Select("count(*) count").Scan(&count).Error
			if err !=nil {
				rsp.Code = RC_SYS_ERR
				rsp.Msg = M(RC_SYS_ERR)
				logs.Error(err)
				c.JSON(http.StatusOK, rsp)
			}
		}
		rsp.ShopList = ShopInfoList
		rsp.PageModel.Current = page.Current
		rsp.PageModel.Total = count.Count
		rsp.PageModel.PageSize = page.PageSize
		c.JSON(http.StatusOK, rsp)
	} else {
		rsp.Code = RC_PARM_ERR
		rsp.Msg = M(RC_PARM_ERR)
		c.JSON(http.StatusOK, rsp)
	}
}

func AddShop(c *gin.Context) {
	g := Gin{c}
	var req ReqShopInfo
	if err := c.ShouldBind(&req); err == nil {
		tx := db.SqlDB.Begin()
		req.Shop.ShopInfo.CreateAt = utils.GetStringDateTime(time.Now())
		req.Shop.ShopInfo.UpdateAt = utils.GetStringDateTime(time.Now())
		shop := req.Shop.ShopInfo
		err := tx.Create(&shop).Error
		if err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		// 1商品 2款式
		//插入商品图片
		for i, _ := range req.Shop.ImgShop {
			req.Shop.ImgShop[i].TargetId = shop.ID
			req.Shop.ImgShop[i].CreateAt = utils.GetStringDateTime(time.Now())
			err = tx.Create(&req.Shop.ImgShop[i]).Error
			if err != nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
		}

		//插入款式跟款式图片
		for i, _ := range req.ShopStyles {
			req.ShopStyles[i].ShopStyle.ShopId = shop.ID
			style := req.ShopStyles[i].ShopStyle
			err = tx.Create(&style).Error
			if err != nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
			for j, _ := range req.ShopStyles[i].ImgStyle {
				req.ShopStyles[i].ImgStyle[j].CreateAt = utils.GetStringDateTime(time.Now())
				req.ShopStyles[i].ImgStyle[j].TargetId = style.ID
				err = tx.Create(&req.ShopStyles[i].ImgStyle[j]).Error
				if err != nil {
					tx.Rollback()
					g.FResponse(nil)
					return
				}
			}
		}
		tx.Commit()
		g.Response(nil)
	} else {
		g.PResponse(err)
	}
}

func DeleteShop(c *gin.Context) {
	g := Gin{c}
	id := c.Param("id")
	if id != "" {
		var style []ShopStyle
		db.SqlDB.Preload("Image","img_type=2").Find(&style, "shop_id=? ", id)
		var shopImages []Image
		var deleteImage []string
		db.SqlDB.Find(&shopImages, "target_id=?", id)
		tx := db.SqlDB.Begin()
		//删除商品
		if err := tx.Where("id=?", id).Delete(ShopInfo{}).Error; err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		//删除商品图片
		if err := tx.Where("target_id=? and img_type=1", id).Delete(Image{}).Error; err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		for _, v := range shopImages {
			deleteImage = append(deleteImage, RuntimeRootPath+v.ImgUrl)
		}
		//删除款式
		if err := tx.Where("shop_id=?", id).Delete(ShopStyle{}).Error; err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		//删除款式图片
		for _, v := range style {
			for _, img := range v.Image {
				deleteImage = append(deleteImage, RuntimeRootPath+img.ImgUrl)
				if err := tx.Delete(&img).Error; err != nil {
					tx.Rollback()
					g.FResponse(nil)
					return
				}
			}
		}
		tx.Commit()
		g.Response(nil)
		//删除磁盘存储图片
		go deleteImages(deleteImage)
	}
}

func UpdateShop(c *gin.Context) {
	g := Gin{c}
	var req ReqShopInfo
	if err := c.ShouldBind(&req); err == nil {
		tx := db.SqlDB.Begin()
		//删除旧的商品图片
		if err := tx.Where("target_id=? and img_type=1", req.Shop.ShopInfo.ID).Delete(Image{}).Error; err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		//添加新的商品图片
		for _, v := range req.Shop.ImgShop {
			v.TargetId = req.Shop.ShopInfo.ID
			v.CreateAt = utils.GetStringDateTime(time.Now())
			if err := tx.Create(&v).Error; err != nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
		}

		//删除旧的款式图片
		var oldStyle []ShopStyle
		db.SqlDB.Find(&oldStyle, "shop_id=?", req.Shop.ShopInfo.ID)
		for _, v := range oldStyle {
			if err := tx.Where("target_id=? and img_type=2", v.ID).Delete(Image{}).Error; err != nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
		}
		//删除旧的款式
		if err := tx.Where("shop_id=?", req.Shop.ShopInfo.ID).Delete(ShopStyle{}).Error; err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}

		//添加新的款式和新图片链接
		for i, v := range req.ShopStyles {
			//添加款式
			v.ShopStyle.ID = 0
			v.ShopStyle.ShopId = req.Shop.ShopInfo.ID
			shopStyle := v.ShopStyle
			if err := tx.Create(&shopStyle).Error; err != nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
			//添加款式图片
			for j, _ := range req.ShopStyles[i].ImgStyle {
				req.ShopStyles[i].ImgStyle[j].CreateAt = utils.GetStringDateTime(time.Now())
				req.ShopStyles[i].ImgStyle[j].TargetId = shopStyle.ID
				err = tx.Create(&req.ShopStyles[i].ImgStyle[j]).Error
				if err != nil {
					tx.Rollback()
					g.FResponse(nil)
					return
				}
			}
		}

		req.Shop.ShopInfo.UpdateAt = utils.GetStringDateTime(time.Now())
		//更新商品
		shopMap := utils.StructToJsonMap(req.Shop.ShopInfo)
		delete(shopMap, "id")
		delete(shopMap, "create_at")
		delete(shopMap, "shop_style")
		delete(shopMap, "image")
		delete(shopMap, "shop_type")
		if err := tx.Table("shop_info").Where("id=?", req.Shop.ShopInfo.ID).Update(shopMap).Error; err != nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		tx.Commit()
		g.Response(nil)
	} else {
		g.PResponse(err)
	}
	//newShopImages := req.Shop.ImgShop
	//var oldShopImages []Image
	//var deleteImage []string
	//var oldStyle []ShopStyle
	//db.SqlDB.Find(&oldShopImages,"target_id=?",req.Shop.ShopInfo.ID)
	//db.SqlDB.Find(&oldStyle,"shop_id=?",req.Shop.ShopInfo.ID)
	//tx := db.SqlDB.Begin()
	//for _,new:=range newShopImages{
	//	flag :=false
	//	for i,old := range oldShopImages{
	//		if new.ImgUrl==old.ImgUrl {
	//			//后面旧的ID不为0的为要删除的
	//			oldShopImages[i].ID=0
	//			flag=true
	//		}
	//	}
	//	//需要插入
	//	if !flag {
	//		err :=tx.Create(&new).Error
	//		if err!=nil {
	//			tx.Rollback()
	//			g.FResponse(nil)
	//			return
	//		}
	//	}
	//}
	////删除旧的商品图片数据
	//for _,v:=range oldShopImages{
	//	if v.ID!=0 {
	//		if err :=tx.Delete(&v).Error;err!=nil{
	//			tx.Rollback()
	//			g.FResponse(nil)
	//			return
	//		}
	//		deleteImage=append(deleteImage, RuntimeRootPath+v.ImgUrl)
	//	}
	//}
	//req.Shop.ShopInfo.UpdateAt=utils.GetStringDateTime(time.Now())
	////更新商品
	//shopMap:=utils.StructToJsonMap(req.Shop.ShopInfo)
	//delete(shopMap, "id")
	//delete(shopMap,"create_at")
	//if err :=tx.Table("shop_info").Where("id=?",req.Shop.ShopInfo.ID).Update(shopMap).Error;err!=nil{
	//	tx.Rollback()
	//	g.FResponse(nil)
	//	return
	//}
	//
	////款式修改
	//for _ ,v :=range req.ShopStyles {
	//	newStyleImages :=v.ImgStyle
	//	var oldStyleImages []Image
	//	db.SqlDB.Find(&oldStyleImages,"target_id=?",v.ShopStyle.ID)
	//	for _,new:=range newStyleImages{
	//		flag := false
	//		for i,old:=range oldStyleImages{
	//			if  new.ImgUrl==old.ImgUrl{
	//				//后面旧的ID不为0的为要删除的
	//				oldStyleImages[i].ID=0
	//				flag=true
	//			}
	//		}
	//		if !flag {
	//			err :=tx.Create(&new).Error
	//			if err!=nil {
	//				tx.Rollback()
	//				g.FResponse(nil)
	//				return
	//			}
	//		}
	//	}
	//	for _,old:=range oldStyleImages {
	//		if old.ID!=0 {
	//			if err :=tx.Delete(&old).Error;err!=nil{
	//				tx.Rollback()
	//				g.FResponse(nil)
	//				return
	//			}
	//			deleteImage=append(deleteImage,RuntimeRootPath+old.ImgUrl)
	//		}
	//	}
	//
	//}
	//tx.Commit()
	//g.Response(nil)
	//删除磁盘存储图片
	//go deleteImages(deleteImage)
}

func deleteImages(imgs []string) {
	for _, v := range imgs {
		DeleteTargetImg(v)
	}
}
