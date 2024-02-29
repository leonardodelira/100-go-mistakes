package main

import (
	"time"

	thestandardlibrary "github.com/leonardodelira/100-go-mistakes/9-the-standard-library"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
		time.Sleep(2 * time.Second)
		ch <- 2
	}(ch)

	thestandardlibrary.Mistake76_consumer(ch)
}
