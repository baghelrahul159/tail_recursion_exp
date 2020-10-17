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
	var n uint64 = 50
	log.Println(recurse(n))

	wg.Wait()
}

func recurse(n uint64) uint64 {
	//breaking condition
	if n < 2 {
		return n
	}

	//recursing and processing
	return recurse(n-1) + recurse(n-2)
}
