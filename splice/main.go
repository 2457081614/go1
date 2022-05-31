package main

import "fmt"

func main() {

	b := make([]int, 0, 3)
	//a :=[]int{1,2,3,4,5,6}
	// 指定附加多少长度
	b = append(b, 1, 2, 3, 4, 5, 6, 7, 8, 9)[0:8]
	fmt.Printf("b 容量 %v", cap(b))
	fmt.Printf("ints 长度 %v", len(b))
}
