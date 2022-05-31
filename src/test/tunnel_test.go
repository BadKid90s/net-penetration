package test

import (
	"io"
	"net"
	"strconv"
	"testing"
)

const (
	serverAddr = "0.0.0.0:23222"
	tunnelAddr = "0.0.0.0:24222"
	BufSize    = 10
)

//server
func TestServer(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		t.Fatal(err)
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	for true {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		var b = make([]byte, 0)
		for true {
			var buf [BufSize]byte
			read, err := tcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			b = append(b, buf[:read]...)

			if read < BufSize {
				break
			}
		}
		t.Log("服务端接收：" + string(b))
		atoi, err := strconv.Atoi(string(b))
		if err != nil {
			t.Fatal(err)
		}
		atoi = atoi + 2

		t.Log("服务端发送：" + string(strconv.Itoa(atoi)))
		_, err = tcpConn.Write([]byte(strconv.Itoa(atoi)))
		if err != nil {
			t.Fatal(err)
		}
	}

}

//client
func TestClient(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tcpConn.Write([]byte("1200"))
	t.Log("客户端发送：" + "1200")
	if err != nil {
		t.Fatal(err)
	}
	var b = make([]byte, 0)
	for true {
		var buf [BufSize]byte
		read, err := tcpConn.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}
		b = append(b, buf[:read]...)
		if read < BufSize {
			break
		}
	}
	t.Log("客户端接收：" + string(b))
	t.Log(string(b))

}

//tunnel
func TestTunnel(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	for true {
		// 客户端连接
		clientTcpConnect, err := tcpListener.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}
		////获取客户端传递的数据
		//var b = make([]byte, 0)
		//for true {
		//	var buf [BufSize]byte
		//	read, err := clientTcpConnect.Read(buf[:])
		//	if err != nil {
		//		t.Fatal(err)
		//	}
		//	b = append(b, buf[:read]...)
		//	if read < BufSize {
		//		break
		//	}
		//}
		//与服务端创建连接
		serverTCPAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		serverTcpConn, err := net.DialTCP("tcp", nil, serverTCPAddr)
		if err != nil {
			t.Fatal(err)
		}
		//t.Log("通道接收：" + string(b))
		//serverTcpConn.Write(b)
		//
		////获取服务端响应的数据
		//var b2 = make([]byte, 0)
		//for true {
		//	var buf [BufSize]byte
		//	read, err := serverTcpConn.Read(buf[:])
		//	if err != nil {
		//		t.Fatal(err)
		//	}
		//	b2 = append(b2, buf[:read]...)
		//	if read < BufSize {
		//		break
		//	}
		//}
		//t.Log("通道发送：" + string(b2))
		//clientTcpConnect.Write(b2)

		go io.Copy(clientTcpConnect, serverTcpConn)
		go io.Copy(serverTcpConn, clientTcpConnect)
	}
}
