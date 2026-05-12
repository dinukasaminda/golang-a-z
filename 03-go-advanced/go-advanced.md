# Advanced Go — what you will learn

This folder is the **Part 3 companion** for the root curriculum. It covers advanced Go topics: concurrency, cancellation, synchronization, generics, reflection, runtime behavior, profiling, and performance engineering.

Use it the same way as `golang-basic`: open each guide, read the problem first, try to write your own Go file, then compare your work with the snippets.

No `.go` files are included here on purpose. The snippets are reference material; you should create the runnable examples yourself.

---

## Part 3 — topic guides (`part3/`)

| # | Topic | Guide |
|---|-------|-------|
| 1 | Goroutines | [part3/01-goroutines.md](part3/01-goroutines.md) |
| 2 | Channels | [part3/02-channels.md](part3/02-channels.md) |
| 3 | `select` and timeouts | [part3/03-select-and-timeouts.md](part3/03-select-and-timeouts.md) |
| 4 | `context` | [part3/04-context.md](part3/04-context.md) |
| 5 | `sync` primitives | [part3/05-sync-primitives.md](part3/05-sync-primitives.md) |
| 6 | `sync/atomic` and memory model | [part3/06-atomic-and-memory-model.md](part3/06-atomic-and-memory-model.md) |
| 7 | Worker pools | [part3/07-worker-pool.md](part3/07-worker-pool.md) |
| 8 | Pipelines | [part3/08-pipelines.md](part3/08-pipelines.md) |
| 9 | Generics | [part3/09-generics.md](part3/09-generics.md) |
| 10 | Reflection | [part3/10-reflection.md](part3/10-reflection.md) |
| 11 | `unsafe` and cgo | [part3/11-unsafe-and-cgo.md](part3/11-unsafe-and-cgo.md) |
| 12 | Benchmarks and `pprof` | [part3/12-benchmarks-and-pprof.md](part3/12-benchmarks-and-pprof.md) |
| 13 | Race detector | [part3/13-race-detector.md](part3/13-race-detector.md) |
| 14 | `embed` and `go:generate` | [part3/14-embed-and-generate.md](part3/14-embed-and-generate.md) |
| 15 | Escape analysis | [part3/15-escape-analysis.md](part3/15-escape-analysis.md) |
| 16 | Allocations and struct layout | [part3/16-allocations-and-struct-layout.md](part3/16-allocations-and-struct-layout.md) |
| 17 | False sharing and cache lines | [part3/17-false-sharing-and-cache-lines.md](part3/17-false-sharing-and-cache-lines.md) |
| 18 | GC tuning with `GOGC` and `GOMEMLIMIT` | [part3/18-gc-tuning-gogc-gomemlimit.md](part3/18-gc-tuning-gogc-gomemlimit.md) |
| 19 | `GOMAXPROCS` and scheduler | [part3/19-gomaxprocs-and-scheduler.md](part3/19-gomaxprocs-and-scheduler.md) |
| 20 | Profile-guided optimisation | [part3/20-profile-guided-optimisation-pgo.md](part3/20-profile-guided-optimisation-pgo.md) |
| 21 | Lock-free programming | [part3/21-lock-free-programming.md](part3/21-lock-free-programming.md) |

---

## Lesson format

Each guide uses this structure:

1. **Problem** — the real pain this feature solves.
2. **Solution** — the Go concept and how to think about it.
3. **Code snippet** — a practical reference snippet.
4. **Conclusion** — what to remember after typing it yourself.
5. **What improves if you use this** — the engineering benefit.
6. **Final things to know** — quick questions and answers.

Start with goroutines and channels before jumping into performance topics. Most advanced Go bugs come from using powerful tools before the mental model is ready.
