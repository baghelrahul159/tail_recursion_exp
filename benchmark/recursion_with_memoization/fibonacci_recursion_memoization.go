package main

var masterFib = map[uint64]uint64{}

func fibWithRecursionMemoization(n uint64) uint64 {
	//breaking condition
	if n < 2 {
		return n
	}

	//memoization
	var num uint64
	num, ok := masterFib[n]
	if !ok {
		//recursing and processing
		num = fibWithRecursionMemoization(n-1) + fibWithRecursionMemoization(n-2)
		masterFib[n] = num
	}
	return num
}
