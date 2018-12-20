package main

import (
	"fmt"
	"os"
	"shop/api/config"
	"shop/api/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	m "shop/api/handler/v1/smobile"
	"git.jiaxianghudong.com/go/logs"
	"git.jiaxianghudong.com/go/redis"
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

	//移动端
	mobile := router.Group("/mobile/v1")
	mobile.GET("shop/tags",m.QueryTag)
	mobile.GET("shop/tagstype",m.GetTagsType)
	mobile.GET("shop/type",m.GetShopType) //获取商品类别
	mobile.GET("shop/list",m.GetShopList) //商品列表
	mobile.GET("shop/details/:id",m.GetShopDetails) //商品详情
	mobile.GET("shop/hot",m.GetHotList)
	mobile.GET("shop/search",m.GetShopSearch) //搜索商品
	mobile.GET("shop/latestversion/:app_type",m.GetLatestVersion) //获取app最新版本号

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



