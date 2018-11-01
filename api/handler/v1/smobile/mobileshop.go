package smobile

import (
	"github.com/gin-gonic/gin"
	. "shop/handler/model"
	"shop/mysql"
)

type RspShop struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	ShopType ShopType `json:"shop_type"`
}

// @Summary 获取分类信息
// @Produce  json
// @Success 200 {string} json "{"code":0,"msg":"ok","shop_type":{"id":2,"parent_id":0,"c_name":"测试2","r_name":" тест 2","created_at":"14:57:22","update_at":"14:57:27"}}"
// @Router /mobile/v1/shop/type [get]
func GetShopType(c *gin.Context)  {
	g :=Gin{c}
	var shop ShopType
	db.SqlDB.Find(&shop)
	//rsp := RspShop{Code:RC_OK,Msg:M(RC_OK),ShopType:shop}
	//rsp := RspShop{Code:RC_SYS_ERR,Msg:M(RC_SYS_ERR)}
	g.Response(shop)
	//c.JSON(http.StatusOK,rsp)
}

func GetShopList(c *gin.Context)  {
	fmt
}

func GetShopDetails(c *gin.Context)  {

}

func GetShopSearch(C *gin.Context)  {

}

