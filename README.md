Simple bench 



# Mutex

```
go test -bench=.
goos: linux
goarch: amd64
BenchmarkIncrementMutex-8               20000000                50.2 ns/op
BenchmarkIncrementMutexNoDefer-8        100000000               16.6 ns/op
BenchmarkIncrementAtomic-8              300000000                5.32 ns/op
PASS
```
