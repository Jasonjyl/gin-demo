package main

import (
	"fmt"
	"github.com/Jasonjyl/gin-demo/dao"
	"github.com/Jasonjyl/gin-demo/models"
	"github.com/Jasonjyl/gin-demo/routers"
	"github.com/Jasonjyl/gin-demo/setting"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./gin-demo/conf/conf.ini")
		return
	}
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, error: %v\n", err)
		return
	}

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer dao.Close()

	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err: %v\n", err)
	}
}
