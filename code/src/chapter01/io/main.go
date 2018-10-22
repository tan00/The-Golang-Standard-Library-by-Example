package main

import (
	"fmt"

	"util" //export GOPATH="/d/1src/go/The-Golang-Standard-Library-by-Example/code"
)

func main() {
	util.Welcome()
	MainMenu()
}

func MainMenu() {
MAINFOR:
	for {
		fmt.Println("")
		fmt.Println("*******请选择示例：*********")
		fmt.Println("1 表示 io.Reader 示例")
		fmt.Println("2 表示 io.ByteReader/ByteWriter 示例")
		fmt.Println("3 表示 ReadDir 示例")
		fmt.Println("4 表示 Formatter接口")
		fmt.Println("5 表示 Scanf")
		fmt.Println("6 表示 bufio.ReadSlice")
		fmt.Println("7 表示 bufio.ReadString")

		fmt.Println("q 退出")
		fmt.Println("***********************************")

		var ch string
		fmt.Scanln(&ch)

		switch ch {
		case "1":
			ReaderExample()
		case "2":
			ByteRWerExample()
		case "3":
			fmt.Println("在当前目录遍历, 遍历深度为1")
			ListAllFile(".", 1)
		case "4":
			sb := Person{"james", 18, 1}
			fmt.Printf("Person %%L is %L \n  Person %%s is %s   Person %%v is %v  ", sb, sb, sb)
		case "5":
			testScanf()
		case "6":
			testReadSlice()
		case "7":
			testReadBytes()

		case "q":
			fmt.Println("程序退出！")
			break MAINFOR
		default:
			fmt.Println("输入错误！")
			continue
		}
	}
}
