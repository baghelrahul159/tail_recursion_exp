# Tail Recursion Benchmarking and Profiling

This repository is an attempt to `benchmark` & `profile` different ways for listing `n-th Fibonacci number` using `Golang` as programming language.

These are 4 candidates for caculating n-th fibonacci number that we have considered.
1. Iterative method.
2. Recursie method without using memoization.
3. Recursive method using memoization for optimization.
4. Tail Recursion.

## Prerequisites

* You have installed the latest version of `Golang`
* You can use any OS, but that would have an impact on benchmark and CPU profile generated. The results here are from Macbook Air using `1.8 GHz Intel Core i5` and `8 GB 1600 MHz DDR3`
* If you want any help with profiling, consider reading https://medium.com/@openmohan/profiling-in-golang-3e51c68eb6a8

## Instructions 

If you want to run benchmarks on your own machine.

1. Clone the repository.
2. Move to benchmark folder and inside any sub-folder run below command.
```
cd benchmark/without_recursion
go test -run=XXX -bench=.
```

## Results and Analysis

Here's the benchmarks results:-

1. Recursion without memoization

```
goos: darwin
goarch: amd64
pkg: github.com/baghelrahul159/tail_recursion_exp/benchmark/non-tail-recursion
BenchmarkFibWithRecursion1-4    	443103949	         2.60 ns/op
BenchmarkFibWithRecursion2-4    	165883412	         7.25 ns/op
BenchmarkFibWithRecursion5-4    	30739672	        41.1 ns/op
BenchmarkFibWithRecursion10-4   	 2572047	       455 ns/op
BenchmarkFibWithRecursion20-4   	   18165	     80904 ns/op
PASS
```

It is clear, with recursion without memoization, even 20-th fibonacci number takes 80904 ns/op. Thus, we will omit it from further discussions.

2. Recursion with Memoization

```
goos: darwin
goarch: amd64
pkg: github.com/baghelrahul159/tail_recursion_exp/benchmark/recursion_with_memoization
BenchmarkFibWithRecursionMemoization1-4      	415333540	         3.00 ns/op
BenchmarkFibWithRecursionMemoization2-4      	149585394	         8.03 ns/op
BenchmarkFibWithRecursionMemoization5-4      	100000000	        10.8 ns/op
BenchmarkFibWithRecursionMemoization10-4     	49885395	        25.9 ns/op
BenchmarkFibWithRecursionMemoization20-4     	66632116	        20.2 ns/op
BenchmarkFibWithRecursionMemoization50-4     	45081136	        27.7 ns/op
BenchmarkFibWithRecursionMemoization100-4    	41284546	        33.4 ns/op
BenchmarkFibWithRecursionMemoization300-4    	45022812	        28.7 ns/op
BenchmarkFibWithRecursionMemoization500-4    	75619054	        19.7 ns/op
BenchmarkFibWithRecursionMemoization700-4    	69009540	        21.3 ns/op
BenchmarkFibWithRecursionMemoization1000-4   	80567641	        16.3 ns/op
PASS
```

3. Tail Recursion

```
goos: darwin
goarch: amd64
pkg: github.com/baghelrahul159/tail_recursion_exp/benchmark/tail_recursion
BenchmarkTailFib1-4      	373255960	         4.33 ns/op
BenchmarkTailFib2-4      	157686687	         7.42 ns/op
BenchmarkTailFib5-4      	73597358	        24.2 ns/op
BenchmarkTailFib10-4     	16147897	        69.8 ns/op
BenchmarkTailFib20-4     	13154060	       126 ns/op
BenchmarkTailFib50-4     	 5215443	       270 ns/op
BenchmarkTailFib100-4    	 2578944	       477 ns/op
BenchmarkTailFib300-4    	  965191	      1368 ns/op
BenchmarkTailFib500-4    	  438627	      3005 ns/op
BenchmarkTailFib700-4    	  339271	      4191 ns/op
BenchmarkTailFib1000-4   	  263703	      5076 ns/op
PASS
```

4. Iterative

```
goos: darwin
goarch: amd64
pkg: github.com/baghelrahul159/tail_recursion_exp/benchmark/without_recursion
BenchmarkFib1-4      	441498200	         2.52 ns/op
BenchmarkFib2-4      	460432479	         2.40 ns/op
BenchmarkFib5-4      	231182103	         6.86 ns/op
BenchmarkFib10-4     	118818930	        12.0 ns/op
BenchmarkFib20-4     	58351034	        20.1 ns/op
BenchmarkFib50-4     	23743794	        59.4 ns/op
BenchmarkFib100-4    	17107134	        68.1 ns/op
BenchmarkFib300-4    	 6996146	       194 ns/op
BenchmarkFib500-4    	 4176632	       359 ns/op
BenchmarkFib700-4    	 2957839	       398 ns/op
BenchmarkFib1000-4   	 2071720	      1038 ns/op
BenchmarkFib2000-4   	 1000000	      1153 ns/op
BenchmarkFib3000-4   	  611624	      3048 ns/op
PASS
```

Plotting the above values using matplotlib library from python

![Benchmarking Plot](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/plot.png)

This illustrates the fact that memoization greatly improves the speed at the cost of more memory. However interesting to note here is that Tail recursion and Iterative method are very close to each other. 

### Flamegraphs for Heap profling

1. Recursion without memoization

![Heap Profile Non Tail Recursion](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/non-tail-recursion.png)

2. Recursion with Memoization

![Heap Profile Recursion with memoization](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/recursion_memoization.png)

3. Tail Recursion

![Heap Profile Tail Recursion](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/tail_recursion.png)

4. Iterative

![Heap Profile Iterative](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/iterative.png)

`Apart from memoization technique which uses almost 30 MB, rest all methods are using few KB.`

### Flamegraphs for CPU profling

1. Recursion without memoization

![CPU Profile Non Tail Recursion](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/non-tail-recursion-cpu.png)

2. Recursion with Memoization

![CPU Profile Recursion with memoization](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/recursion_memoization-cpu.png)

3. Tail Recursion

![CPU Profile Tail Recursion](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/tail_recursion-cpu.png)

4. Iterative

![CPU Profile Iterative](https://github.com/baghelrahul159/tail_recursion_exp/blob/master/images/iterative-cpu.png)

`Without memoization recursion takes almost 18s which is reduced to 160ms with memoization, tail recursion and iterative function takes almost same amount of cpu time.`
