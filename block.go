package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

//定义区块结构
type Block struct {
	//版本号
	Version uint64
	//前区块哈希
	PrevHash []byte
	//Merkel根
	MerkelRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Diffculty uint64
	//随机值
	Nonce uint64
	//当前区块哈希
	Hash []byte
	//数据
	Data []byte
}

//uint64转成btye
func Uint64ToByte(num uint64) []byte{
	var buffer bytes.Buffer
	err := binary.Write(&buffer,binary.BigEndian,num)
	if err != nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}

//创建区块的函数
func NewBlock(data string,prevBlockHash []byte) *Block{
	block := Block{
		Version: 00,
		PrevHash: prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Diffculty: 0,
		Nonce: 0,
		Hash: []byte{},//先填空后面计算 //TODO
		Data: []byte(data),
	}
	block.setHash();
	return &block
}

//计算区块哈希函数
func (blcok *Block)setHash(){
	var blockInfo []byte
	//拼装数据
	tmp := [][]byte{
		Uint64ToByte(blcok.Version),
		blcok.PrevHash,
		blcok.MerkelRoot,
		Uint64ToByte(blcok.TimeStamp),
		Uint64ToByte(blcok.Diffculty),
		Uint64ToByte(blcok.Nonce),
		blcok.Data,
	}
	blockInfo = bytes.Join(tmp,[]byte{})
	//计算sha256
	hash := sha256.Sum256(blockInfo)
	blcok.Hash = hash[:]
}