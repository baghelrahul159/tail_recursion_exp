package main

func fibWithoutRecursion(n uint64) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0, 1

	for i := uint64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}
