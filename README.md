# one-euro-filter

Simple golang implementation of a [One Euro Filter](https://jaantollander.com/post/noise-filtering-using-one-euro-filter/).

```go
filter := NewOneEuroFilter(t0, x0, dx0, minCutoff, beta, dCutoff float64)
xHat := filter.Update(t, x)
```

Based on https://github.com/jaantollander/OneEuroFilter
