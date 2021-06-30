package main

import "fmt"

func main() {
	//创建区块链
	blockChain := NewBlockChain()
	blockChain.AddBlock("1111111111")
	blockChain.AddBlock("2222222222")
	blockChain.AddBlock("3333333333")
	it := blockChain.NewIterator()
	//调用迭代器
	for {
		//得到block
		block := it.Next()
		fmt.Printf("========================================\n\n")
		fmt.Printf("前区块哈希：%x\n",block.PrevHash)
		fmt.Printf("当前区块哈希：%x\n",block.Hash)
		fmt.Printf("当前数据：%s\n",block.Data)
		if len(block.PrevHash) == 0{
			fmt.Printf("区块链遍历结束！！！！！！！！")
			break
		}
	}
}
