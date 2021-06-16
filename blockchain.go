package main

//定义区块链结构
type BlcokChain struct{
	//定义一个区块链数组
	blocks []*Block
}

//创建一个区块链
func NewBlockChain() *BlcokChain{
	//创建创世区块并加到链上
	genesisBlock := GenesisBlock()
	return &BlcokChain{
		blocks: []*Block{genesisBlock},
	}
}

//创建初始区块
func GenesisBlock() *Block{
	block := NewBlock("创世区块",[]byte{})
	return block
}

//区块链添加区块方法
func (blockChain  *BlcokChain)AddBlock(data string){
	//取最后一个区块哈希
	lastBlock := blockChain.blocks[len(blockChain.blocks)-1]
	prevHash := lastBlock.Hash
	//创建区块
	block := NewBlock(data,prevHash)
	//将新区块加入链
	blockChain.blocks = append(blockChain.blocks,block)
}


