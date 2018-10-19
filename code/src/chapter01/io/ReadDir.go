package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

//列出路径下的文件
//path 路径
//向下递归深度
func ListAllFile(path string, deepth int) {
	if deepth < 0 {
		return
	}

	var dirSeprater = "/"
	ostype := runtime.GOOS
	if -1 != strings.Index(ostype, "windows") {
		dirSeprater = "\\"
	}

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := deepth; tmpHier > 0; tmpHier-- {
				fmt.Printf("|  ")
			}
			fmt.Println(info.Name(), "/")
			ListAllFile(path+dirSeprater+info.Name(), deepth-1)
			//ExapmleReadDir(path+dirSeprater+info.Name(), curHier+1)
		} else {
			for tmpHier := deepth; tmpHier > 0; tmpHier-- {
				fmt.Printf("|  ")
			}
			fmt.Println(info.Name())
		}
	}
}
