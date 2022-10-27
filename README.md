I read this article https://go101.org/blog/2022-10-01-three-way-string-comparison.html

It states that `strings.Compare` in Golang is not optimzed, so it run slow.

It suggest to implement something like `bytes.Compare`, that relies on machine implementation and is thus more performant.

The article proposal is to "optimize `strings.Compare`". The Golang solution is to **never use** `strings.Compare` and leverage built-in operators instead (`==`, `<`, `>`).

I was curios so I decided to write some benchmark to understand what's the difference.

Initial assumptions:
- `strings.Compare` is 2x slower than `bytes.Compare`
- built-ins are faster than `strings.Compare`
- converting string to bytes and using `bytes.Compare` may be beneficial for longer strings

I implemented 4 test cases: 
1. string comparison using built-in operators
2. string comparison using `strings.Compare`
3. bytes comparison using `bytes.Compare` with string to bytes conversion just for the comparison
4. bytes comparison using `bytes.Compare` with byte sequences

The results? Here testing `foobar` with `hellow` (same length to test for worst case scenario):

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
BenchmarkBuiltIn-12                     76338814                14.79 ns/op
BenchmarkBytes
BenchmarkBytes-12                       27815754                43.59 ns/op
BenchmarkBytesWithBytes
BenchmarkBytesWithBytes-12              139774934                8.487 ns/op
BenchmarkStrings
BenchmarkStrings-12                     83908431                14.46 ns/op
BenchmarkBuiltIn
PASS
ok      strcmp  6.329s

```

It seems `strings.Compare` is optimised, at least as `bytes.Compare`.

Is this due to the small string size? Running test with longer strings (255 char each):

```
❯ go test -bench=. -v > test3.txt]
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
BenchmarkBuiltIn-12                     76338814                14.79 ns/op
BenchmarkBytes
BenchmarkBytes-12                       27815754                43.59 ns/op
BenchmarkBytesWithBytes
BenchmarkBytesWithBytes-12              139774934                8.487 ns/op
BenchmarkBuiltIn_long
BenchmarkBuiltIn_long-12                75807630                15.70 ns/op
BenchmarkBytes_long
BenchmarkBytes_long-12                   2300572               468.1 ns/op
BenchmarkBytesWithBytes_long
BenchmarkBytesWithBytes_long-12         137647111                8.276 ns/op
BenchmarkStrings_long
BenchmarkStrings_long-12                85086740                13.62 ns/op
BenchmarkStrings
BenchmarkStrings-12                     83908431                14.46 ns/op
PASS
ok      strcmp  11.701s

```

Final results

|bench|short|long|
|---|---|---|
|built-in|14.79 ns/op|15.70 ns/op|
|strings|14.46 ns/op|13.62 ns/op|
|bytes (w/ conversion)|43.59 ns/op|468.1 ns/op|
|bytes (no conversion)|8.487 ns/op|8.276 ns/op|

I find the results interesting:
1. **do not** convert a string to bytes, you're nuking performances 🪃
2. `strings.Compare` is slightly less than 2x more CPU than `bytes.Compare` but seems consistent across input length
3. built-ins are not more performant 🤷
