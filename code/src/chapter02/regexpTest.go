package main

import (
	"fmt"
	"regexp"
)

func TestRegExp() {
	str1 := "dfsfn3iuhrui23qhrduwfbufwef134"
	str2 := "123456453werwegfwbrehytrewfvsdf"

	if preg, err := regexp.Compile("^df"); err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(preg.FindString(str1))
		fmt.Println(preg.FindString(str2))
	}

	if preg, err := regexp.Compile("df$"); err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(preg.FindString(str1))
		fmt.Println(preg.FindString(str2))
	}

}
