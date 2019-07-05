Simple bench 



# Mutex

```
go test -bench=.
goos: linux
goarch: amd64
BenchmarkIncrementMutex-8               30000000                50.4 ns/op
BenchmarkIncrementMutexNoDefer-8        100000000               17.3 ns/op
BenchmarkIncrementAtomic-8              300000000                5.34 ns/op
BenchmarkConcurrentMutex-8               1000000              2144 ns/op
BenchmarkConcurrentMutexNoDefer-8        1000000              1738 ns/op
PASS
```
