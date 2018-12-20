package sadmin

import (

	"github.com/gin-gonic/gin"

	"shop/admin/handler/model"
	"git.jiaxianghudong.com/go/logs"
	"shop/admin/mysql"
	"os"
	"strings"
	"shop/admin/config"
)

type reqImg struct {
	ImgUrl string `json:"img_url" binding:"required"`
	//Id int `json:"id" binding:"required"`
}

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
	//savePath := GetImagePath()
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
		"img_url":     "http://"+GetImageFullUrl(imageName),
		//"img_show_url": GetImageFullUrl(imageName),
	})
}

func DeleteTargetImg(src string){
	src = strings.Replace(src,"http://"+config.GetExternalIp() + "/","",1)
	os.Remove(RuntimeRootPath+src)
}

func DeleteImg(c *gin.Context)  {
	appG := model.Gin{c}
	var reqImg reqImg
	var img  model.Image
	if err :=c.ShouldBind(&reqImg);err==nil{
		DeleteTargetImg(reqImg.ImgUrl)
		db.SqlDB.Find(&img,"img_url=?",reqImg.ImgUrl)
		if img.ID!=0 {
			db.SqlDB.Delete(&img)
		}
		appG.Response(nil)
	}else {
		appG.PResponse(err)
	}
}
