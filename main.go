package main

func main() {
	//创建区块链
	blockChain := NewBlockChain()
	blockChain.AddBlock("1111111111")
	blockChain.AddBlock("2222222222")
	blockChain.AddBlock("3333333333")
	//for i,block := range blockChain.blocks{
	//	fmt.Printf("========当前区块高度:%d========\n",i)
	//	fmt.Printf("前区块哈希：%x\n",block.PrevHash)
	//	fmt.Printf("当前区块哈希：%x\n",block.Hash)
	//	fmt.Printf("当前数据：%s\n",block.Data)
	//}
}
