package main

import (
	"github.com/boltdb/bolt"
	"log"
)

//定义区块链结构
type BlcokChain struct{
	//使用数据库代替数组
	db *bolt.DB
	//最后一个区块的哈希
	tail []byte
}

//定义数据库名常量
const blockChainDb = "blockChain.db"
//定义Bucket名常量
const blockBucket = "blockBucket"

//创建一个区块链
func NewBlockChain() *BlcokChain{
	//定义最后一个区块哈希
	var lastHash []byte
	//打开数据库
	db,err := bolt.Open(blockChainDb,0600,nil)
	if err != nil{
		log.Panic("打开数据库失败！")
	}
	//操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屉（如果没有就创建）
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有抽屉，就去创建
			bucket,err = tx.CreateBucket([]byte(blockBucket))
			if err != nil{
				log.Panic("创建bucket失败")
			}
			//创建创世区块并加到链上
			genesisBlock := GenesisBlock()
			//将创世区块序列化后存入数据库
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"),genesisBlock.Hash)
			//写到内存里面
			lastHash = genesisBlock.Hash
		}else{
			//如果抽屉有就返回
			bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})
	return &BlcokChain{db,lastHash}
}

//创建初始区块
func GenesisBlock() *Block{
	block := NewBlock("创世区块",[]byte{})
	return block
}

//区块链添加区块方法
func (blockChain  *BlcokChain)AddBlock(data string){
	//取最后一个区块哈希
	db := blockChain.db
	lastHash := blockChain.tail
	//创建区块
	db.Update(func(tx *bolt.Tx) error {
		//完成数据的添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			log.Panic("bucket不应该为空，请检查!")
		}
		block := NewBlock(data,lastHash)
		//添加到区块链DB中
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("lastHashKey"),block.Hash)
		//更新到内存中
		blockChain.tail = block.Hash
		return nil
	})
}


