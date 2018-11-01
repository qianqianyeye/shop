package sadmin

import (
	"github.com/gin-gonic/gin"
	."shop/api/handler/model"
	"time"
	"shop/api/utils"
	"shop/api/mysql"
)



func GetShopType(c *gin.Context)  {

}

func AddShopType(c *gin.Context)  {

}

func DeleteShopType(c *gin.Context)  {

}

func UpdateShopType(c *gin.Context)  {


}
func GetHotKey(c *gin.Context)  {

}

func AddHotKey(c *gin.Context)  {

}

func DeleteHotKey(c *gin.Context)  {

}

func UpdateHotKey(c *gin.Context)  {

}

func GetShopList(c *gin.Context)  {

}

type ReqShopInfo struct {
	ShopInfo ShopInfo `json:"shop_info"`
	ShopStyle ShopStyle `json:"shop_style"`
	ImgStyle []Image `json:"img_style"`
	ImgShop  []Image `json:"img_shop"`
}

// @Summary 添加商品
// @Description add by json ReqShopInfo
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param  shop_info body model.ShopInfo true "商品信息"
// @Param  shop_style body model.ShopStyle true "款式"
// @Param  {array} img_style body model.Image true "款式图片"
// @Param  {array} img_shop body model.Image true "商品图片"
// @Success 200 {string} json "{"code":0,"data":null,"msg":"ok"}"
// @Router /admin/v1/shop/ashop [post]
func AddShop(c *gin.Context)  {
	g :=Gin{c}
	var req ReqShopInfo
	if err :=c.ShouldBind(&req);err==nil{
		tx := db.SqlDB.Begin()
		req.ShopInfo.CreateAt=utils.GetStringDateTime(time.Now())
		req.ShopInfo.UpdateAt=utils.GetStringDateTime(time.Now())
		shop := req.ShopInfo
		err :=tx.Create(&shop).Error
		if err!=nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		req.ShopStyle.ShopId=shop.ID
		style := req.ShopStyle
		err =tx.Create(&style).Error
		if err!=nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		// 1商品 2款式
		for i,_:=range req.ImgShop {
			req.ImgShop[i].CreateAt=utils.GetStringDateTime(time.Now())
			req.ImgShop[i].TargetId=shop.ID
			err=tx.Create(&req.ImgShop[i]).Error
			if err!=nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
		}
		for i,_:=range req.ImgStyle {
			req.ImgStyle[i].CreateAt=utils.GetStringDateTime(time.Now())
			req.ImgStyle[i].TargetId=style.ID
			err=tx.Create(&req.ImgStyle[i]).Error
			if err!=nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
		}
		tx.Commit()
		g.Response(nil)
	}else {
		g.PResponse(err)
	}
}

func DeleteShop(c *gin.Context)  {

}

func UpdateShop(c *gin.Context)  {

}
