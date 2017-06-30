Stream Bench
---

A stream lifts weights above its chest.

Huh?
---

Is it faster to use `json.Marshal()` or `json.NewEncoder().Encode()`?

Benchmark Results
---

```
BenchmarkMarshal-4   1000000        1246 ns/op
BenchmarkEncode-4    2000000         952 ns/op
ok    github.com/bentranter/streambench 4.150s
```

Yup, encoding is faster.
