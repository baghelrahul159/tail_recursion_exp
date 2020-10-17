package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	//logic starts
	var n uint64 = 1000000
	log.Println(recurse(n, 0, 1))

	wg.Wait()
}

func recurse(n uint64, a uint64, b uint64) uint64 {
	//breaking condition
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}

	//recursing and processing
	return recurse(n-1, b, a+b)
}
