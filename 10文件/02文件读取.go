package main

import (
	"fmt"
	"io/ioutil"
)

/**
bufio的写操作
bufio封装了io.Reader、io.Writer接口对象，并创建了另一个也实现了该接口的对象：bufio.Reader、bufio.Writer。通过该实现，bufio实现了文件的缓冲区设计，可以大大提高文件I/O的效率。

使用bufio读取文件时，先将数据读入内存的缓冲区（缓冲区一般比要比程序中设置的文件接收对象要大），这样就可以有效降低直接I/O的次数。

bufio.Read([]byte)相当于读取大小len(p)的内容：

当缓冲区有内容时，将缓冲区内容全部填入p并清空缓冲区
当缓冲区没有内容且len(p)>len(buf)，即要读取的内容比缓冲区还要大，直接去文件读取即可
当缓冲区没有内容且len(p)<len(buf)，即要读取的内容比缓冲区小，读取文件内容并填满缓冲区，并将p填满
以后再次读取时，缓冲区有内容，将缓冲区内容全部填入p并清空缓冲区（和第一步一致）
 */
func main() {
reads()

}
func reads(){
	file, err := ioutil.ReadFile("test.txt")
	if err!=nil{
		fmt.Println("文件读取失败")
	}
	fmt.Printf("%s",file)

}
func writes()  {
	s:="伟大的无产阶级政党"
	err:= ioutil.WriteFile("test.txt", []byte(s), 6)
	if err!=nil{
		fmt.Println(err)
	}

}