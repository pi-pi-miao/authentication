package main

import (
	"authentication/config"
	"authentication/pkg/db"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"

	"authentication/api"
)

var (
	path = flag.String("path","","")
)

func main(){
	flag.Parse()
	fmt.Println("path is ",*path)
	if *path == "" {
		fmt.Println("please set config path")
		return
	}

	if err := config.Init(*path);err != nil {
		fmt.Println("init config err",err)
		return
	}
	if err := db.InitDb(config.Auth.AuthConfig.Name,config.Auth.AuthConfig.Passwd,
		config.Auth.AuthConfig.Ip,config.Auth.AuthConfig.Port,config.Auth.AuthConfig.Store);err != nil {
		fmt.Println("init db err",err)
		return
	}
	r := gin.Default()
	api.Router(r).Run(fmt.Sprintf("%v:%v",config.Auth.HttpConfig.Ip,config.Auth.HttpConfig.Port))
}