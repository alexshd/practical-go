# Practical Go for Developers

ArdanLabs ∴  2025 <br />

This repo contains the material for the "Practical Go Foundations" class.
The code & links are synced with the [online class](for Developers).

This is an assorted collection of exercises for teaching, not a real Go project.

## Setting Up

We highly recommend that you set up a local working environment and follow along with the videos.
To setup a local environment install the following:
- The Go SDK either from your package manager (`brew`, `apt`, `choco` ...) or from [here](https://go.dev/dl/)
- `git`
- And IDE such as [VSCode](https://code.visualstudio.com/) with [the Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) or [GoLand](https://www.jetbrains.com/go/)

However, if you want to jump right in, you can use GitHub codespaces with the course repository.
To start a codespaces follow these steps:

- Click on the green "<> Code" button
- Select the "Codespaces" tab
- Click on the green "Create codespace on main" button

After a while you should have a new tab with Visual Studio Code already set up.
You can read more about codespaces [here](https://github.com/features/codespaces).

Miki Tebeka
<i class="far fa-envelope"></i> [miki@ardanlabs.com](mailto:miki@ardanlabs.com), <i class="fab fa-twitter"></i> [@tebeka](https://twitter.com/tebeka), <i class="fab fa-linkedin-in"></i> [mikitebeka](https://www.linkedin.com/in/mikitebeka/), <i class="fab fa-blogger-b"></i> [blog](https://www.ardanlabs.com/blog/), <i class="fa-brands fa-medium"></i>[medium](https://medium.com/@tebeka).

#### Shameless Plugs

- [Go Essential Training](https://www.linkedin.com/learning/go-essential-training/) - LinkedIn Learning
    - [Rest of classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Go Brain Teasers](https://pragprog.com/titles/d-gobrain/go-brain-teasers/) book
    - [Rest of books](https://pragprog.com/search/?q=miki+tebeka)

---

### Code


- [hw.go](hw/hw.go) - Hello World
- [banner.go](banner/banner.go) - Working with strings, Unicode
- [github.go](github/github.go) - Calling REST APIs, JSON
- [kill_server.go](kill_server/kill_server.go) - Working with files, defer, error handling
- [div.go](div/div.go) - Handling panics
- [sha1.go](sha1/sha1.go) - Using `io.Reader` and `io.Writer`
- [cart.go](cart/cart.go) - Working with slices
- [game.go](game/game.go) - Structs, methods & interfaces
- [empty.go](empty/empty.go) - The empty interface
- [stats.go](stats/stats.go) - Using generics
- [freq.go](freq/freq.go) - Processing text, regular expressions
- [go_chan.go](go_chan/go_chan.go) - Goroutines and channels
- [urls.go](urls/urls.go) - Fan out patterns
- [taxi_check.go](taxi_check/taxi_check.go) - Convert sequential code to concurrent
- [count.go](count/count.go) - Using mutex and atomic, the race detector
- [select.go](select/select.go) - Using select, context for timeout and cancellation
- [rtb.go](rtb/rtb.go) - Using context for timeouts
- [nlp](nlp) - `nlp` project


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

- `G☺`
- `♡`
- [http.log.gz](_extra/http.log.gz)
- `https://api.github.com/users/tebeka`
- [Slides](_extra/slides.pdf)
- [Unicode](_extra/unicode.pdf)
- [Cost of errors](_extra/cost.png)
- [sherlock.txt](_extra/sherlock.txt)
- [slices](_extra/slices.md)
- [rtb.go](_extra/rtb.go)
- [taxi.tar](https://storage.googleapis.com/353solutions/c/data/taxi.tar)
- [flags](https://github.com/tebeka/flags) - More types for the built-in `flag` package
- [nlp.go](_extra/nlp.go)
- [stemmer.go](_extra/stemmer.go)
- [tokenize_cases.toml](_extra/tokenize_cases.toml)
    - github.com/BurntSushi/toml
