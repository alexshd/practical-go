# Practical Go for Developers

ArdanLabs ∴  2025 <br />

Miki Tebeka
<i class="far fa-envelope"></i> [miki@ardanlabs.com](mailto:miki@ardanlabs.com), <i class="fab fa-twitter"></i> [@tebeka](https://twitter.com/tebeka), <i class="fab fa-linkedin-in"></i> [mikitebeka](https://www.linkedin.com/in/mikitebeka/), <i class="fab fa-blogger-b"></i> [blog](https://www.ardanlabs.com/blog/), <i class="fa-brands fa-medium"></i>[medium](https://medium.com/@tebeka).

#### Shameless Plugs

- [Go Essential Training](https://www.linkedin.com/learning/go-essential-training/) - LinkedIn Learning
    - [Rest of classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Go Brain Teasers](https://pragprog.com/titles/d-gobrain/go-brain-teasers/) book
    - [Rest of books](https://pragprog.com/search/?q=miki+tebeka)

[FINAL EXERCISE](_extra/dld.md)

---

## Day 1

### Agenda

- Strings & formatted output
    - What's a string?
    - Unicode basics
    - Using fmt package for formatted output
- Calling REST APIs
    - Making HTTP calls with net/http
    - Defining structs
    - Serializing JSON
- Working with files
    - Handling errors
    - Using defer to manage resources
    - Working with io.Reader & io.Writer interfaces

### Code


- [hw.go](ws/sandboxaq/hw/hw.go) - Hello World
- [banner.go](ws/sandboxaq/banner/banner.go) - Working with strings
- [github.go](ws/sandboxaq/github/github.go) - Calling REST APIs
- [kill_server.go](ws/sandboxaq/kill_server/kill_server.go) - Working with files, errors
- [sha1.go](ws/sandboxaq/sha1/sha1.go) - Using `io.Reader` & `io.Writer`
- [scope.go](ws/sandboxaq/scope/scope.go) - Variable scope

### Links

- [GoReleaser](https://goreleaser.com/) - Building Go executables
- [HTTP status cats](https://http.cat/)
- [errors](https://pkg.go.dev/errors/) package ([Go 1.13](https://go.dev/blog/go1.13-errors))
- [encoding/json](https://pkg.go.dev/encoding/json)
- [net/http](https://pkg.go.dev/net/http)
- Numbers
    - [math/big](https://pkg.go.dev/math/big/) - Big numbers
    - [Numeric types](https://go.dev/ref/spec#Numeric_types)
- Strings
    - [Unicode table](https://unicode-table.com/)
    - [strings](https://pkg.go.dev/strings/) package - string utilities
    - [Go strings](https://go.dev/blog/strings)
    - [Plain Text](https://www.youtube.com/watch?v=gd5uJ7Nlvvo) - Great Video on Unicode
- [Go Proverbs](https://go-proverbs.github.io/) - Think about them ☺
- [Annotated "Hello World"](https://www.353solutions.com/annotated-go)
- [Effective Go](https://go.dev/doc/effective_go.html) - Read this!
- [Go standard library](https://pkg.go.dev/) - official documentation
- [A Tour of Go](https://do.dev/tour/)
- Setting Up
    - [The Go SDK](https://go.dev/dl/)
    - [git](https://git-scm.com/)
    - IDE's: [Visual Studio Code](https://code.visualstudio.com/) + [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) or [Goland](https://www.jetbrains.com/go/) (paid)

### Data & Other

- `G☺`
- `♡`
- [http.log.gz](_extra/http.log.gz)
- `https://api.github.com/users/tebeka`
- [Slides](_extra/slides.pdf)
- [Unicode](_extra/unicode.pdf)

---

## Day 2

### Agenda

- Sorting
    - Working with slices
    - Writing methods
    - Understanding interfaces
- Catching panics
    - The built-in recover function
    - Named return values
- Processing text
    - Reading line by line with bufio.Scanner
    - Using regular expressions
    - Working with maps

### Code

- [cart.go](cart/cart.go) - Slices
- [game.go](game/game.go) - Structs, methods & interfaces
- [empty.go](empty/empty.go) - The empty interface
- [stats.go](stats/stats.go) - Using generics

### Links

- [Hyrum's Law](https://www.hyrumslaw.com/)
- [regex101](https://regex101.com/) - Regular expression builder
- [sort examples](https://pkg.go.dev/sort/#pkg-examples) - Read and try to understand
- [When to use generics](https://go.dev/blog/when-generics)
- [Generics tutorial](https://go.dev/doc/tutorial/generics)
- [Methods, interfaces & embedded types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
- [Methods & Interfaces](https://go.dev/tour/methods/1) in the Go tour
- Slices
    - [Slices](https://go.dev/blog/slices) & [Slice internals](https://go.dev/blog/go-slices-usage-and-internals) on the Go blog
    - [Slice tricks](https://github.com/golang/go/wiki/SliceTricks)
- Error Handling
    - [Defer, Panic and Recover](https://go.dev/blog/defer-panic-and-recover)
    - [errors](https://pkg.go.dev/errors/) package ([Go 1.13](https://go.dev/blog/go1.13-errors))
    - [pkg/errors](https://github.com/pkg/errors)
- [Composition over inheritance](https://en.wikipedia.org/wiki/Composition_over_inheritance)

### Data & Other

- [Cost of errors](_extra/cost.png)
- [sherlock.txt](_extra/sherlock.txt)
- [slices](_extra/slices.md)
- [Method sets](_extra/method-sets.pdf)

---

## Day 3

### Agenda

- Distributing work
    - Using goroutines & channels
    - Using the sync package to coordinate work
- Timeouts & cancellation
    - Working with multiple channels using select
    - Using context for timeouts & cancellations
    - Standard library support for context

### Code


- [div.go](div/div.go) - Handling panics
- [freq.go](freq/freq.go) - Processing text
- [go_chan.go](go_chan/go_chan.go) - Channels and goroutines
- [taxi_check.go](taxi/taxi_check.go) - Converting sequential code to concurrent
- [rtb.go](rtb/rtb.go) - Using context for timeouts
- [counter.go](counter/counter.go) - Using mutex and atomic

### Links

- [The race detector](https://go.dev/doc/articles/race_detector)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [Data Race Patterns in Go](https://eng.uber.com/data-race-patterns-in-go/)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)
- [The Behavior of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
- [Channel Semantics](https://www.353solutions.com/channel-semantics)
- [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
- [Amdahl's Law](https://en.wikipedia.org/wiki/Amdahl%27s_law) - Limits of concurrency
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576/photo/1)
- [Concurrency is not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) by Rob Pike
- [Scheduling in Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html) by Bill Kennedy
- [The cost of Go's panic and recover](https://jub0bs.com/posts/2025-02-28-cost-of-panic-recover/)

### Data & Other

- [rtb.go](_extra/rtb.go)
- [site_times.go](_extra/site_time.go)
- [taxi.tar](https://storage.googleapis.com/353solutions/c/data/taxi.tar)

---

## Day 4

### Agenda

- Testing your code
    - Working with the testing package
    - Using testify
    - Managing dependencies with go mod
- Structuring your code
    - Writing sub-packages
- Writing an HTTP server
    - Writing handlers
    - Testing handlers
Adding metrics & logging
    - Using expvar for metrics
    - Using the log package and a look at user/zap
- Configuration patterns
    - Reading environment variables and a look at external packages
    - Using the flag package for command line processing

### Code

- [nlp](nlp) - `nlp` project

### Links

- [flags](https://github.com/tebeka/flags) - More types for the built-in `flag` package
- [Conway's Law](https://martinfowler.com/bliki/ConwaysLaw.html)
- [The Twelve-Factor App](https://12factor.net)
- Configuration
    - [conf](https://pkg.go.dev/github.com/ardanlabs/conf/v3)
    - [viper](https://github.com/spf13/viper) & [cobra](https://github.com/spf13/cobra)
- Logging 
    - Built-in [slog](https://pkg.go.dev/log/slog/)
    - [uber/zap](https://pkg.go.dev/go.uber.org/zap)
- Metrics
    - Built-in [expvar](https://pkg.go.dev/expvar/)
    - [Open Telemetry](https://opentelemetry.io/)
    - [Prometheus](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus)
- External web servers:
    - [chi](https://go-chi.io/) - Miki's favorite (after stdlib :)
    - [gorilla](https://gorilla.github.io/) - Old timer
    - [Gin](https://gin-gonic.com/) - Popular framework
    - [fasthttp](https://github.com/valyala/fasthttp) - Low level and fast
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
- [Example Project Structure](https://github.com/ardanlabs/service)
- [Organizing a Go module](https://go.dev/doc/modules/layout)
- [How to Write Go Code](https://go.dev/doc/code.html)
- Documentation
    - [Godoc: documenting Go code](https://go.dev/blog/godoc)
    - [Go Doc Comments](https://go.dev/doc/comment)
    - [Testable examples in Go](https://go.dev/blog/examples)
    - [Go documentation tricks](https://godoc.org/github.com/fluhus/godoc-tricks)
    - [gob/doc.go](https://github.com/golang/go/blob/master/src/encoding/gob/doc.go) of the `gob` package. Generates [this documentation](https://pkg.go.dev/encoding/gob/)
    - `go install golang.org/x/pkgsite/cmd/pkgsite@latest` (require go 1.18+)
    - `pkgsite -http=:8080` (open browser on http://localhost:8080/${module name})
- [Out Software Dependency Problem](https://research.swtch.com/deps) - Good read on dependencies by Russ Cox
- Linters (static analysis)
    - [staticcheck](https://staticcheck.io/)
    - [golangci-lint](https://golangci-lint.run/)
    - [gosec](https://github.com/securego/gosec) - Security oriented
    - [vulncheck](https://pkg.go.dev/golang.org/x/vuln/vulncheck) - Check for CVEs
    - [golang.org/x/tools/go/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) - Helpers to write analysis tools (see [example](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/))
- Testing
    - [testing](https://pkg.go.dev/testing/)
    - [testify](https://pkg.go.dev/github.com/stretchr/testify) - Many test utilities (including suites & mocking)
    - [Ginkgo](https://onsi.github.io/ginkgo/)
    - [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
        - [testing/quick](https://pkg.go.dev/testing/quick) - Initial fuzzing library
    - [test containers](https://golang.testcontainers.org/)
- HTTP Servers
    - [net/http](https://pkg.go.dev/net/http/)
    - [net/http/httptest](https://pkg.go.dev/net/http/httptest)
    - [chi](https://github.com/go-chi/chi) - A nice web framework

### Data & Other

- [nlp.go](_extra/nlp.go)
- [stemmer.go](_extra/stemmer.go)
- [tokenize_cases.toml](_extra/tokenize_cases.toml)
    - github.com/BurntSushi/toml
