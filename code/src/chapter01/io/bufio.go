package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

//bufio.ReadSlice  读取后对象的值被改变了,ReadSlice返回的[]byte是指向Reader中的buffer，而不是copy一份返回
//正因为ReadSlice返回的数据会被下次的I/O操作重写，因此许多的客户端会选择使用ReadBytes或者ReadString来代替
func testReadSlice() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

//bufio.ReadBytes() 内部调用ReadSlice, 并发读取可能会有问题. 下面进行尝试
func testReadBytes() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	go func() {
		line, _ := reader.ReadString('\n')
		fmt.Println("line1" + line)
	}()
	go func() {
		line, _ := reader.ReadString('\n')
		fmt.Println("line2" + line)
	}()
	go func() {
		line, _ := reader.ReadString('\n')
		fmt.Println("line3" + line)
	}()

	time.Sleep(2000)
}
