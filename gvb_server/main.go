package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb/gvb_server/core"
	"gvb/gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	//初始化日志
	global.Log = core.Initlogger()
	global.Log.Warnln("123")
	global.Log.Error("123")
	global.Log.Infof("123")

	logrus.Warnln("123")
	logrus.Error("123")
	logrus.Infof("123")

	//连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)

}
