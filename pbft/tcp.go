package pbft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"pbft_blockchain/conf"
)

//客户端使用的tcp监听
func ClientTcpListen(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	replyCount := map[int]int{}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		b, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}
		var r Reply
		err = json.Unmarshal(b, &r)
		if err == nil {
			replyCount[r.MessageID]++
			if replyCount[r.MessageID] > conf.GlobalConfig.NodeCount/3 {
				println(r.MessageID, " has reply success!")
				replyCount[r.MessageID] = -conf.GlobalConfig.NodeCount
				// TODO: 消息失败重启
			}
		}

	}
}

//节点使用的tcp监听
func (p *pbft) TcpListen() {
	listen, err := net.Listen("tcp", p.node.addr)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("节点开启监听，地址：%s\n", p.node.addr)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		b, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}
		p.handleRequest(b)
	}

}

//使用tcp发送消息
func tcpDial(context []byte, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("connect error", err)
		return
	}

	_, err = conn.Write(context)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
