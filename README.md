I read this article https://go101.org/blog/2022-10-01-three-way-string-comparison.html

It states that `strings.Compare` in Golang is not optimzed, so it run slow.

It suggest to implement something like `bytes.Compare`, that relies on machine implementation and is thus more performant.

The article proposal is to "optimize `strings.Compare`". The Golang solution is to **never use** `strings.Compare` and leverage built-in operators instead (`==`, `<`, `>`).

I was curios so I decided to write some benchmark to understand what's the difference.

I implemented 4 test cases: 
1. string comparison using built-in operators
2. string comparison using `strings.Compare`
3. bytes comparison using `bytes.Compare` with string to bytes conversion just for the comparison
4. bytes comparison using `bytes.Compare` with byte sequences

The results? Here:

```
❯ go test -bench=. -v
=== RUN   TestFirst
--- PASS: TestFirst (0.00s)
=== RUN   TestBytes
--- PASS: TestBytes (0.00s)
=== RUN   TestBytesWithBytes
--- PASS: TestBytesWithBytes (0.00s)
=== RUN   TestStrings
--- PASS: TestStrings (0.00s)
goos: linux
goarch: amd64
pkg: strcmp
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkBuiltIn
BenchmarkBuiltIn-12             77651035                14.15 ns/op
BenchmarkBytes
BenchmarkBytes-12               25444228                45.15 ns/op
BenchmarkBytesWithBytes
BenchmarkBytesWithBytes-12      140633712                8.426 ns/op
BenchmarkStrings
BenchmarkStrings-12             142449582                7.818 ns/op
PASS
ok      strcmp  6.329s
```

It seems `strings.Compare` is optimised, at least as `bytes.Compare`.

Is this due to the small string size?

```
❯ go test -bench=. -v
=== RUN   TestFirst
--- PASS: TestFirst (0.00s)
=== RUN   TestBytes
--- PASS: TestBytes (0.00s)
=== RUN   TestBytesWithBytes
--- PASS: TestBytesWithBytes (0.00s)
=== RUN   TestStrings
--- PASS: TestStrings (0.00s)
goos: linux
goarch: amd64
pkg: strcmp
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkBuiltIn
BenchmarkBuiltIn-12                     74178211                14.99 ns/op
BenchmarkBytes
BenchmarkBytes-12                       26920994                42.95 ns/op
BenchmarkBytesWithBytes
BenchmarkBytesWithBytes-12              141717039                8.668 ns/op
BenchmarkBuiltIn_long
BenchmarkBuiltIn_long-12                74486179                15.57 ns/op
BenchmarkBytes_long
BenchmarkBytes_long-12                   3146968               391.6 ns/op
BenchmarkBytesWithBytes_long
BenchmarkBytesWithBytes_long-12         137667508                8.515 ns/op
BenchmarkStrings_long
BenchmarkStrings_long-12                90615900                14.01 ns/op
BenchmarkStrings
BenchmarkStrings-12                     150320742                7.846 ns/op
PASS
ok      strcmp  12.543s

```

Final results

|bench|short|long|
|---|---|---|
|built-in|14.99 ns/op|15.57 ns/op|
|strings|7.846 ns/op|14.01 ns/op|
|bytes (w/ conversion)|42.95 ns/op|391.6 ns/op|
|bytes (no conversion)|8.668 ns/op|8.515 ns/op|
