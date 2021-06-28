package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//定义一个结构Person
type Person struct {
	Name string
	Age uint
}

func main() {
	//使用gob进行序列化（编码）得到字节流
	var xiaoMing Person
	xiaoMing.Name = "小明"
	xiaoMing.Age = 20
	//编码的数据放到buffer里面
	var buffer bytes.Buffer
	//使用gob进行序列化
	//定义一个编码器
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoMing)
	if err != nil{
		log.Panic("编码出错")
	}
	fmt.Printf("编码后的小明： %v\n",buffer.Bytes())
	//使用gob进行反序列化
	//定义一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	var daMing Person
	err = decoder.Decode(&daMing)
	if err != nil{
		log.Panic("解码出错")
	}
	fmt.Printf("解码后的小明：%v\n",daMing)
}
