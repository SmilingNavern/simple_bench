package main

import (
	"testing"
)

func BenchmarkPrepend(b *testing.B) {
	b.StopTimer()
	bla := make([]int, 4096, 8192)
	for i := 0; i < 4096; i++ {
		bla[i] = i
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		newbla := make([]int, len(bla), cap(bla))
		copy(newbla, bla)
		b.StartTimer()
		newbla = append([]int{10}, newbla...)
	}
}

func BenchmarkSmartPrepend(b *testing.B) {
	b.StopTimer()
	bla := make([]int, 4096, 8192)
	for i := 0; i < 4096; i++ {
		bla[i] = i
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		newbla := make([]int, len(bla), cap(bla))
		copy(newbla, bla)
		b.StartTimer()
		newbla = append(newbla[:1], newbla...)
		newbla[0] = 10
	}
}
