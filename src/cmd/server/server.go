package main

import (
	"io"
	"log"
	"net"
	"net-penetration/define"
	"net-penetration/helper"
	"sync"
)

var (
	controlConnect     *net.TCPConn
	userRequestConnect *net.TCPConn
	wg                 sync.WaitGroup
)

func main() {
	wg.Add(1)
	//控制中心监听
	go controlListen()
	//用户请求监听
	go userRequestListen()
	//隧道监听
	go tunnelListen()
	wg.Wait()
}

// tunnelListen 用户请求监听
func tunnelListen() {
	listen, err := helper.CreateListen(define.TunnelServerAddress)
	if err != nil {
		panic(err)
	}
	log.Printf("[隧道] 监听中：%v \n", listen.Addr().String())
	for true {
		//用户进行连接
		tunnelConnect, err := listen.AcceptTCP()
		if err != nil {
			log.Printf("tunnelConnect AcceptTCP Error %v \n", err)
			return
		}
		//数据转发
		go io.Copy(userRequestConnect, tunnelConnect)
		go io.Copy(tunnelConnect, userRequestConnect)

	}
}

//userRequestListen 用户请求监听
func userRequestListen() {
	listen, err := helper.CreateListen(define.UserRequestAddress)
	if err != nil {
		panic(err)
	}
	log.Printf("[用户请求] 监听中：%v \n", listen.Addr().String())

	for true {
		//用户进行连接
		userRequestConnect, err = listen.AcceptTCP()
		if err != nil {
			log.Printf("userRequestListen AcceptTCP Error %v \n", err)
			return
		}
		//发送新的消息给客户端，告诉客户端有新的连接
		_, err := controlConnect.Write([]byte(define.NewConnectionStr))
		if err != nil {
			log.Printf("controlConnect Write Error %v \n", err)
			return
		}
	}
}

//controlListen 控制中心监听
func controlListen() {
	listen, err := helper.CreateListen(define.ControlServerAddress)
	if err != nil {
		panic(err)
	}
	log.Printf("[控制中心] 监听中：%v \n", listen.Addr().String())

	for true {
		//客户端连接
		controlConnect, err = listen.AcceptTCP()
		if err != nil {
			log.Printf("controlListen AcceptTCP Error %v \n", err)
			return
		}
		//进行心跳
		go helper.KeepAlive(controlConnect)
	}
}
