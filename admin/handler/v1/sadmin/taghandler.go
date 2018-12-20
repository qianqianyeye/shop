package sadmin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "shop/admin/constant"
	. "shop/admin/handler/model"
	"shop/admin/mysql"
	"time"
	. "shop/admin/utils"
	"strings"
)
func GetTagsType(c *gin.Context)  {
	id :=c.Query("id")
	ids :=strings.Split(id,",")
	rsp := RspShopTypeList{Code: RC_OK, Msg: M(RC_OK)}
	if id!="" {
		var ShopTypeList []ShopType
		var tags []TagType
		var shopTypeIds []int64
		db.SqlDB.Find(&tags,"tag_id in (?)",ids)
		for _,v:=range tags {
			shopTypeIds=append(shopTypeIds,v.ShopTypeId)
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

func AddTag(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var req Tag
	if err := c.ShouldBind(&req); err == nil {
		req.Time=GetStringDateTime(time.Now())
		err:=db.SqlDB.Create(&req).Error
		if err!=nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
		}
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR),Data:nil}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

func UpdateTag(c *gin.Context)  {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var req Tag
	if err := c.ShouldBind(&req); err == nil {
		req.Time=GetStringDateTime(time.Now())
		Tag :=Tag{ID:req.ID,Name:req.Name,RName:req.RName,Time:req.Time}
		err:=db.SqlDB.Save(&Tag).Error
		if err!=nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
		}
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR),Data:nil}
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

func DeleteTag(c *gin.Context)  {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	id := c.Param("id")
	if id != "" {
		tx := db.SqlDB.Begin()
		if err := tx.Where("tag_id=?", id).Delete(TagType{}).Error; err != nil {
			tx.Rollback()
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
		}

		if err := tx.Where("id=?", id).Delete(Tag{}).Error; err != nil {
			tx.Rollback()
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
		}
		tx.Commit()
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

func DeleteShopTag(c *gin.Context)  {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	tagid := c.Query("tag_id")
	shopid := c.Query("shop_id")
	if tagid != "" && shopid!="" {
		if err := db.SqlDB.Where("tag_id=? and shop_info_id=?", tagid,shopid).Delete(TagType{}).Error; err != nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
		}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}