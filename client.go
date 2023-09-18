package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       -1,
	}
	//连接server
	conn, err := net.Dial(
		"tcp",
		fmt.Sprintf("%s:%d", serverIp, serverPort),
	)
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	//返回对象
	return client
}

func (client *Client) menu() bool {
	var inputFlag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("3.退出")

	scan, err := fmt.Scanln(&inputFlag)
	if err != nil {
		fmt.Println("获取用户的输入错误:", scan)
		return false
	}

	if inputFlag >= 0 && inputFlag <= 3 {
		client.flag = inputFlag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			//公聊模式
			fmt.Println("公聊模式")
			break
		case 2:
			//私聊模式
			fmt.Println("私聊模式")
			break
		case 3:
			//修改用户名
			fmt.Println("修改用户名")
			break
		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>连接服务端失败...")
		return
	}
	fmt.Println(">>>>>服务器连接成功!!!")

	//启动客户端业务
	client.Run()
}
