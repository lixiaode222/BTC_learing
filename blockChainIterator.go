package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainTterator struct {
	//区块链
	db *bolt.DB
	//当前哈希指针
	currentHashPointer []byte
}

func (blockChain *BlcokChain) NewIterator() *BlockChainTterator{
	return &BlockChainTterator{
		blockChain.db,
		blockChain.tail,
	}
}

func (it *BlockChainTterator) Next() *Block{
	var block Block
	//返回当前的区块
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			log.Panic("迭代器在遍历时bucket为空！")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		//解码动作
		block = Deserialize(blockTmp)
		//指针左移
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}













