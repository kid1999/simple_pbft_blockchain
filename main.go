package main

import (
	"log"
	"os"
	. "pbft_blockchain/pbft"
)

func main() {
	//为四个节点生成公私钥
	GenRsaKeys()

	if len(os.Args) != 2 {
		log.Panic("输入的参数有误！")
	}
	nodeID := os.Args[1]
	if nodeID == "client" {
		ClientSendMessageAndListen() //启动客户端程序
	} else if addr, ok := NodeTable[nodeID]; ok {
		p := NewPBFT(nodeID, addr)
		go p.TcpListen() //启动节点
	} else {
		log.Fatal("无此节点编号！")
	}
	select {}
}
