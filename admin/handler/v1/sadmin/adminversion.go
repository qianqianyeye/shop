package sadmin

import (
	"git.jiaxianghudong.com/go/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	. "shop/admin/constant"
	. "shop/admin/handler/model"
	"shop/admin/mysql"
	. "shop/admin/utils"
	"time"
)

type RspVersionList struct {
	Code           int          `json:"code"`
	Msg            string       `json:"msg"`
	Data   interface{} `json:"data"`
	//AppVersionList []AppVersion `json:"app_version_list"`
	PageModel      PageModel    `json:"page"`
}

func Versionlist(c *gin.Context) {
	rsp := RspVersionList{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var req ReqShopList
	if err := c.ShouldBindWith(&req, binding.Query); err == nil {
		page := GetPageInfo(req.Page_size, req.Current)
		var err error
		var AppVersionList []AppVersion
		var count Count
		err = db.SqlDB.Order("update_at desc").Offset(page.OffSet).Limit(page.PageSize).Find(&AppVersionList).Error
		err = db.SqlDB.Table("app_version").Select("count(*) count").Scan(&count).Error
		if err != nil {
			rsp.Code = RC_SYS_ERR
			rsp.Msg = M(RC_SYS_ERR)
			logs.Error(err)
			c.JSON(http.StatusOK, rsp)
		}
		rsp.Data = AppVersionList
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

func AddVersion(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var req AppVersion
	if err := c.ShouldBind(&req); err == nil {
		req.CreateAt = GetStringDateTime(time.Now())
		req.UpdateAt = GetStringDateTime(time.Now())
		if err := db.SqlDB.Create(&req).Error; err != nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR),Data:nil}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

func UpdateVersion(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var req AppVersion
	if err := c.ShouldBind(&req); err == nil {
		req.UpdateAt = GetStringDateTime(time.Now())
		app := StructToJsonMap(req)
		delete(app, "id")
		delete(app, "create_at")
		if err := db.SqlDB.Table("app_version").Where("id=?", req.ID).Update(app).Error; err != nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR),Data:nil}
		}
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR),Data:nil}
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}

func DeleteVersion(c *gin.Context) {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	id := c.Param("id")
	if id != "" {
		db.SqlDB.Where("id=?", id).Delete(AppVersion{})
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
