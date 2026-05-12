# Summary - Questions and Answers (Files through Modules)

## 7) Files & `os`
### Why we use it
- Create, open, read, write, and remove files from Go programs.
- Access process arguments and environment variables.

### Benefits in real systems
- Supports config loading, temp-file workflows, command-line tooling, and local persistence.
- Makes filesystem failures explicit through normal `error` handling.

### What to learn
- Use `os.ReadFile` and `os.WriteFile` for simple whole-file operations.
- Use `os.Open`, `os.Create`, or `os.OpenFile` when streaming or flags matter.
- Close files after successful opens and clean up temp paths with `os.Remove` or `os.RemoveAll`.
- Use `filepath` for cross-platform path joining.

### Optimization impact
- Choosing streaming vs whole-file helpers prevents memory spikes.
- Correct close and cleanup behavior avoids file descriptor leaks and stale temp files.

---

## 8) `time` & Duration
### Why we use it
- Represent instants, elapsed time, scheduling, timeouts, and repeated ticks.
- Parse and format timestamps for APIs, logs, and storage.

### Benefits in real systems
- Enables reliable deadline, retry, cache-expiry, and audit-log behavior.
- Makes time zone and duration handling explicit.

### What to learn
- Use `time.Duration` literals like `5 * time.Second`.
- Use Go's reference layout model for `Parse` and `Format`.
- Prefer RFC3339 for API timestamps unless a domain requires another format.
- Stop timers and tickers when finished.

### Optimization impact
- Correct timer/ticker cleanup prevents goroutine and runtime resource leaks.
- Clear duration handling avoids retry storms, slow timeouts, and scheduling bugs.

---

## 9) `encoding/json`
### Why we use it
- Marshal Go values into JSON and unmarshal JSON into structs, maps, slices, or `any`.
- Control wire names and optional fields with struct tags.

### Benefits in real systems
- Powers HTTP APIs, config files, logs, queues, and external integrations.
- Gives a standard, well-tested conversion path between Go types and JSON payloads.

### What to learn
- Only exported struct fields participate in JSON.
- Use tags like `json:"name"`, `json:"-"`, and `json:"name,omitempty"`.
- Unknown JSON fields are ignored by default when decoding into structs.
- Use `Decoder.UseNumber()` when numeric precision matters in `map[string]any`.

### Optimization impact
- Struct decoding is clearer and safer than repeated dynamic map assertions.
- Avoiding unnecessary pretty-printing and large intermediate payloads helps hot API paths.

---

## 10) Flags & CLI
### Why we use it
- Parse command-line options using the standard `flag` package.
- Separate flags from positional arguments.

### Benefits in real systems
- Makes small operational tools, demos, and admin commands configurable.
- Keeps CLI defaults and help text close to the code that uses them.

### What to learn
- Define flags before calling `flag.Parse()`.
- `flag.String`, `flag.Int`, and similar helpers return pointers.
- Use `flag.Args()` for remaining positional arguments after parsing.
- `os.Args[0]` is the program name; user-provided args start at index 1.

### Optimization impact
- Simple standard-library CLIs avoid unnecessary dependencies for small tools.
- Clear flag parsing reduces misconfiguration and support/debug time.

---

## 11) Modules & Workspaces
### Why we use it
- Manage dependencies and module identity with `go.mod` and `go.sum`.
- Work across multiple local modules with `go.work`.

### Benefits in real systems
- Reproducible dependency resolution across developer machines and CI.
- Easier local development when applications depend on sibling modules.

### What to learn
- `go mod init` creates the module definition.
- `go get module@version` adds or changes dependency requirements.
- `go mod tidy` removes unused requirements and adds missing ones.
- Commit `go.sum` for checksum-based reproducibility.
- Use semantic versioning rules carefully, especially `v2+` import paths.

### Optimization impact
- Clean module files reduce dependency drift and CI surprises.
- Workspaces speed local multi-module iteration without publishing temporary versions.

---

## Checklist (Files through Modules)
- Are file operations using the right API for the data size and access pattern?
- Are file handles closed and temporary files/directories cleaned up?
- Are timestamps parsed, formatted, and stored with clear time zone expectations?
- Are timers and tickers stopped when no longer needed?
- Are JSON structs exported and tagged to match the expected wire format?
- Are CLI flags defined before parsing, with sensible defaults and positional args handled?
- Are `go.mod` and `go.sum` tidy, committed, and pinned to intentional dependency versions?

## Answers (Practical View)
- **File API choice:** Yes, use `ReadFile` and `WriteFile` for small bounded files, and stream through `*os.File` or `io.Reader`/`io.Writer` for larger data. Always handle returned errors.
- **Resource cleanup:** Yes, call `defer f.Close()` after a successful open and remove temporary paths when the workflow is done. Close or flush before reading data that was just written.
- **Time correctness:** Yes, use explicit layouts and time zones, favor RFC3339 for API values, and stop tickers or timers to avoid leaks.
- **JSON contracts:** Yes, export fields that need to be encoded, use tags for stable API names, and choose structs over `map[string]any` when the schema is known.
- **CLI reliability:** Yes, define flags before `flag.Parse()`, read pointer results after parsing, and use `flag.Args()` for remaining inputs.
- **Module hygiene:** Yes, keep `go.mod` and `go.sum` committed and tidy. Pin production dependencies intentionally, and use `go.work` for local multi-module development.
