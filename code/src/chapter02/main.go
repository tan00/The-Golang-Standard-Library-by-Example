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
		fmt.Println("1 表示 stirngs 示例")
		fmt.Println("2 表示 bytes 示例")
		fmt.Println("3 表示 strconv 示例")

		fmt.Println("q 退出")
		fmt.Println("***********************************")

		var ch string
		fmt.Scanln(&ch)

		switch ch {
		case "1":
			TestStrings()
		case "2":
			TestBytes()
		case "3":
			TestStrConv()

		case "q":
			fmt.Println("程序退出！")
			break MAINFOR
		default:
			fmt.Println("输入错误！")
			continue
		}
	}
}
