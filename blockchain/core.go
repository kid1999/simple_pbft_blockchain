package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/kid1999/simple_pbft_blockchain/conf"
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
	NodeID            string
	LeaderId          string
	BroadcastQueue    chan Message
	IncomingMessages  chan Message
	TransactionsQueue chan *Transaction
	BlockQueue        chan Block
	ReceivedMessages  chan Message
}

var Core = struct {
	*Keypair
	*Blockchain
	*Network
}{}

func SetupNetwork() *Network {
	n := new(Network)
	n.BroadcastQueue, n.IncomingMessages = make(chan Message), make(chan Message)
	n.TransactionsQueue, n.BlockQueue = make(chan *Transaction, conf.GlobalConfig.BlockSize*2), make(chan Block)
	n.ReceivedMessages = make(chan Message)
	return n
}

func Start(id string, LeaderId string) {
	// Generating keypair
	fmt.Println("Generating keypair...")
	keypair := GenerateNewKeypair()
	Core.Keypair = keypair

	// Setup Network by pbft
	Core.Network = SetupNetwork()
	Core.Network.NodeID = id
	Core.Network.LeaderId = LeaderId

	// Setup blockchain
	Core.Blockchain = SetupBlockchan()

	// 只有Leader Run
	if id == LeaderId {
		go Core.Blockchain.LeaderRun()
	}
	// follower 只接受Block 加入区块
	go Core.Blockchain.Run()

	// 接受来自terminal的消息
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

func CreateTransaction(to, txt string) *Transaction {
	t := NewTransaction(Core.Keypair.Public, []byte(to), []byte(txt))
	t.Signature = t.Sign(Core.Keypair)
	return t
}

// 处理消息
func HandleIncomingMessage(txt string) {
	t := CreateTransaction(Core.NodeID, txt)
	// send block to leader
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	Core.Network.BroadcastQueue <- bytes
}
