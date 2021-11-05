package main

import (
	"bufio"
	"fmt"
	"github.com/kid1999/simple_pbft_blockchain/blockchain"
	"github.com/kid1999/simple_pbft_blockchain/conf"
	. "github.com/kid1999/simple_pbft_blockchain/pbft"
	"log"
	"os"
	"strings"
)

func main() {
	//为四个节点生成公私钥
	GenRsaKeys()

	// 获取配置
	conf.NewConfig()

	if len(os.Args) != 2 {
		log.Panic("输入的参数有误！")
	}
	nodeID := os.Args[1]

	// 启动区块链，监听trans消息 发送给leader
	blockchain.Start(nodeID, conf.GlobalConfig.LeaderID)

	// 启动pbft
	go ClientSendMessageAndListen(conf.GetNode(nodeID).ClientAddr)
	p := NewPBFT(nodeID, conf.GetNode(nodeID).NodeAddr)
	go p.TcpListen() //启动节点
	if nodeID == conf.GlobalConfig.LeaderID {
		go p.BroadcastBlock()
	}

	// 读取数据发送消息
	fmt.Println("请启动全部节点后再发送消息！")
	stdReader := bufio.NewReader(os.Stdin)
	for {
		data, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}
		// create a transaction to the queue
		context := strings.TrimSpace(data)

		// send str to blockchain
		blockchain.Core.Network.IncomingMessages <- []byte(context)
	}

	select {}
}
