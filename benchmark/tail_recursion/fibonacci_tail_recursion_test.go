package main

import "testing"

func Test_fibTailRecursion(t *testing.T) {
	type args struct {
		n uint64
		a uint64
		b uint64
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
				a: 0,
				b: 1,
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				n: 3,
				a: 0,
				b: 1,
			},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{
				n: 6,
				a: 0,
				b: 1,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibTailRecursion(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("fibTailRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFibTailRecursion(i uint64, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibTailRecursion(i, 0, 1)
	}
}

func BenchmarkTailFib1(b *testing.B)    { benchmarkFibTailRecursion(1, b) }
func BenchmarkTailFib2(b *testing.B)    { benchmarkFibTailRecursion(2, b) }
func BenchmarkTailFib5(b *testing.B)    { benchmarkFibTailRecursion(5, b) }
func BenchmarkTailFib10(b *testing.B)   { benchmarkFibTailRecursion(10, b) }
func BenchmarkTailFib20(b *testing.B)   { benchmarkFibTailRecursion(20, b) }
func BenchmarkTailFib50(b *testing.B)   { benchmarkFibTailRecursion(50, b) }
func BenchmarkTailFib100(b *testing.B)  { benchmarkFibTailRecursion(100, b) }
func BenchmarkTailFib300(b *testing.B)  { benchmarkFibTailRecursion(300, b) }
func BenchmarkTailFib500(b *testing.B)  { benchmarkFibTailRecursion(500, b) }
func BenchmarkTailFib700(b *testing.B)  { benchmarkFibTailRecursion(700, b) }
func BenchmarkTailFib1000(b *testing.B) { benchmarkFibTailRecursion(1000, b) }
