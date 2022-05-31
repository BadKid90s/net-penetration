package helper

import (
	"log"
	"net"
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

func GetDataFromConnect(bufSize int, conn net.Conn) ([]byte, error) {
	b := make([]byte, 0)
	for true {
		var buf = make([]byte, bufSize)
		read, err := conn.Read(buf[:])
		if err != nil {
			return nil, err
		}
		b = append(b, buf[:read]...)
		if read < bufSize {
			break
		}
	}
	return b, nil
}
