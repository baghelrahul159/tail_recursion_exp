package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Test Case 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{
				n: 6,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibWithRecursionMemoization(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFibWithRecursionMemoization(i uint64, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibWithRecursionMemoization(i)
	}
}

func BenchmarkFibWithRecursionMemoization1(b *testing.B)  { benchmarkFibWithRecursionMemoization(1, b) }
func BenchmarkFibWithRecursionMemoization2(b *testing.B)  { benchmarkFibWithRecursionMemoization(2, b) }
func BenchmarkFibWithRecursionMemoization5(b *testing.B)  { benchmarkFibWithRecursionMemoization(5, b) }
func BenchmarkFibWithRecursionMemoization10(b *testing.B) { benchmarkFibWithRecursionMemoization(10, b) }
func BenchmarkFibWithRecursionMemoization20(b *testing.B) { benchmarkFibWithRecursionMemoization(20, b) }
func BenchmarkFibWithRecursionMemoization50(b *testing.B) { benchmarkFibWithRecursionMemoization(50, b) }
func BenchmarkFibWithRecursionMemoization100(b *testing.B) {
	benchmarkFibWithRecursionMemoization(100, b)
}
func BenchmarkFibWithRecursionMemoization300(b *testing.B) {
	benchmarkFibWithRecursionMemoization(300, b)
}
func BenchmarkFibWithRecursionMemoization500(b *testing.B) {
	benchmarkFibWithRecursionMemoization(500, b)
}
func BenchmarkFibWithRecursionMemoization700(b *testing.B) {
	benchmarkFibWithRecursionMemoization(700, b)
}
func BenchmarkFibWithRecursionMemoization1000(b *testing.B) {
	benchmarkFibWithRecursionMemoization(1000, b)
}
