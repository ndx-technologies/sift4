SIFT4 — fast approximate string distance algorithm

* zero memory copy
* 100% test coverage

```bash
$ go test -benchmem -bench .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/sift4
cpu: Apple M3 Max
BenchmarkSIFT4Distance/empty-16                         956523358                1.109 ns/op           0 B/op          0 allocs/op
BenchmarkSIFT4Distance/one_empty-16                     1000000000               1.093 ns/op           0 B/op          0 allocs/op
BenchmarkSIFT4Distance/equal-16                         562000238                2.137 ns/op           0 B/op          0 allocs/op
BenchmarkSIFT4Distance/different-16                     16788264                71.80 ns/op           48 B/op          2 allocs/op
BenchmarkSIFT4Distance/long_different-16                 1488578               809.1 ns/op            24 B/op          1 allocs/op
BenchmarkSIFT4Distance/buffer/empty-16                  1000000000               1.122 ns/op           0 B/op          0 allocs/op
BenchmarkSIFT4Distance/buffer/one_empty-16              1000000000               1.116 ns/op           0 B/op          0 allocs/op
BenchmarkSIFT4Distance/buffer/equal-16                  552020344                2.179 ns/op           0 B/op          0 allocs/op
BenchmarkSIFT4Distance/buffer/different-16              29124799                41.24 ns/op            0 B/op          0 allocs/op
BenchmarkSIFT4Distance/buffer/long_different-16          1520445               789.8 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/sift4       11.848s
```
