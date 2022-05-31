package main

import (
	"bufio"
	"io"
	"log"
	"net-penetration/define"
	"net-penetration/helper"
)

func main() {
	//连接服务端控制中心
	connect, err := helper.CreateConnect(define.ControlServerAddress)
	if err != nil {
		log.Printf("连接错误 %v", err)
		panic(err)
	}
	log.Printf("[连接成功] :%v", connect.RemoteAddr().String())
	for true {
		data, err := bufio.NewReader(connect).ReadString('\n')
		if err != nil {
			log.Printf("获取数据错误 %v", err)
			continue
		}
		if string(data) == define.NewConnectionStr {
			go messageForward()
		}
	}
}

func messageForward() {
	//连接服务端隧道
	tunnelConn, err := helper.CreateConnect(define.TunnelServerAddress)
	if err != nil {
		log.Printf("连接错误 %v", err)
		panic(err)
	}
	//连接客户端服务
	localConn, err := helper.CreateConnect(define.LocalServerAddress)
	if err != nil {
		log.Printf("连接错误 %v", err)
		panic(err)
	}

	go io.Copy(localConn, tunnelConn)
	go io.Copy(tunnelConn, localConn)
}
