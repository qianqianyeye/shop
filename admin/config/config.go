package config

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"

	"git.jiaxianghudong.com/go/utils"
)

var ostype = runtime.GOOS

var conf ConfAPI

type ConfAPI struct {
	Listen          int               `yaml:"listen"`           // 监听端口
	RunMode         string            `yaml:"runmode"`          // 服务运行模式
	TimestampExpire int               `yaml:"timestamp_expire"` // 请求时间戳过期时长
	NonceRepeat     int               `yaml:"nonce_repeat"`     // nonce串可重复间隔
	SessionExpire   int               `yaml:"session_expire"`   // session过期时长
	ServerUrl       string            `yaml:"server_url"`
	PayTransferUrl  string            `yaml:"pay_transfer_url"`
	IapVerifyUrl    string            `yaml:"iap_verify_url"`
	ExternalIp      string            `yaml:"external_ip"`
	Logs            EntityLogs        `yaml:"logs"` // 日志
	Mysql           EntityMysql       `yaml:"mysql"`
	Redis           EntityRedis       `yaml:"redis"`
}

//
type EntityLogs struct {
	Dir      string `yaml:"dir"`
	File     string `yaml:"file"`
	Level    int    `yaml:"level"`
	SaveFile bool   `yaml:"savefile"`
}

type EntityMysql struct {
	Addr     string `yaml:"addr"`
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
}

//
type EntityRedis struct {
	Addr string `yaml:"addr"`
	Pwd  string `yaml:"password"`
	DB   int64  `yaml:"db"`
}



func Init() {
	fmt.Println("--init config start")

	// 初始化日志
	str := utils.ReadConfFile("shop/admin/shop.yaml")
	/* 替换注释 */
	reg := regexp.MustCompile(`\/\*[^(\*\/)]*\*\/`)
	str = reg.ReplaceAllString(str, "")

	/* 解析yaml */
	err := yaml.Unmarshal([]byte(str), &conf)
	/* 配制全局变量 */
	if nil == err {

		// host config
		fmt.Println(fmt.Sprintf(" Listen:%d", conf.Listen))
		fmt.Println(fmt.Sprintf(" RunMode:%s", conf.RunMode))
		fmt.Println(fmt.Sprintf(" TimestampExpire:%d", conf.TimestampExpire))
		fmt.Println(fmt.Sprintf(" RunMode:%s", conf.RunMode))


		// logs
		fmt.Println(fmt.Sprintf(" LOG:{Dir:%s\tFile:%s\tLevel:%d\tSaveFile:%t}",
			conf.Logs.Dir, conf.Logs.File, conf.Logs.Level, conf.Logs.SaveFile))
		// mysql
		fmt.Println(fmt.Sprintf(" Mysql:{Addr:%s\tUserName:%s\tPwd:%s\tDB:%s\tMaxOpen:%d\tMaxIdle:%d}}",
			conf.Mysql.Addr, conf.Mysql.UserName, conf.Mysql.Password,
			conf.Mysql.Db, conf.Mysql.MaxOpen, conf.Mysql.MaxIdle))
		// redis
		fmt.Println(fmt.Sprintf(" REDIS:{Addr:%s\tPwd:%s\tDB:%d\t}",
			conf.Redis.Addr, conf.Redis.Pwd, conf.Redis.DB))

		fmt.Println("--init config end")

	} else {
		/* 打印日志 */
		fmt.Println("--init config err :", err.Error())
	}
}


// 获取日志配置
func GetLogs() *EntityLogs {
	if GetRunMode() == "debug" {
		conf.Logs.Dir = "../log"
		if ostype == "windows" {
			conf.Logs.Dir = "..\\log"
		}
	}
	return &conf.Logs
}

// 获取运行模式
func GetRunMode() string {
	return strings.TrimSpace(strings.ToLower(conf.RunMode))
}

func GetMysql() *EntityMysql {
	return &conf.Mysql
}

func GetRedis() *EntityRedis {
	return &conf.Redis
}

func GetListen() int {
	return conf.Listen
}