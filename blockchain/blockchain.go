package blockchain

import (
	"fmt"
	"reflect"
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

// TODO: blockchain 从这里 拿到交易 或者 block
func (bl *Blockchain) Run() {
	for {
		select {
		// all node deal with transaction
		case tr := <-Core.Network.TransactionsQueue:

			fmt.Println("get pbft trans: ", string(tr.Payload))

			// TODO 验证交易信息

			// Put currentBlock into chan block
			bl.CurrentBlock.AddTransaction(tr)
			fmt.Println("data: ", string(tr.Payload))

			//TODO: Make the Block include N transactions
			if bl.CurrentBlock.TransactionSlice.Len() >= BLOCK_SIZE {
				block := bl.CurrentBlock
				fmt.Println("Add a block in Queue.")
				block.BlockHeader.MerkelRoot = block.GenerateMerkelRoot()
				block.BlockHeader.Nonce = 0
				block.BlockHeader.Timestamp = uint32(time.Now().Unix())
				// TODO 验证区块信息
				fmt.Println("New block!", block.Hash())
				// broadcast the block
				Core.Network.BlockQueue <- block
				// Add block into blockchain
				bl.AddBlock(block)
				//New Block
				bl.CurrentBlock = bl.CreateNewBlock()
			}
		}
	}
}

func DiffTransactionSlices(a, b TransactionSlice) (diff TransactionSlice) {
	//Assumes transaction arrays are sorted (which maybe is too big of an assumption)
	lastj := 0
	for _, t := range a {
		found := false
		for j := lastj; j < len(b); j++ {
			if reflect.DeepEqual(b[j].Signature, t.Signature) {
				found = true
				lastj = j
				break
			}
		}
		if !found {
			diff = append(diff, t)
		}
	}

	return
}
