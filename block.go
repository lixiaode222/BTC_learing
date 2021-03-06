package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
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
	//block.setHash();
	pow := NewProofOfWork(&block)
	//挖矿
	hash,nonce := pow.Run()
	//根据挖矿结果对数据进行赋值
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//计算区块哈希函数
func (block *Block)setHash(){
	var blockInfo []byte
	//拼装数据
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Diffculty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockInfo = bytes.Join(tmp,[]byte{})
	//计算sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//序列化
func (block *Block) Serialize() []byte{
	var buffer bytes.Buffer
	//使用gob进行序列化
	//定义一个编码器
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil{
		log.Panic("序列化出错")
	}
	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) Block{
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic("反序列化出错")
	}
	return block
}