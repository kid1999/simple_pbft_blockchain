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

// 封装为一个Transaction消息发送
func ClientSendMessageAndListen(addr string) {
	//开启客户端的本地监听（主要用来接收节点的reply信息）
	go ClientTcpListen(addr)
	fmt.Printf("客户端开启监听，地址：%s\n", addr)

	for {
		select {
		case transaction := <-blockchain.Core.Network.BroadcastQueue:
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
			tcpDial(content, NodeTable[LeaderID])
			// TODO: 消息失败重启
		}
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
