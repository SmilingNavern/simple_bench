package main

import (
    "testing"
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
