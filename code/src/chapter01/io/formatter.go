package main

import (
	"fmt"
)

//一个类型只要有 String() string 方法，我们就说它实现了 Stringer 接口。

//实现 Format(f State, c rune)方法, 既实现了Formatter 接口
//实现了Formatter 会在屏蔽 其他格式化输出. 即任何%s %v 这样的输出都为空, 除非在Format 重新定义

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (p Person) String() string {
	return fmt.Sprintf("From String() person=%s", p.Name)
}

func (p Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		str := fmt.Sprintf("name=%s age=%d ", p.Name, p.Age)
		if p.Sex == 1 {
			str += "sex=male"
		} else {
			str += "sex=female"
		}
		f.Write([]byte(str))
		//f.Write([]byte(" Person has three fields."))
	} else if c == 's' {
		// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}

func testScanf() {
	for i := 0; i < 2; i++ {
		var name string
		fmt.Print("Input Name:")
		n, err := fmt.Scanf("%s ", &name)
		fmt.Println(n, err, name)
	}
}
