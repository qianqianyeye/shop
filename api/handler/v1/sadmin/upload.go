package sadmin

import (

	"github.com/gin-gonic/gin"

	"shop/api/handler/model"
	"git.jiaxianghudong.com/go/logs"
	"net/http"
)

// @Summary 上传图片
// @Accept  multipart/form-data
// @Produce  json
// @Param image formData file true "图片文件"
// @Success 200 {string} json "{"code":0,"data":{"img_save_url":"upload/images/8f81dcf6-b970-4ed1-bd02-ccf1c908c0f1.jpg","img_url":"localhost:5054/upload/images/8f81dcf6-b970-4ed1-bd02-ccf1c908c0f1.jpg"},"msg":"ok"}"
// @Router /admin/v1/upload [post]
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