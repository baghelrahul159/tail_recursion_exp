package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

//curl http://localhost:6060/debug/pprof/heap --output heap.tar.gz

var masterFib = map[uint64]uint64{}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	//logic starts
	var n uint64 = 1000000
	log.Println(recurse(n))

	wg.Wait()
}

func recurse(n uint64) uint64 {
	//breaking condition
	if n < 2 {
		return n
	}

	//memoization
	var num uint64
	num, ok := masterFib[n]
	if !ok {
		//recursing and processing
		num = recurse(n-1) + recurse(n-2)
		masterFib[n] = num
	}
	return num
}
