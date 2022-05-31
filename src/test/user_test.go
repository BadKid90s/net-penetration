package test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net-penetration/define"
	"net-penetration/helper"
	"sync"
	"testing"
)

var wg sync.WaitGroup
var clientConnect net.Conn

//服务端
func TestUserServer(t *testing.T) {
	wg.Add(1)
	//监听控制中心
	go ControlServer()
	//监听用户请求
	go RequestServer()
	wg.Wait()
}

func ControlServer() {
	tcpListener, err := helper.CreateListen(define.ControlServerAddress)
	if err != nil {
		panic(err)
	}
	for true {
		clientConnect, err = tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		// 心跳机制
		go helper.KeepAlive(clientConnect)

	}
}

func RequestServer() {
	tcpListener, err := helper.CreateListen(define.RequestServerAddress)
	if err != nil {
		panic(err)
	}
	for true {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go io.Copy(clientConnect, tcpConn)
		go io.Copy(tcpConn, clientConnect)

	}
}

//客户端
func TestUserClient(t *testing.T) {
	tcpConn, err := helper.CreateConnect(define.ControlServerAddress)
	if err != nil {
		log.Printf("[连接失败] %s", err)
	}
	for true {
		readString, err := bufio.NewReader(tcpConn).ReadString('\n')
		if err != nil {
			log.Printf("Get Data Error: %v", err)
			continue
		}
		log.Printf("Get Data: %v", readString)
		_, err = tcpConn.Write([]byte("I Get\n"))
		if err != nil {
			log.Printf("Send Data Error:%v", err)
		}
	}
}

//用户端
func TestUserRequest(t *testing.T) {
	conn, err := helper.CreateConnect(define.RequestServerAddress)
	if err != nil {
		log.Printf("[连接失败] %s", err)
	}
	_, err = conn.Write([]byte("中文 \n"))
	if err != nil {
		log.Printf("[发送失败] %s", err)
	}
	s, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("[接收失败] %s", err)
	}
	fmt.Println(s)

}
