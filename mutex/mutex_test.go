package main

import (
	"testing"
	"runtime"
)

func BenchmarkIncrementMutex(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		bs.IncrementMutex()
	}
}

func BenchmarkIncrementMutexNoDefer(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		bs.IncrementMutexNoDefer()
	}
}

func BenchmarkIncrementAtomic(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		bs.IncrementAtomic()
	}

}

func BenchmarkConcurrentMutex(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < runtime.NumCPU(); j++ {
			go bs.IncrementMutex()
		}
	}
}

func BenchmarkConcurrentMutexNoDefer(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < runtime.NumCPU(); j++ {
			go bs.IncrementMutexNoDefer()
		}
	}
}


func BenchmarkConcurrentAtomic(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < runtime.NumCPU(); j++ {
			go bs.IncrementAtomic()
		}
	}
}

