package smobile

import (
	"github.com/gin-gonic/gin"
	"shop/api/mysql"
	"shop/api/handler/model"
	"net/http"
	."shop/api/constant"
)
type RspVersion struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	AppVersion model.AppVersion `json:"app_version"`
}
func GetLatestVersion (c *gin.Context)  {
	appType :=c.Param("app_type")
	rsp :=RspVersion{Code:RC_OK,Msg:M(RC_OK)}
	if appType != "" {
		var appVersion model.AppVersion
		db.SqlDB.Order("id desc").Limit(1).Find(&appVersion,"type=?",appType)
		rsp.AppVersion=appVersion
		c.JSON(http.StatusOK,rsp)
	}else {
		rsp.Code=RC_PARM_ERR
		rsp.Msg=M(RC_PARM_ERR)
		c.JSON(http.StatusOK,rsp)
	}

}