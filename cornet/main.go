package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

var (
	ConfigNamePath = "cornet/config/app"
	clientManager  = NewClientManager()
)

func init() {
	viper.SetConfigName(ConfigNamePath)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	webSocketPort := viper.GetString("app.webSocketPort")

	// 开始监控用户事件
	go clientManager.start()

	// uid 唯一id 最好是用户的id/或者能够关联到用户的
	// app_id：应用id--> web/app/小程序
	// ws://127.0.0.1:8888/ws?uid=342342&app_id=1
	http.HandleFunc("/ws", WsController)
	if err := http.ListenAndServe(":"+webSocketPort, nil); err != nil {
		panic(err)
	}
}
