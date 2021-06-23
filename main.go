package main

func main() {
	//创建区块链
	blockChain := NewBlockChain()
	blockChain.AddBlock("李晓德")
	blockChain.AddBlock("李晓晓")
	//for i,block := range blockChain.blocks{
	//	fmt.Printf("========当前区块高度:%d========\n",i)
	//	fmt.Printf("前区块哈希：%x\n",block.PrevHash)
	//	fmt.Printf("当前区块哈希：%x\n",block.Hash)
	//	fmt.Printf("当前数据：%s\n",block.Data)
	//}
}
