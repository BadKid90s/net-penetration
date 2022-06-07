package helper

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net"
	"net-penetration/conf"
	"net-penetration/define"
	"time"
)

func CreateListen(addr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	return tcpListener, err

}

func CreateConnect(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	return tcpConn, err
}

func KeepAlive(conn net.Conn) {
	for true {
		_, err := conn.Write([]byte(define.KeepAliveStr))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}

func KeepAliveReply(conn net.Conn) {
	for true {
		_, err := conn.Write([]byte(define.KeepAliveStr))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}

// GetServerConf 解析server.yml
func GetServerConf() (*conf.Server, error) {
	s := new(conf.Server)
	b, err := ioutil.ReadFile("./conf/server.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return s, nil

}
