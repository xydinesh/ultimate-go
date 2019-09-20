package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var i int32

func foo() {
	fmt.Println("foo~~")
}

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			once.Do(foo) // goroutine
			ch <- i
		}
	}(ch)

	for i := 0; i < 5; i++ {
		once.Do(foo) // main routine
		fmt.Println(<-ch)
	}

	i = -1
	fmt.Println(i)

}
