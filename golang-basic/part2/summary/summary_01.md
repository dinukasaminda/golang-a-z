# Summary - Questions and Answers (Interfaces through I/O)

## 1) Interfaces
### Why we use it
- Define behavior-only contracts without tying code to one concrete type.
- Let types satisfy contracts implicitly when their method set matches.

### Benefits in real systems
- Enables flexible APIs such as accepting `io.Reader` instead of `*os.File`.
- Makes code easier to test with small fakes and focused mocks.

### What to learn
- Prefer small interfaces that describe one capability.
- Accept interfaces and return concrete types when designing functions.
- Understand the nil interface vs nil concrete value gotcha.

### Optimization impact
- Smaller interfaces reduce coupling and make refactoring cheaper.
- Interface boundaries let hot paths swap implementations without changing callers.

---

## 2) Type Assertions & Type Switches
### Why we use it
- Recover concrete types from `any` or wider interfaces when static type information is lost.
- Handle dynamic values safely at boundaries like JSON decoding, plugins, and generic APIs.

### Benefits in real systems
- Prevents panics from unsafe assumptions about runtime values.
- Keeps boundary-handling logic explicit and readable.

### What to learn
- Use the two-value assertion form: `v, ok := x.(T)`.
- Use type switches when several concrete types are valid.
- Avoid one-value assertions on untrusted or external data.

### Optimization impact
- Safe checks reduce crash risk in request and data-processing paths.
- Clear type-switch logic lowers debugging time when dynamic input changes.

---

## 3) Struct Embedding
### Why we use it
- Compose structs by embedding fields and promoting methods/fields.
- Reuse behavior through composition rather than inheritance.

### Benefits in real systems
- Useful for wrappers, decorators, and layered types such as HTTP middleware helpers.
- Keeps shared behavior close to the type that provides it.

### What to learn
- Embedded fields are promoted only when names are unambiguous.
- Use explicit named fields when promotion makes the API hard to read.
- Be careful with embedded pointers because nil embedded values can panic.

### Optimization impact
- Composition avoids deep inheritance-style designs and reduces maintenance cost.
- Clear type layering makes API changes easier to isolate.

---

## 4) Error Wrapping
### Why we use it
- Preserve original errors while adding context with `fmt.Errorf("...: %w", err)`.
- Detect sentinel or custom errors through the wrap chain.

### Benefits in real systems
- Produces useful operational messages while keeping programmatic checks reliable.
- Helps callers decide whether to retry, report, ignore, or translate failures.

### What to learn
- Use `%w` when wrapping an error; `%v` only formats text.
- Use `errors.Is` for sentinel errors and `errors.As` for custom error types.
- Do not compare wrapped errors with `==`.

### Optimization impact
- Context-rich error chains speed up incident diagnosis.
- Reliable classification reduces brittle string matching in production code.

---

## 5) `defer`, `panic`, `recover`
### Why we use it
- Schedule cleanup actions that must run when a function exits.
- Understand panic recovery for exceptional boundaries, not normal error flow.

### Benefits in real systems
- Ensures files, locks, spans, and temporary resources are released consistently.
- Keeps cleanup close to acquisition, reducing leak risk.

### What to learn
- Deferred calls run in LIFO order when the surrounding function returns.
- Use `defer` after successful resource acquisition.
- `recover` only works inside a deferred function while a panic is unwinding.
- Prefer `error` returns over `panic` for expected failures.

### Optimization impact
- Correct cleanup prevents resource leaks and long-running process degradation.
- Scoped defer usage avoids accumulating too many deferred calls in large loops.

---

## 6) `io` & `bufio`
### Why we use it
- Stream bytes through common interfaces: `io.Reader` and `io.Writer`.
- Add buffering for efficient line and delimiter-based text handling.

### Benefits in real systems
- Standardizes data flow across files, network connections, buffers, and strings.
- Makes code composable with helpers like `io.Copy`, `io.ReadAll`, and `bufio.Scanner`.

### What to learn
- Treat `io.EOF` as normal end-of-input.
- Use `io.ReadAll` only when input size is safe to hold in memory.
- Check `Scanner.Err()` after scanner loops.
- Use `bufio.Reader` for huge tokens or delimiter-based reads beyond scanner limits.

### Optimization impact
- Streaming avoids loading large payloads into memory.
- Buffered reads/writes reduce syscall overhead and improve throughput.

---

## Checklist (Interfaces through I/O)
- Are interfaces small, behavior-focused, and accepted at package boundaries?
- Are dynamic values handled with safe assertions or type switches?
- Is struct embedding used for composition only where promotion stays readable?
- Are wrapped errors checked with `errors.Is` or `errors.As` instead of strings or `==`?
- Are cleanup actions deferred immediately after successful resource acquisition?
- Is I/O streamed or buffered appropriately instead of reading unbounded data into memory?

## Answers (Practical View)
- **Small interface design:** Yes, define interfaces around one behavior such as `Read`, `Write`, or `String`. Accept interfaces at boundaries and return concrete values so callers keep full type information.
- **Safe dynamic typing:** Yes, use `v, ok := x.(T)` or a type switch for dynamic input. Avoid panicking assertions unless the program invariant is truly guaranteed.
- **Readable embedding:** Yes, use embedding to compose behavior when the promoted API is obvious. Prefer a named field when explicit access would make the design clearer.
- **Error wrapping discipline:** Yes, wrap with `%w` when preserving the cause matters, then detect with `errors.Is` or `errors.As`. Add context at each layer without losing machine-checkable error identity.
- **Cleanup with defer:** Yes, defer cleanup immediately after successful open/create/lock operations. Use panic/recover only at exceptional boundaries, not as routine control flow.
- **I/O strategy:** Yes, stream large or unknown-size data with `io.Reader` and `io.Writer`, buffer when useful, and reserve `io.ReadAll` for bounded inputs.
