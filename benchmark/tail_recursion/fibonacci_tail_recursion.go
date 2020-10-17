package main

func fibTailRecursion(n, a, b uint64) uint64 {

	//breaking condition
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	//recursing and processing
	return fibTailRecursion(n-1, b, a+b)
}
