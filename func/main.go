package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//闭包函数
	f := func(data string) {
		fmt.Printf("value is %T", data)
	}
	f("hello world")

	//闭包优点
	// 1.加强模块化  2.抽象

	f2 := add()
	for i := 0; i < 10; i++ {
		f2(i)
	}

	f3 := rand.Float64()
	fmt.Printf("value is %v", f3)
}

func add() func(int) int {
	sum := 0
	return func(i int) int {
		fmt.Printf("sum1 = %v \n", sum)
		sum = sum + i
		fmt.Printf("sum2 = %v \n", sum)
		return sum
	}
}
