package blockchain

import (
	"encoding/json"
	"fmt"
	"log"
)

/**
* @Description: block chain core
* @author : kid1999
* @date Date : 2021/10/21 16:17
* @version V1.0
 */

type Message []byte

// 统一让 PBFT 处理
type Network struct {
	BroadcastQueue    chan Message
	IncomingMessages  chan Message
	TransactionsQueue chan *Transaction
	BlockQueue        chan Block
}

var Core = struct {
	*Keypair
	*Blockchain
	*Network
	TransactionCount int
}{}

func SetupNetwork() *Network {
	n := new(Network)
	n.BroadcastQueue, n.IncomingMessages = make(chan Message), make(chan Message)
	n.TransactionsQueue, n.BlockQueue = make(chan *Transaction), make(chan Block)

	return n
}

func Start() {
	// Generating keypair
	fmt.Println("Generating keypair...")
	keypair := GenerateNewKeypair()
	Core.Keypair = keypair

	// Setup Network by pbft
	Core.Network = SetupNetwork()

	// Setup blockchain
	Core.Blockchain = SetupBlockchan()
	go Core.Blockchain.Run()

	// TODO 处理消息，节点在pbft中验证block消息
	go func() {
		for {
			select {
			case msg := <-Core.Network.IncomingMessages:
				// 注意此处是已经验证过的消息
				HandleIncomingMessage(string(msg))
			}
		}
	}()
}

// TODO 创建交易信息
func CreateTransaction(txt string) *Transaction {

	t := NewTransaction(Core.Keypair.Public, nil, []byte(txt))
	t.Signature = t.Sign(Core.Keypair)

	return t
}

// TODO 处理消息
func HandleIncomingMessage(txt string) {
	t := CreateTransaction(txt)
	// send block to leader
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	Core.Network.BroadcastQueue <- bytes
}

func logOnError(err error) {

	if err != nil {
		log.Println("[Todos] Err:", err)
	}
}
