package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	path := "D:\\maven\\repository\\antlr\\antlr\\2.7.2"
	stat, _ := os.Stat(path)
	fmt.Printf("is dir %v \n", stat.IsDir())
	fmt.Printf("name %v \n", stat.Name())
	fmt.Printf("file mode %v \n", stat.Mode())
	fmt.Printf("size dir %v \n", stat.Size())

	// 递归创建所有目录
	os.MkdirAll("G:\\杨柳真帅", os.ModePerm)

	// 只会创建一层目录
	err := os.Mkdir("G:\\杨柳真帅111\\2222222222", os.ModePerm)
	if nil != err {
		fmt.Println("error is", err)
	}

	path = "E:\\学习\\filebeat\\nacos.yml"
	open, err := os.Open(path)
	if nil != err {
		fmt.Println("打开文件错误:", err)
	} else {
		b := make([]byte, 1, 1024*8)
		count := -1
		for {
			count, err = open.Read(b)
			if count == 0 || err == io.EOF {
				fmt.Println("读取文件结束")
				break
			}
			fmt.Println("输出", string(b))
		}
	}
	open.Close()
}
