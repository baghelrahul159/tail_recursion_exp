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
	var n uint64 = 10000000000
	var n2, n1 uint64 = 0, 1

	for i := uint64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	log.Println(n2 + n1)

	wg.Wait()
}
