package main

import(
	//"time"
	"sync"
	"fmt"
)

var lock sync.Mutex

func main() {

	testMap()
}
func testMap() {
	a := make(chan int, 1)
	b := make(chan string, 1)

	a <- 1
	b <- "aaaa"

	select {
	case v := <- a:
		fmt.Println(v)
	case f := <- b:
		panic(f)
	}
}
