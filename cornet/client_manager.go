package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// 连接管理
type ClientManager struct {
	Clients *ConcurrentMap // 全部的连接
	//ClientsLock sync.RWMutex       // 读写锁
	Users      map[string]*Client // 登录的用户 // appId+uuid
	UserLock   sync.RWMutex       // 读写锁
	Connect    chan *Client       // 连接连接处理
	DisConnect chan *Client       // 断开连接处理程序
	Broadcast  chan []byte        // 广播 向全部成员发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	concurrentMap := NewConcurrentMap()
	clientManager = &ClientManager{
		Clients:    &concurrentMap,
		Users:      make(map[string]*Client),
		Connect:    make(chan *Client, 1000),
		DisConnect: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
	return
}

func (manager *ClientManager) addClient(client *Client) {
	manager.Clients.Set(client.SignId, client)
}

func (manager *ClientManager) removeClient(client *Client) {
	manager.Clients.Remove(client.SignId)
}

func (manager *ClientManager) getClient(signId string) (client *Client) {
	if obj, ex := manager.Clients.Get(signId); ex {
		client = obj.(*Client)
	}
	return
}

func (manager *ClientManager) handleConnect(client *Client) {
	manager.addClient(client)
	logrus.WithFields(logrus.Fields{
		"appId":      client.AppId,
		"userId":     client.UserId,
		"remoteAddr": client.ClientAddr,
		"clientId":   client.RpcServerAddr,
	}).Info("客户端已连接")
}

func (manager *ClientManager) handleDisConnect(client *Client) {
	manager.removeClient(client)
	logrus.WithFields(logrus.Fields{
		"userId":     client.UserId,
		"appId":      client.AppId,
		"remoteAddr": client.ClientAddr,
		"clientId":   client.RpcServerAddr,
	}).Info("客户端已断开")
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Connect:
			// 建立链接
			manager.handleConnect(conn)
		case conn := <-manager.DisConnect:
			// 断开链接
			manager.handleDisConnect(conn)

		}
	}
}
