# run benchmark from terminal

```bash
  go test -v -run="nonde" -bench="BenchmarkSprintf"
  go test -v -run="nonde" -bench="BenchmarkSprintf" -benchtime=3s
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