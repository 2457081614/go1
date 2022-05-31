package main

import (
	"fmt"
)

func main() {
	/*ch := make(chan string)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	ch <- "e"
	ch <- "f"

	aa := <-ch
	fmt.Printf(aa)*/

	//ch := make(chan string)

	for i := 0; i < 10; i++ {
		go jc(i)
	}
	select {}

}

func jc(x int) int {
	sum := 0
	for i := 0; i < x; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
	return sum
}
