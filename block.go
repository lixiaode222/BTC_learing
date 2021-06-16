package main

import "crypto/sha256"

//定义区块结构
type Block struct {
	//前区块哈希
	PrevHash []byte
	//当前区块哈希
	Hash []byte
	//数据
	Data []byte
}

//创建区块的函数
func NewBlock(data string,prevBlockHash []byte) *Block{
	block := Block{
		PrevHash: prevBlockHash,
		Hash: []byte{},//先填空后面计算 //TODO
		Data: []byte(data),
	}
	block.setHash();
	return &block
}

//计算区块哈希函数
func (blcok *Block)setHash(){
	//拼装数据
	blockInfo := append(blcok.PrevHash,blcok.Data...)
	//计算sha25
	hash := sha256.Sum256(blockInfo);
	blcok.Hash = hash[:]
}