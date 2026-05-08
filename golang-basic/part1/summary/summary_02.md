# Summary - Questions and Answers (Arrays through Errors)

## 7) Arrays & Slices
### Why we use it
- Arrays give fixed-size, value-type collections; slices give dynamic, reference-like views over arrays.
- `append`, `len`, `cap` let you grow and inspect slices without manual memory management.

### Benefits in real systems
- Slices are the workhorse collection in Go — used for request batches, query results, middleware chains, and more.
- Understanding backing-array sharing prevents subtle mutation bugs across function boundaries.

### What to learn
- Always reassign: `s = append(s, x)` — forgetting the reassign is a common source of lost data.
- Be deliberate about pre-allocating with `make([]T, 0, cap)` when the final size is known.
- Know the difference between `nil` slice and empty slice for JSON serialisation (`null` vs `[]`).

### Optimization impact
- Pre-sized slices reduce allocations and GC pressure in hot loops.
- Using `copy` instead of repeated `append` avoids unnecessary capacity doublings.

---

## 8) Maps
### Why we use it
- Provide O(1) key-value lookups for caches, indexes, counters, and configuration.
- The comma-ok idiom (`v, ok := m[k]`) cleanly separates "missing" from "zero value".

### Benefits in real systems
- Natural fit for lookup tables, deduplication, frequency counting, and in-memory caches.
- `delete` and iteration make runtime data management straightforward.

### What to learn
- Always initialise maps with `make` or a literal — writing to a `nil` map panics.
- Never depend on iteration order; sort keys explicitly when deterministic output matters.
- Default maps are not concurrent-safe; plan for `sync.Mutex` or `sync.Map` early.

### Optimization impact
- Pre-sizing with `make(map[K]V, hint)` reduces rehashing for known cardinalities.
- Replacing slice scans with map lookups can drop O(n) paths to O(1).

---

## 9) Strings & Runes
### Why we use it
- Go strings are UTF-8 bytes; `rune` (int32) represents a single Unicode code point.
- Correct text handling requires distinguishing byte length from character count.

### Benefits in real systems
- Proper rune handling prevents garbled output, wrong truncation, and broken Unicode in APIs serving international users.
- `strings.Builder` enables efficient string concatenation without repeated allocations.

### What to learn
- `len(s)` returns bytes, not characters — use `utf8.RuneCountInString` for logical length.
- Iterate with `range` to get runes; raw byte indexing breaks multi-byte characters.
- Conversions `[]byte(s)` and `string(b)` copy data — avoid in hot paths.

### Optimization impact
- Using `strings.Builder` instead of `+=` concatenation avoids quadratic allocation in loops.
- Byte-level processing where rune safety is guaranteed avoids unnecessary rune decoding overhead.

---

## 10) Structs & Methods
### Why we use it
- Structs group related fields into a single type; methods attach behavior to that type.
- Pointer receivers enable mutation; value receivers guarantee immutability of the original.

### Benefits in real systems
- Structs model domain entities (User, Order, Config) with clear field contracts.
- Methods keep behavior co-located with data, improving discoverability and encapsulation.

### What to learn
- If any method needs `*T`, keep all methods on `*T` for API consistency.
- Value receivers get a copy — mutations inside them do not affect the caller's value.
- Named types (e.g., `type Celsius float64`) can also carry methods, adding domain meaning to primitives.

### Optimization impact
- Pointer receivers avoid copying large structs on every method call.
- Well-designed struct layouts reduce memory padding and improve cache locality.

---

## 11) Packages & Imports
### Why we use it
- Packages organise code into reusable, independently compilable units.
- Exported (uppercase) vs unexported (lowercase) names enforce encapsulation at the package boundary.

### Benefits in real systems
- Clear package boundaries reduce coupling and enable parallel team development.
- Encapsulation hides implementation details, making refactoring safer.

### What to learn
- One directory = one package; folder name and package clause should match by convention.
- Import cycles are compile errors — design dependency direction carefully.
- Blank imports (`_ "pkg"`) register side effects (drivers, codecs); dot imports are rare and discouraged.

### Optimization impact
- Smaller, focused packages compile faster and produce more cacheable build artifacts.
- Good package design reduces the blast radius of changes and speeds up CI pipelines.

---

## 12) Errors — Basics
### Why we use it
- `error` is a built-in interface; the `(T, error)` return pattern is Go's primary error-handling idiom.
- `errors.New` and `fmt.Errorf` create descriptive error values without exceptions or stack unwinding.

### Benefits in real systems
- Explicit error returns force callers to handle failures — no silent swallowing.
- Sentinel errors (`var ErrNotFound = errors.New(...)`) enable reliable programmatic checks.

### What to learn
- Always check `if err != nil` — ignoring errors leads to silent data corruption and hard-to-trace bugs.
- Use `fmt.Errorf` for context-rich messages; reserve `%w` wrapping for Part 2.
- `panic` is not for normal control flow — use `error` returns instead.

### Optimization impact
- Early error returns keep the happy path flat and fast, avoiding unnecessary computation.
- Consistent error patterns reduce debugging time and speed up incident response.

---

## Checklist (Arrays through Errors)
- Can the team explain slice backing-array sharing and when `append` allocates?
- Are maps always initialised before use, with concurrency guards where needed?
- Is string handling rune-aware for any user-facing or internationalised text?
- Do structs use pointer receivers consistently and model domain entities clearly?
- Are packages small, cycle-free, and with a clear exported API surface?
- Does every function that can fail return `error`, and is every `error` checked?

## Answers (Practical View)
- **Slice sharing and append:** Yes, the team should understand that sub-slices share memory and that `append` may or may not allocate. Pre-allocate with `make` when the size is predictable, and always reassign the result of `append`.
- **Map initialisation and concurrency:** Yes, never write to a `nil` map. Use `make` or a literal, and protect concurrent access with `sync.Mutex` or `sync.Map` from day one in concurrent code.
- **Rune-aware string handling:** Yes, use `range` for rune iteration and `utf8.RuneCountInString` for logical length. Avoid raw byte indexing unless the encoding is guaranteed ASCII.
- **Structs and receiver consistency:** Yes, prefer pointer receivers when any method mutates state, and apply the same receiver type across all methods on a type. Model domain concepts as named structs with clear field semantics.
- **Package design:** Yes, keep packages focused on a single responsibility, export only what consumers need, and resolve import cycles by extracting shared types into a separate package.
- **Error handling discipline:** Yes, every fallible function returns `(T, error)`, every call site checks `err != nil`, and error messages include enough context to diagnose failures without a debugger.
