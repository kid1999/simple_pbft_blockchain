package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pbft_blockchain/blockchain"
	. "pbft_blockchain/pbft"
	"strings"
)

func main() {
	//为四个节点生成公私钥
	GenRsaKeys()

	if len(os.Args) != 2 {
		log.Panic("输入的参数有误！")
	}
	nodeID := os.Args[1]

	// 启动区块链，监听trans消息 发送给leader
	blockchain.Start()

	// 启动pbft
	if addr, ok := NodeTable[nodeID]; ok {
		go ClientSendMessageAndListen(ClientTable[nodeID])
		p := NewPBFT(nodeID, addr)
		go p.TcpListen() //启动节点
	} else {
		log.Fatal("无此节点编号！")
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
