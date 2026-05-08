# Go basics — what you will learn

This folder is the **entry point** for learning Go in this repository. It aligns with **Part 1 (Go Language Fundamentals)** and **Part 2 (Intermediate Go)** from the root curriculum.

**How to use the topic guides:** open each linked `.md` file below. Read the explanation, then **write your own** `main.go` (or package) from scratch. Use the **reference snippets** in that file only after you try, or to check your work — they are not checked in as `.go` files on purpose.

---

## Part 1 — topic guides (`part1/`)

| # | Topic | Guide |
|---|--------|--------|
| 1 | Hello World & toolchain | [part1/01-hello-world-and-toolchain.md](part1/01-hello-world-and-toolchain.md) |
| 2 | Variables & constants | [part1/02-variables-and-constants.md](part1/02-variables-and-constants.md) |
| 3 | Primitive types | [part1/03-primitive-types.md](part1/03-primitive-types.md) |
| 4 | Control flow | [part1/04-control-flow.md](part1/04-control-flow.md) |
| 5 | Functions | [part1/05-functions.md](part1/05-functions.md) |
| 6 | Pointers | [part1/06-pointers.md](part1/06-pointers.md) |
| 7 | Arrays & slices | [part1/07-arrays-and-slices.md](part1/07-arrays-and-slices.md) |
| 8 | Maps | [part1/08-maps.md](part1/08-maps.md) |
| 9 | Strings & runes | [part1/09-strings-and-runes.md](part1/09-strings-and-runes.md) |
| 10 | Structs & methods | [part1/10-structs-and-methods.md](part1/10-structs-and-methods.md) |
| 11 | Packages & imports | [part1/11-packages-and-imports.md](part1/11-packages-and-imports.md) |
| 12 | Errors — basics | [part1/12-errors-basics.md](part1/12-errors-basics.md) |

---

## Part 2 — topic guides (`part2/`)

| # | Topic | Guide |
|---|--------|--------|
| 1 | Interfaces | [part2/01-interfaces.md](part2/01-interfaces.md) |
| 2 | Type assertions & type switches | [part2/02-type-assertions-and-type-switches.md](part2/02-type-assertions-and-type-switches.md) |
| 3 | Struct embedding | [part2/03-struct-embedding.md](part2/03-struct-embedding.md) |
| 4 | Error wrapping | [part2/04-error-wrapping.md](part2/04-error-wrapping.md) |
| 5 | `defer`, `panic`, `recover` | [part2/05-defer-panic-recover.md](part2/05-defer-panic-recover.md) |
| 6 | `io` & `bufio` | [part2/06-io-and-bufio.md](part2/06-io-and-bufio.md) |
| 7 | Files & `os` | [part2/07-files-and-os.md](part2/07-files-and-os.md) |
| 8 | `time` & duration | [part2/08-time-and-duration.md](part2/08-time-and-duration.md) |
| 9 | `encoding/json` | [part2/09-encoding-json.md](part2/09-encoding-json.md) |
| 10 | Flags & CLI | [part2/10-flags-and-cli.md](part2/10-flags-and-cli.md) |
| 11 | Modules & workspaces | [part2/11-modules-and-workspaces.md](part2/11-modules-and-workspaces.md) |

---

## Part 1 — Language fundamentals (summary table)

These topics teach you to read and write correct Go: syntax, types, control flow, and the building blocks every program uses.

| # | Topic | What you learn (short explanation) |
|---|--------|-------------------------------------|
| 1 | **Hello World & toolchain** | How to install Go, use `go run` and `go build`, and understand the layout of a minimal program. |
| 2 | **Variables & constants** | Declaring values with `var` and `:=`, constants with `const`, and how **zero values** initialise variables before you assign anything. |
| 3 | **Primitive types** | `int`, `float`, `bool`, `string`, and how **explicit conversions** work (Go does not implicitly widen or narrow types). |
| 4 | **Control flow** | `if`, `for` (the only loop in Go), `switch`, and writing clear branching logic without unnecessary complexity. |
| 5 | **Functions** | Parameters, return values, **multiple return values**, and **named results** — idiomatic for errors and small tuples. |
| 6 | **Pointers** | Address-of (`&`), dereference (`*`), when Go passes by value vs when you need a pointer, and avoiding nil pointer mistakes. |
| 7 | **Arrays & slices** | Fixed-size arrays vs **slices** (dynamic views); `len`, `cap`, `append`, and how slices share underlying arrays. |
| 8 | **Maps** | Key–value maps, existence checks with the two-value form, and iteration order (not guaranteed). |
| 9 | **Strings & runes** | UTF-8 in Go: `string` as bytes, **`rune`** as a Unicode code point, and using `unicode/utf8` when you need character-level logic. |
| 10 | **Structs & methods** | Grouping data with `struct`, defining **methods** with receivers (value vs pointer), and the start of object-like behaviour without classes. |
| 11 | **Packages & imports** | Organising code into packages, `import` paths, exported vs unexported names (capital letter = exported), and `go mod init`. |
| 12 | **Errors — basics** | The `error` interface, creating errors with `errors.New`, returning `(T, error)` from functions, and checking errors with `if err != nil`. |

---

## Part 2 — Intermediate Go

These topics move you from “correct Go” to **idiomatic Go**: interfaces, composition, richer errors, and everyday standard-library use.

| # | Topic | What you learn (short explanation) |
|---|--------|-------------------------------------|
| 1 | **Interfaces** | Defining behaviour-only types; **implicit implementation** (no `implements` keyword); small interfaces and the **accept interfaces, return structs** idea. |
| 2 | **Type assertions & type switches** | Extracting concrete types from an `interface{}` or a small interface; safe two-value form vs panic-prone single-value form. |
| 3 | **Struct embedding** | Embedding one struct inside another for **composition** instead of classical inheritance; promoted fields and methods. |
| 4 | **Error wrapping** | `fmt.Errorf` with `%w`, **`errors.Is`** and **`errors.As`** for comparing and inspecting errors in a chain. |
| 5 | **`defer`, `panic`, `recover`** | `defer` for cleanup and ordering; when **panic** is for truly exceptional cases; **`recover`** only inside deferred functions — not for normal control flow. |
| 6 | **`io` & `bufio`** | The `io.Reader` / `io.Writer` abstractions; buffered I/O with `bufio` for efficient reads and writes. |
| 7 | **Files & `os`** | Opening files, reading and writing paths, environment variables, and basic filesystem operations. |
| 8 | **`time` & duration** | `time.Time`, `time.Duration`, monotonic clocks, formatting and parsing, and scheduling with `time.After` / `time.Ticker` basics. |
| 9 | **`encoding/json`** | Marshalling and unmarshalling structs, tags like `` `json:"field"` ``, and dealing with unknown or optional fields. |
| 10 | **Flags & CLI** | Building small command-line tools with the `flag` package and `os.Args`. |
| 11 | **Modules & workspaces** | **`go.mod`** for dependencies and versions, **`go.sum`** for integrity, and **`go.work`** for multi-module local development. |

---

## How this folder will grow

Each topic has a **guide** under `part1/` or `part2/` (this page links to them). You type the Go yourself; the guides hold explanations and **reference snippets** only.

Later, each topic can also get a **module folder** (for example `01-hello-world/`) with checked-in `main.go`, tests, Mermaid diagrams, and `exercises.md` if you want a fuller repo layout.

Start with **Part 1** in order; **Part 2** assumes you are comfortable with Part 1.

For the full learning path (CS, architecture, distributed systems, and more), see the root [`README.md`](../README.md).
