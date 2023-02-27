package main

import (
	"fmt"
	"go-web/bluebell/controller"
	"go-web/bluebell/dao/mysql"
	"go-web/bluebell/pkg/snowflake"
	"go-web/bluebell/routers"

	"go-web/bluebell/logger"
	"go-web/bluebell/settings"

	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}
	// 加载配置
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	//if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init redis failed, err:%v\n", err)
	//	return
	//}
	//defer redis.Close()

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	//if err := controller.InitTrans("zh"); err != nil {
	//	fmt.Printf("init validator trans failed, err:%v\n", err)
	//	return
	//}

	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
