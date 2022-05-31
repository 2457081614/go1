package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://www.baidu.com")
	//b := make([]byte, 1024)
	//var result string
	//for {
	//	count, _ := response.Body.Read(b)
	//	result = result + string(b)
	//	if count < 1024 {
	//		break
	//	}
	//}

	all, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("response body  %v", string(all))

}
