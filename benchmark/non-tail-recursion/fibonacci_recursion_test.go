package main

import "testing"

func Test_fibWithRecursion(t *testing.T) {
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
			if got := fibWithRecursion(tt.args.n); got != tt.want {
				t.Errorf("fibWithRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFibWithRecursion(i uint64, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibWithRecursion(i)
	}
}

func BenchmarkFibWithRecursion1(b *testing.B)  { benchmarkFibWithRecursion(1, b) }
func BenchmarkFibWithRecursion2(b *testing.B)  { benchmarkFibWithRecursion(2, b) }
func BenchmarkFibWithRecursion5(b *testing.B)  { benchmarkFibWithRecursion(5, b) }
func BenchmarkFibWithRecursion10(b *testing.B) { benchmarkFibWithRecursion(10, b) }
func BenchmarkFibWithRecursion20(b *testing.B) { benchmarkFibWithRecursion(20, b) }

// func BenchmarkFibWithRecursion50(b *testing.B) { benchmarkFibWithRecursion(50, b) }

// func BenchmarkFibWithRecursion100(b *testing.B) {
// 	benchmarkFibWithRecursion(100, b)
// }
// func BenchmarkFibWithRecursion300(b *testing.B) {
// 	benchmarkFibWithRecursion(300, b)
// }
// func BenchmarkFibWithRecursion500(b *testing.B) {
// 	benchmarkFibWithRecursion(500, b)
// }
// func BenchmarkFibWithRecursion700(b *testing.B) {
// 	benchmarkFibWithRecursion(700, b)
// }
// func BenchmarkFibWithRecursion1000(b *testing.B) {
// 	benchmarkFibWithRecursion(1000, b)
// }
