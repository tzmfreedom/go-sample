package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
	"net"
	"net/rpc"
	"os/exec"
	"time"
)

//func main() {
//	c := goridge_rpc.NewClientCodec()
//	err := c.WriteRequest(&rpc.Request{
//		ServiceMethod: "Test",
//	}, "hoge")
//	if err != nil {
//		panic(err)
//	}
//}

type App struct{}

func (s *App) Hi(name string, r *string) error {
	*r = fmt.Sprintf("Hello, %s!", name)
	return nil
}

type Client struct{}

func (c *Client) Call(data interface{}) ([]byte, error) {
	conn, err := net.Dial("unix", "test.sock")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write(buf)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte("\n"))
	if err != nil {
		return nil, err
	}
	buf, _, err = bufio.NewReader(conn).ReadLine()
	if err != nil {
		return nil, err
	}
	return buf, err
}

type PHPServer struct{}

func (s *PHPServer) Serve() func() {
	cmd := exec.Command("php", "php/server.php")
	go func() {
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	return func () {
		err := cmd.Process.Kill()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if true {
		closeFunc := (&PHPServer{}).Serve()
		defer closeFunc()
		c := &Client{}
		buf, err := c.Call(map[string]string{"hoge":"fuga"})
		if err != nil {
			panic(err)
		}
		var res map[string]string
		err = json.Unmarshal(buf, &res)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
		return
	}
	ln, err := net.Listen("tcp", ":6001")
	if err != nil {
		panic(err)
	}

	_ = rpc.Register(new(App))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		_ = conn
		go rpc.ServeCodec(goridgeRpc.NewCodec(conn))
	}
}
