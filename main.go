package main

import (
	"bubble/dao"
	"bubble/routers"
	"bubble/setting"
	"fmt"
	"os"
)

/**
	运行：go build
	bubble.exe conf/config.ini
 */

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./bubble conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}


	err := dao.InitSqlite()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 程序退出关闭数据库连接
	defer dao.CloseSqlite()
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
