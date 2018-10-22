package main

import (
	"fmt"
	"strconv"
)

func TestStrConv() {

	//strconv.Atoi("123")  //123

	i1, _ := strconv.ParseInt("128", 10, 64)
	fmt.Println("strconv.ParseInt ", i1)

	//如果字符串有前缀, 则 base必须为0 , 否则转换失败
	i2, err := strconv.ParseInt("0xff", 0, 64)
	fmt.Println("strconv.ParseInt ", i2, err)

	i3, err := strconv.ParseInt("ff", 16, 64)
	fmt.Println("strconv.ParseInt ", i3, err)
	fmt.Println("strconv.FormatInt ", strconv.FormatInt(i3, 10))

	f1, err := strconv.ParseFloat("1.19002", 64)
	fmt.Println("strconv.ParseFloat ", f1, err)

	//infinity
	f2, err := strconv.ParseFloat("inf", 64)
	fmt.Println("strconv.ParseFloat  infinity", f2, err)
	//对于 'e', 'E' 和 'f'，有效数字用于小数点之后的位数；对于 'g' 和 'G'，则是所有的有效数字
	fmt.Println("strconv.FormatFloat ", strconv.FormatFloat(f1, 'g', 5, 64))

	fmt.Println("strconv.Quote ", strconv.Quote("quote"))
}
