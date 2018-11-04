package main

import (
	"git.jiaxianghudong.com/go/logs"
	"fmt"
	"git.jiaxianghudong.com/go/redis"
	"os"
	"shop/admin/config"
	"shop/admin/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	a "shop/admin/handler/v1/sadmin"
	_ "shop/admin/docs"
	"net/http"
)


func main()  {
	// 初始化配置
	config.Init()

	// 初始化日志
	initLogs()
	// 初始化mysql
	err := initMysql()
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to init mysql ,err: %s", err.Error()))
		os.Exit(1)
	}
	defer db.SqlDB.Close()
	//err = initRedis()
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("failed to init redis notice,err: %s", err.Error()))
	//	os.Exit(1)
	//}

	router := gin.Default()
	ginCorsCfg := cors.DefaultConfig()
	ginCorsCfg.AllowAllOrigins = true
	router.Use(cors.New(ginCorsCfg))
	router.StaticFS("/upload/images/", http.Dir(a.GetImageFullPath()))

	//后台
	//商品类别
	admin := router.Group("/admin/v1")
	admin.GET("shop/typelist",a.GetShopType)  //商品类别
	admin.POST("shop/atype",a.AddShopType)  //添加商品类别
	admin.PATCH("shop/utype",a.UpdateShopType)  //修改商品类别
	admin.DELETE("shop/dtype/:id",a.DeleteShopType) //删除商品类别

	//关键词
	admin.GET("shop/hotkey",a.GetHotKey)
	admin.POST("shop/ahotkey",a.AddHotKey)
	admin.PATCH("shop/uhotkey",a.UpdateHotKey)
	admin.DELETE("shop/dhotkey/:id",a.DeleteHotKey)

	//商品
	admin.GET("shop/list",a.GetShopList)
	admin.POST("shop/ashop",a.AddShop)
	admin.PATCH("shop/ushop",a.UpdateShop)
	admin.DELETE("shop/dshop/:id",a.DeleteShop)

	//版本控制
	admin.GET("shop/versionlist",a.Versionlist)
	admin.POST("shop/aversion",a.AddVersion)
	admin.PATCH("shop/uversion",a.UpdateVersion)
	admin.DELETE("shop/dversion/:id",a.DeleteVersion)

	//管理员用户
	//admin.GET("user/list",a.GetUserList)
	//admin.POST("user/auser",a.AddUser)
	//admin.PATCH("user/uuser",a.UpdateUser)
	//admin.DELETE("user/duser/:id",a.DeleteUser)

	//登陆
	admin.POST("login",a.Login)
	admin.PATCH("update/userinfo",a.UpdateUserInfo)

	//图片上传
	admin.POST("/upload", a.UploadImage)
	admin.POST("/deleteImage",a.DeleteImg)
	//router.Run(":5054")
	router.Run(fmt.Sprintf(":%d", config.GetListen()))
}


// 初始化日志模块
func initLogs() {
	options := config.GetLogs()
	logs.Init(options.Dir, options.File, options.Level, options.SaveFile)
}

// 初始化mysql
func initMysql() error {
	options := config.GetMysql()
	err := db.InitDB(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		options.UserName, options.Password, options.Addr, options.Db), options.MaxOpen, options.MaxIdle)
	return err
}

// 初始化redis
func initRedis() error {
	options := config.GetRedis()
	err := redis.Init(options.Addr, options.Pwd, options.DB)
	return err
}

