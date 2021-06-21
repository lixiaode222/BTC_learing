package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	//block
	Block *Block
	//目标值
	Target *big.Int
}

//创建POW函数
func NewProofOfWork(block *Block) *ProofOfWork{
	pow := ProofOfWork{
		Block: block,
	}
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr,16)
	pow.Target = &tmpInt
	return  &pow
}

//工作量证明
func (pow *ProofOfWork) Run() ([]byte,uint64) {
	//定义随机值变量
	var nonce uint64
	var hash [32]byte
	var tmpInt big.Int
	block := pow.Block
	for {
		//拼装数据
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Diffculty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(tmp,[]byte{})
		//进行哈希运算
		hash = sha256.Sum256(blockInfo)
		//与pow中的target进行比较
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.Target) == -1{
			fmt.Printf("挖矿成功! hash : %x，nonce : %d\n",tmpInt,nonce)
			return  hash[:],nonce
		}else{
			nonce++
		}
	}
}
