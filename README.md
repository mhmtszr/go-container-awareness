# Go Container Awareness [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Go Report Card][go-report-img]][go-report]

**Go Container Awareness** is an open-source Go library that provides container awareness for CPU and memory limits, facilitating efficient resource management within containerized environments.

Automatically set `GOMEMLIMIT` to match Linux container memory quota(%90 of memory limit).

Automatically set `GOGC` off.

Automatically set `GOMAXPROCS` to match Linux container CPU quota.


## Installation

```go
import _ "github.com/mhmtszr/go-container-awareness"

func main() {
    ...
}
```


[doc-img]: https://godoc.org/github.com/mhmtszr/go-container-awareness?status.svg
[doc]: https://godoc.org/github.com/mhmtszr/go-container-awareness
[ci-img]: https://github.com/mhmtszr/go-container-awareness/actions/workflows/build-test.yml/badge.svg
[ci]: https://github.com/mhmtszr/go-container-awareness/actions/workflows/build-test.yml
[go-report-img]: https://goreportcard.com/badge/github.com/mhmtszr/go-container-awareness
[go-report]: https://goreportcard.com/report/github.com/mhmtszr/go-container-awareness