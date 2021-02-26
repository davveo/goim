package main

import (
	"log"
	"net/http"

	"github.com/davveo/goim/cornet/consts/code"

	"github.com/gorilla/websocket"
)

const (
	// 最大的消息大小
	maxMessageSize = 8192
)

type Data struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WsController(w http.ResponseWriter, req *http.Request) {
	uid := req.FormValue("uid")
	appId := req.FormValue("app_id")

	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, req, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		http.NotFound(w, req)
		return
	}

	if len(uid) == 0 || len(appId) == 0 {
		conn.WriteJSON(Data{
			Code: code.SYSTEM_ID_ERROR,
			Msg:  "用户id或者appId非法",
			Data: nil,
		})
		conn.Close()
		return
	}

	conn.SetReadLimit(maxMessageSize)
	remoteAddr := conn.RemoteAddr().String()

	client := NewClient(remoteAddr, appId, uid, conn)

	go client.read()
	go client.write()

	// 注册用户连接
	clientManager.Connect <- client
}
