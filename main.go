package main

import (
	"fmt"
	"time"

	concurrencypractice "github.com/leonardodelira/100-go-mistakes/8-concurrency-practice"
)

func main() {
	account := concurrencypractice.NewAccount2Mistake74()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			account.Mistake74_CopyingSyncType("foo")
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			account.Mistake74_CopyingSyncType("bar")
		}
	}()
	for {
		time.Sleep(1 * time.Second)
		fmt.Print("...")
	}
}
