package main

func fibWithRecursion(n uint64) uint64 {
	//breaking condition
	if n < 2 {
		return n
	}

	//recursing and processing
	return fibWithRecursion(n-1) + fibWithRecursion(n-2)
}
