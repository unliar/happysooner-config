package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	happyConfig "github.com/unliar/happysooner-config"
)

func main() {
	// my personal default config
	c, err := happyConfig.NewHappyDefaultConf()
	if err != nil {
		fmt.Println("初始化出错")
		panic(err)
	}
	r, err := c.GetMySQL("mysql")
	if err!=nil {
		fmt.Println("读取配置出错")
		panic(err)
	}
	fmt.Println(r)

	// customer config
	fmt.Println("customer config")
	hf:=happyConfig.Conf{}
	ccc,err := config.NewConfig()
	if err!=nil {
		panic(err)
	}
	hf.Init = ccc
	rr,err := hf.GetMySQL("mysql")
	fmt.Println(rr,err)
}
