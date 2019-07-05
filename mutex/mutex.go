package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Bla struct {
	i   uint64
	mux sync.Mutex
}

func (b *Bla) IncrementMutex() {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.i += 1
}

func (b *Bla) IncrementMutexNoDefer() {
	b.mux.Lock()

	b.i += 1

	b.mux.Unlock()
}

func (b *Bla) IncrementAtomic() {
	atomic.AddUint64(&b.i, 1)
}

func main() {
	fmt.Println("BENCH")
}
