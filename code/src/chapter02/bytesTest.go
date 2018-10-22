package main

import (
	"bytes"
	"fmt"
)

func TestBytes() {
	var (
		str    = "some test string\nsome test string 你好 \t 你好"
		substr = "好"

		b  = []byte(str)
		sb = []byte(substr)
	)
	fmt.Println("bytes.Contains", bytes.Contains(b, sb))
	fmt.Println(" bytes.Count", bytes.Count(b, sb))

}
