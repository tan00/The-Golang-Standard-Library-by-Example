package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isAscii(r rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, r)
}
func isHan(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

func TestStrings() {
	var (
		str = "some test string\nsome test string 你好 \t 你好"
	)

	//长度
	fmt.Println("字符串长度为\t", len(str))

	//子串包含计数
	fmt.Println("包含子串计数Count\t", strings.Count(str, "你好"))

	//包含指定子串
	fmt.Println("包含子串Contains \t", strings.Contains("failure", "re"))

	//包含任意字符
	fmt.Println("是否包含任意指定字符 ContainsAny\t", strings.ContainsAny("failure", "ui"))

	//字符串分割
	fields := strings.Fields(str)
	fmt.Printf("字符串分割  %v\n", fields)
	//判断是否为汉字
	fmt.Println("龙是否为汉字 %v", isHan('龙'), "a是否为汉字 %v", isHan('a'))

	//Join
	sliceStr := []string{"this", "is", "a", "book"}
	fmt.Println("Join函数", strings.Join(sliceStr, ";"))

}
