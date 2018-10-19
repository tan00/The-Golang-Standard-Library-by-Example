package main

import (
	"bufio"
	"fmt"
	"strings"
)

//ReadSlice 读取时有一个curpos指针, 读取后对象的值被改变了
func testReadSlice() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}
