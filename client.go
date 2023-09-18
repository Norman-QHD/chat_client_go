package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

// 处理server回应的消息,直接显示到标准输出即可
func (client *Client) DealResponse() {
	//等同于创建buf,然后读取buf,再打印到页面上.
	//一旦client.conn有数据,就直接copy到stdout标准输出上,并且永久阻塞监听
	i, err := io.Copy(os.Stdout, client.conn)
	if err != nil {
		fmt.Println("读取消息失败:", err)
		fmt.Println("读取消息返回结果:", i)
		return
	}

	/*for {
	buf:
		make()
	client.conn.Read(buf)
	fmt.Printf(buf)
	}*/
}

func (client *Client) UpdateName() bool {
	fmt.Println("请输入用户名")
	scanned, err := fmt.Scanln(&client.Name)
	if err != nil {
		fmt.Println("获取用户输入错误:", scanned)
		return false
	}
	sendMsg := "rename|" + client.Name + "\n"
	_, err = client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
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
			client.UpdateName()
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

	//单独开启一个goroutine 处理server回传的消息
	go client.DealResponse()

	//启动客户端业务
	client.Run()
}
