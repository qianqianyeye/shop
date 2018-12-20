package sadmin

import (
	"github.com/gin-gonic/gin"
	. "shop/admin/handler/model"
	. "shop/admin/constant"
	"net/http"
	"strings"
	"git.jiaxianghudong.com/go/utils"
	"github.com/satori/go.uuid"
	"shop/admin/mysql"
	"crypto/md5"
	"fmt"
	"sync"
	"time"
	"shop/api/constant"
)

type RspAdmin struct {
	Code         int        `json:"code"`
	Msg          string     `json:"msg"`
	Data interface{} `json:"data"`
	Token string `json:"token"`
}
var AdminMap *sync.Map
func init() {
	AdminMap=new(sync.Map)
}

type AdminMapModel struct {
	Admin Admin
	t time.Time
}

func GetCurStatus(c *gin.Context)  {
	token:=c.Request.Header.Get("token")
	if v,ok:=AdminMap.Load(token);ok {
		temp:=v.(AdminMapModel)
		if time.Now().Unix()-temp.t.Unix()>60*60*2{
			AdminMap.Delete(token)
			c.JSON(http.StatusOK,RspCommon{Code:constant.RC_ADMIN_EXPIRE,Msg:M(constant.RC_ADMIN_EXPIRE),Data:nil})
			return
		}
		var AM AdminMapModel
		AM.t=time.Now()
		AM.Admin=temp.Admin
		AdminMap.Store(token,AM)
		c.JSON(http.StatusOK,RspCommon{Code:constant.RC_OK,Msg:M(constant.RC_OK),Data:temp.Admin})
		return
	}
	c.JSON(http.StatusOK,RspCommon{Code:constant.RC_ADMIN_EXPIRE,Msg:"token not found",Data:nil})
	return
}
func Login(c *gin.Context)  {
	rsp := RspAdmin{Code: RC_OK, Msg: M(RC_OK),Data:nil}
	var req Admin
	if err := c.ShouldBind(&req); err == nil {
		//tx:=db.SqlDB.Begin()
		var admin Admin
		err :=db.SqlDB.Find(&admin).Error
		if err !=nil {
			rsp = RspAdmin{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
			c.JSON(http.StatusOK,rsp)
			return
		}
		if admin.ID == 0 {
			rsp.Code = RC_ADMIN_NOT_FOUND
			rsp.Msg = M(RC_ADMIN_NOT_FOUND)
			c.JSON(http.StatusOK,rsp)
			return
		}
		// 验证密码
		if strings.ToUpper(admin.Password) != strings.ToUpper(GetMd5(req.Password)) {
			rsp.Code = RC_PASSWORD_ERR
			rsp.Msg = M(RC_PASSWORD_ERR)
			c.JSON(http.StatusOK,rsp)
			return
		}
		guid, _ := uuid.NewV4()
		token := utils.Md5Sum(guid.String())
		rsp.Token=token
		rsp.Data=admin
		AM:=AdminMapModel{Admin:admin,t:time.Now()}
		AdminMap.Store(token,AM)
		c.JSON(http.StatusOK,rsp)
	} else {
		rsp = RspAdmin{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR),Data:nil}
		c.JSON(http.StatusOK,rsp)
	}
}

func UpdateUserInfo(c *gin.Context)  {
	rsp := RspCommon{Code: RC_OK, Msg: M(RC_OK)}
	var req Admin
	if err := c.ShouldBind(&req); err == nil {
		adminMap :=make(map[string]string)
		adminMap["name"]=req.Name
		adminMap["password"]=GetMd5(req.Password)
		err :=db.SqlDB.Table("admin").Where("id=?",req.ID).Update(adminMap).Error
		if err!=nil {
			rsp = RspCommon{Code: RC_SYS_ERR, Msg: M(RC_SYS_ERR)}
		}
		c.JSON(http.StatusOK,rsp)
	} else {
		rsp = RspCommon{Code: RC_PARM_ERR, Msg: M(RC_PARM_ERR)}
		c.JSON(http.StatusOK,rsp)
	}
}

func GetMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}