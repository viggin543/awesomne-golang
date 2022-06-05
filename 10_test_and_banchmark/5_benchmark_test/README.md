# run benchmark from terminal

```bash
  go test -v -run="nonde" -bench="BenchmarkSprintf"
  go test -v -run="nonde" -bench="BenchmarkSprintf" -benchtime=3s
  go test -v -run="nonde" -bench="BenchmarkSprintf" --benchmem
```

## Response looks like that
```bash
    BenchmarkSprintf-8   	30061965	        39.72 ns/op
```

The first result number represents how many times the action was called
The second is the avg tim it took

in our case, 302,962,279 times. ( 300+ million times)
and avg time per operation is 39 nanoseconds

wow...Golang on Apple Silicon is fast :)


## bench meme resp looks like
```bash
BenchmarkSprintf-8      30528499                39.63 ns/op            2 B/op          1 allocs/op
```

allocs/op value represents the number of heap allocations per operation
B/op value represents the number of bytes per operation
The call to Sprintf result in 2 bytes of memory being allocated per operation.
