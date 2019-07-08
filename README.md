Simple bench 



# Mutex

```
go test -bench=.
goos: linux
goarch: amd64
BenchmarkIncrementMutex-8               20000000                50.7 ns/op
BenchmarkIncrementMutexNoDefer-8        100000000               16.6 ns/op
BenchmarkIncrementAtomic-8              300000000                5.36 ns/op
BenchmarkConcurrentMutex-8               1000000              2334 ns/op
BenchmarkConcurrentMutexNoDefer-8        1000000              2095 ns/op
BenchmarkConcurrentAtomic-8              1000000              1995 ns/op
PASS
```


# Prepend

```
 go test -bench=.
goos: linux
goarch: amd64
BenchmarkPrepend-8                300000              4484 ns/op
BenchmarkSmartPrepend-8          2000000               838 ns/op
PASS
```
