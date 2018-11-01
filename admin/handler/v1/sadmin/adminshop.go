package sadmin

import (
	"github.com/gin-gonic/gin"
	."shop/admin/handler/model"
	"time"
	"shop/admin/utils"
	"shop/admin/mysql"
	"net/http"
	"fmt"
	"encoding/json"
)



func GetShopType(c *gin.Context)  {
	maps := map[string]interface{}{"code":0,"msg":"ok"}
	c.JSON(http.StatusOK,gin.H{"s":maps})
}

func AddShopType(c *gin.Context)  {
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded;charset=utf-8;param=value"
	parm := make(map[string]string)
	rsp, err :=PostMap("http:127.0.0.1:5054/admin/v1/shop/type", parm, header, true)
	if err != nil {
		fmt.Println(err)
	}
	if len(rsp) == 0 {
		fmt.Println("fali to notify h5server")
	}
	rspH5Notify := &RspH5Notify{}
	err = json.Unmarshal(rsp, rspH5Notify)
	if err != nil {
		fmt.Println( err)
	}

	if rspH5Notify.Code != 0 {
		fmt.Println( err)
	}
}
type RspH5Notify struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
// post
func PostMap(apiUrl string, parm map[string]string, header map[string]string, isHttps bool) ([]byte, error) {

	//data := url.Values{}
	//for k, v := range parm {
	//	data.Set(k, v)
	//}
	//
	//reqParams := ioutil.NopCloser(strings.NewReader(data.Encode()))
	//client := &http.Client{}
	//
	//if isHttps {
	//	client.Transport = &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	}
	//}
	//reqest, _ := http.NewRequest("POST", apiUrl, reqParams)
	//
	//for k, v := range header {
	//	reqest.Header.Set(k, v)
	//}
	//
	//response, err := client.Do(reqest)
	//if nil != err {
	//	return nil, err
	//}
	//
	//defer response.Body.Close()
	//if response.StatusCode != 200 {
	//	return nil, errors.New(response.Status)
	//}
	//
	//body, err := ioutil.ReadAll(response.Body)
	//if nil != err {
	//	return nil, err
	//}
	//
	//return body, nil
	return nil,nil
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
	c.Redirect(302,"www.baidu.com")
}

type ReqAddShopStyle struct {
	ShopStyle ShopStyle `json:"shop_style"`
	ImgStyle  []Image `json:"img_style"`
}
type ReqAddShopInfo struct {
	ShopInfo ShopInfo `json:"shop_info"`
	ImgShop []Image `json:"img_shop"`
}
type ReqShopInfo struct {
	Shop ReqAddShopInfo `json:"shop"`
	ShopStyles []ReqAddShopStyle `json:"shop_styles"`
}


func AddShop(c *gin.Context)  {
	g :=Gin{c}
	var req ReqShopInfo
	if err :=c.ShouldBind(&req);err==nil{
		tx := db.SqlDB.Begin()
		req.Shop.ShopInfo.CreateAt=utils.GetStringDateTime(time.Now())
		req.Shop.ShopInfo.UpdateAt=utils.GetStringDateTime(time.Now())
		shop := req.Shop.ShopInfo
		err :=tx.Create(&shop).Error
		if err!=nil {
			tx.Rollback()
			g.FResponse(nil)
			return
		}
		// 1商品 2款式
		//插入商品图片
		for i,_ := range req.Shop.ImgShop {
			req.Shop.ImgShop[i].TargetId=shop.ID
			req.Shop.ImgShop[i].CreateAt=utils.GetStringDateTime(time.Now())
			err=tx.Create(&req.Shop.ImgShop[i]).Error
			if err!=nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
		}

		//插入款式跟款式图片
		for i,_:=range req.ShopStyles {
			req.ShopStyles[i].ShopStyle.ShopId=shop.ID
			style := req.ShopStyles[i].ShopStyle
			err =tx.Create(&style).Error
			if err!=nil {
				tx.Rollback()
				g.FResponse(nil)
				return
			}
			for j,_ := range req.ShopStyles[i].ImgStyle{
				req.ShopStyles[i].ImgStyle[j].CreateAt=utils.GetStringDateTime(time.Now())
				req.ShopStyles[i].ImgStyle[j].TargetId=style.ID
				err=tx.Create(&req.ShopStyles[i].ImgStyle[j]).Error
				if err!=nil {
					tx.Rollback()
					g.FResponse(nil)
					return
				}
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
