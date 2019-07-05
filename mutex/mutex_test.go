package main

import (
	"testing"
	"runtime"
	"sync"
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
		var wg sync.WaitGroup
		for j := 0; j < runtime.NumCPU(); j++ {
			wg.Add(1)
			go func() {
				bs.IncrementMutex()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentMutexNoDefer(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < runtime.NumCPU(); j++ {
			wg.Add(1)
			go func() {
				bs.IncrementMutexNoDefer()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}


func BenchmarkConcurrentAtomic(b *testing.B) {
	bs := &Bla{}
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < runtime.NumCPU(); j++ {
			wg.Add(1)
			go func() {
				bs.IncrementAtomic()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

