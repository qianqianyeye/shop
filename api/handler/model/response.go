package model

import (
	"github.com/gin-gonic/gin"
	."shop/constant"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": RC_OK,
		"msg": M(RC_OK) ,
		"data": data,
	})
	return
}

func (g *Gin) FResponse(data interface{})  {
	g.C.JSON(http.StatusOK, gin.H{
		"code": RC_SYS_ERR,
		"msg": M(RC_SYS_ERR) ,
		"data": data,
	})
	return
}

func (g *Gin) PResponse(data interface{})  {
	g.C.JSON(http.StatusOK, gin.H{
		"code": RC_PARM_ERR,
		"msg": M(RC_PARM_ERR) ,
		"data": data,
	})
	return
}

func (g *Gin) UpResponse(data interface{})  {
	g.C.JSON(http.StatusOK, gin.H{
		"code": RC_UPLOAD_FORMAT,
		"msg": M(RC_UPLOAD_FORMAT) ,
		"data": data,
	})
	return
}

func (g *Gin)UpFResponse(data interface{})  {
	g.C.JSON(http.StatusOK, gin.H{
		"code": RC_UPLOAD_FORMAT,
		"msg": M(RC_UPLOAD_FORMAT) ,
		"data": data,
	})
	return
}
