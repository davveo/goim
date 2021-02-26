package main

import (
	"time"

	"github.com/davveo/goim/cornet/lib/helper"

	"github.com/gorilla/websocket"
)

type Client struct {
	ClientAddr    string          // 客户端地址
	RpcServerAddr string          // 链接唯一标示
	Conn          *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         string          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	SignId        string          // appId+userId
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

func NewClient(remoteAddr, appId, userId string, conn *websocket.Conn) *Client {
	rpcServerAddr := helper.GenClientId()
	currentTime := uint64(time.Now().Unix())
	return &Client{
		Conn:          conn,
		AppId:         appId,
		UserId:        userId,
		ClientAddr:    remoteAddr,
		RpcServerAddr: rpcServerAddr,
		FirstTime:     currentTime,
		HeartbeatTime: currentTime,
		SignId:        appId + userId,
		Send:          make(chan []byte, 100),
	}
}

func (c *Client) read() {
	// 读取客户端数据
}

func (c *Client) write() {
	// 向客户端写数据
}
