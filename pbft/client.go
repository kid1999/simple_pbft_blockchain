package pbft

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"pbft_blockchain/blockchain"
	"time"
)

// TODO 控制此处，在论文中100个node中只有21个可以发送Transactions
//func ClientSendMessageAndListen() {
//	//开启客户端的本地监听（主要用来接收节点的reply信息）
//	go ClientTcpListen()
//	fmt.Printf("客户端开启监听，地址：%s\n", clientAddr)
//
//	fmt.Println(" ---------------------------------------------------------------------------------")
//	fmt.Println("|  已进入PBFT测试Demo客户端，请启动全部节点后再发送消息！ :)  |")
//	fmt.Println(" ---------------------------------------------------------------------------------")
//	fmt.Println("请在下方输入要存入节点的信息：")
//	//首先通过命令行获取用户输入
//	stdReader := bufio.NewReader(os.Stdin)
//	for {
//		data, err := stdReader.ReadString('\n')
//		if err != nil {
//			fmt.Println("Error reading from stdin")
//			panic(err)
//		}
//		r := new(Request)
//		r.Timestamp = time.Now().UnixNano()
//		r.ClientAddr = clientAddr
//		r.Message.ID = getRandom()
//
//		// TODO 修改为Transaction
//		//消息内容就是用户的输入
//		r.Message.Content = strings.TrimSpace(data)
//		br, err := json.Marshal(r)
//		if err != nil {
//			log.Panic(err)
//		}
//		fmt.Println(string(br))
//		content := jointMessage(cRequest, br)
//
//		// TODO 修改为当前主节点
//		//默认N0为主节点，直接把请求信息发送至N0
//		tcpDial(content, NodeTable["N0"])
//	}
//}

// TODO 封装为一个Transaction消息发送function
func ClientSendMessageAndListen(addr string) {
	//开启客户端的本地监听（主要用来接收节点的reply信息）
	go ClientTcpListen(addr)
	fmt.Printf("客户端开启监听，地址：%s\n", addr)

	for {
		transaction := <-blockchain.Core.Network.BroadcastQueue
		r := new(Request)
		r.Timestamp = time.Now().UnixNano()
		r.ClientAddr = addr
		r.Message.ID = getRandom()

		//消息内容就是用户的输入
		r.Message.Content = transaction
		br, err := json.Marshal(r)
		if err != nil {
			log.Panic(err)
		}
		content := jointMessage(cRequest, br)

		//默认N0为主节点，直接把请求信息发送至N0
		tcpDial(content, NodeTable[masterID])
	}
}

//返回一个十位数的随机数，作为msgid
func getRandom() int {
	x := big.NewInt(10000000000)
	for {
		result, err := rand.Int(rand.Reader, x)
		if err != nil {
			log.Panic(err)
		}
		if result.Int64() > 1000000000 {
			return int(result.Int64())
		}
	}
}
