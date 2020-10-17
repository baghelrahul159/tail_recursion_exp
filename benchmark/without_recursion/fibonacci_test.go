package main

import "testing"

func Test_fibWithoutRecursion(t *testing.T) {
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
			if got := fibWithoutRecursion(tt.args.n); got != tt.want {
				t.Errorf("fibWithoutRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFib(i uint64, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibWithoutRecursion(i)
	}
}

func BenchmarkFib1(b *testing.B)    { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)    { benchmarkFib(2, b) }
func BenchmarkFib5(b *testing.B)    { benchmarkFib(5, b) }
func BenchmarkFib10(b *testing.B)   { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B)   { benchmarkFib(20, b) }
func BenchmarkFib50(b *testing.B)   { benchmarkFib(50, b) }
func BenchmarkFib100(b *testing.B)  { benchmarkFib(100, b) }
func BenchmarkFib300(b *testing.B)  { benchmarkFib(300, b) }
func BenchmarkFib500(b *testing.B)  { benchmarkFib(500, b) }
func BenchmarkFib700(b *testing.B)  { benchmarkFib(700, b) }
func BenchmarkFib1000(b *testing.B) { benchmarkFib(1000, b) }
func BenchmarkFib2000(b *testing.B) { benchmarkFib(2000, b) }
func BenchmarkFib3000(b *testing.B) { benchmarkFib(3000, b) }
