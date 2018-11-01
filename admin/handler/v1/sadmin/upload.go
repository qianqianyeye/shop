package sadmin

import (

	"github.com/gin-gonic/gin"

	"shop/admin/handler/model"
	"git.jiaxianghudong.com/go/logs"
	"net/http"
)

func UploadImage(c *gin.Context) {
	appG := model.Gin{c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.FResponse(nil)
		return
	}

	if image == nil {
		appG.PResponse(nil)
		return
	}

	imageName := GetImageName(image.Filename)
	fullPath := GetImageFullPath()
	savePath := GetImagePath()
	src := fullPath + imageName

	if !CheckImageExt(imageName) || !CheckImageSize(file) {
		appG.UpResponse(nil)
		return
	}

	err = CheckImage(fullPath)
	if err != nil {
		logs.Error(err)
		appG.UpFResponse(nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logs.Error(err)
		appG.UpFResponse(nil)
		return
	}

	appG.Response(map[string]string{
		"img_url":     GetImageFullUrl(imageName),
		"img_save_url": savePath + imageName,
	})
}

func GetImg(c *gin.Context)  {
	img :=http.Dir(GetImageFullPath())
	c.JSON(http.StatusOK,img)
}