Simple bench 



# Mutex

```
go test -bench=.
goos: linux
goarch: amd64
BenchmarkIncrementMutex-8               30000000                50.2 ns/op
BenchmarkIncrementMutexNoDefer-8        100000000               16.6 ns/op
BenchmarkIncrementAtomic-8              200000000                7.37 ns/op
BenchmarkConcurrentMutex-8               1000000              2003 ns/op
BenchmarkConcurrentMutexNoDefer-8        1000000              1727 ns/op
BenchmarkConcurrentAtomic-8              1000000              1920 ns/op
PASS
```
