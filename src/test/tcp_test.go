package test

import (
	"net"
	"testing"
)

const (
	addr    = "0.0.0.0:22222"
	bufSize = 10
)

//服务端
func TestTcpListen(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	listenTCP, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	for true {
		tcpConn, err := listenTCP.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		for true {
			var buf [bufSize]byte
			read, err := tcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(buf[:read]))
			if read < bufSize {
				break
			}
		}
		_, err = tcpConn.Write([]byte("server ==> hello word 你好世界"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

//客户端
func TestCreateTcp(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	_, err = tcpConn.Write([]byte("client ==》 hello word 你好世界"))
	if err != nil {
		t.Fatal(err)
	}

	for true {
		var buf [bufSize]byte
		read, err := tcpConn.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(buf[:read]))
	}
}
