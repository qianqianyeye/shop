package smobile

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/api/mysql"
	."shop/api/handler/model"
	."shop/api/constant"
	"strings"
)
type RspCommon struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func GetTagsType(c *gin.Context)  {
	id :=c.Query("id")
	ids :=strings.Split(id,",")
	rsp := RspShopTypeList{Code: RC_OK, Msg: M(RC_OK)}
	if id!="" {
		var ShopTypeList []ShopType
		var tags []TagType
		var shopTypeIds []int64
		db.SqlDB.Find(&tags,"tag_id in (?)",ids)
		tempMap := make(map[int64]int64)
		for _,v:=range tags {
			tempMap[v.ShopTypeId]=v.TagId
		}
		for k,_:=range tempMap {
			shopTypeIds=append(shopTypeIds,k)
		}
		db.SqlDB.Order("update_at desc").Preload("Image","img_type=3").Find(&ShopTypeList,"id in (?)",shopTypeIds)
		rsp.Data=ShopTypeList
	}else {
		rsp.Code=RC_PARM_ERR
		rsp.Msg=M(RC_PARM_ERR)
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
func QueryTag(c *gin.Context)  {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var tag []Tag
	err :=db.SqlDB.Find(&tag).Error
	if err!=nil {
		rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
	}else {
		rsp.Data=tag
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}