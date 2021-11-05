package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/kid1999/simple_pbft_blockchain/conf"
	"time"
)

/**
* @Description: blockchain
* @author : kid1999
* @date Date : 2021/10/21 15:51
* @version V1.0
 */

type TransactionsQueue chan *Transaction
type BlocksQueue chan Block

type Blockchain struct {
	CurrentBlock Block
	BlockSlice

	TransactionsQueue
	BlocksQueue
}

func SetupBlockchan() *Blockchain {

	bl := new(Blockchain)
	bl.TransactionsQueue, bl.BlocksQueue = make(TransactionsQueue), make(BlocksQueue)

	//Read blockchain from file and stuff...

	bl.CurrentBlock = bl.CreateNewBlock()

	return bl
}

func (bl *Blockchain) CreateNewBlock() Block {

	prevBlock := bl.BlockSlice.PreviousBlock()
	prevBlockHash := []byte{}
	if prevBlock != nil {

		prevBlockHash = prevBlock.Hash()
	}

	b := NewBlock(prevBlockHash)
	b.BlockHeader.Origin = Core.Keypair.Public

	return b
}

func (bl *Blockchain) AddBlock(b Block) {
	bl.BlockSlice = append(bl.BlockSlice, b)
	fmt.Println("Height: ", len(bl.BlockSlice))
}

// Leader 在这里生产区块
func (bl *Blockchain) LeaderRun() {
	for {
		select {
		// 到达出块时间 出块
		case <-time.After(time.Second * time.Duration(conf.GlobalConfig.BlockInterval)):
			block := bl.CurrentBlock
			block.BlockHeader.MerkelRoot = block.GenerateMerkelRoot()
			block.BlockHeader.Nonce = 0
			block.BlockHeader.Timestamp = uint32(time.Now().Unix())
			// TODO 验证区块信息
			fmt.Println("New block!", block.Hash())
			// broadcast the block
			Core.Network.BlockQueue <- block

		// 到达区块大小 出块
		case tr := <-Core.Network.TransactionsQueue:
			// Put currentBlock into chan block
			bl.CurrentBlock.AddTransaction(tr)
			fmt.Println("data: ", string(tr.Payload))

			// Adjustment of the block size or block interval
			if bl.CurrentBlock.TransactionSlice.Len() >= conf.GlobalConfig.BlockSize {
				block := bl.CurrentBlock
				block.BlockHeader.MerkelRoot = block.GenerateMerkelRoot()
				block.BlockHeader.Nonce = 0
				block.BlockHeader.Timestamp = uint32(time.Now().Unix())
				// TODO 验证区块信息
				fmt.Println("New block!", block.Hash())
				// broadcast the block
				Core.Network.BlockQueue <- block
				//New Block
				bl.CurrentBlock = bl.CreateNewBlock()
			}
		}
	}
}

func (bl *Blockchain) Run() {
	for {
		select {
		case <-time.After(time.Second):
		case b := <-Core.Network.ReceivedMessages:
			var block Block
			err := json.Unmarshal(b, &block)
			if err == nil {
				bl.AddBlock(block)
				//New Block
				bl.CurrentBlock = bl.CreateNewBlock()
			}
		}
	}
}
