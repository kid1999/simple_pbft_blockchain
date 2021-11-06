package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/kid1999/simple_pbft_blockchain/conf"
	"log"
	"math/rand"
	"time"
)

/**
* @Description: 持久化区块
* @author : kid1999
* @date Date : 2021/11/05 10:47 PM
* @version V1.0
 */

var bkName = "test"

func DB_init() *bolt.DB {
	db, err := bolt.Open(conf.GlobalConfig.DBPath+RandString(10)+".db", 0600, nil)
	if err != nil {
		log.Panicf("open the Dbfailed! %v\n", err)
	}
	return db
}

func (bc *Blockchain) StoreBlock(block Block) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bkName))
		if b == nil {
			_, err := tx.CreateBucket([]byte(bkName))
			if err != nil {
				log.Fatal(err)
			}
		}
		if b != nil {
			err := b.Put(block.Hash(), block.Serialize())
			if nil != err {
				log.Panicf("put the data of new block into Db failed! %v\n", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panicf("store the block failed! %v\n", err)
	}
}

// TODO 区块数据的读写
func (bc *Blockchain) PrintChain() {
	fmt.Println("——————————————打印区块链———————————————————————")
	var prevBlockHash = bc.CurrentBlock.PrevBlock
	for i := 1; i < int(bc.Height)-1; i++ {
		fmt.Println("—————————————————————————————————————————————")
		bc.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bkName))
			if b != nil {
				println("get: ", prevBlockHash)
				blockBytes := b.Get(prevBlockHash)
				curBlock := DeserilizeBlock(blockBytes)
				fmt.Println(curBlock)
				prevBlockHash = curBlock.PrevBlock
			}
			return nil
		})
	}
}

// 序列化
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panicf("serialize the block to byte failed %v \n", err)
	}
	return result.Bytes()
}

// 反序列化
func DeserilizeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panicf("deserialize the block to byte failed %v \n", err)
	}
	return &block
}

// 随机命名数据库文件
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
