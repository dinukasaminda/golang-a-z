# Part 2 — 11 — Modules & workspaces

**Goal:** Manage dependencies with **`go.mod`** / **`go.sum`** and work on multiple modules locally with **`go.work`**.

## What you are learning

- **`go mod init example.com/mymodule`** creates **`go.mod`** with the module path and Go version.
- **`go get example.com/pkg@v1.2.3`** adds a **require** line; **`go mod tidy`** removes unused requires.
- **`go.sum`** records **checksums** of module contents — commit it for reproducible builds.
- **Semantic versioning:** `v0`, `v1` rules; major version `v2+` often needs **`/v2`** in the import path (advanced rule — look up when you publish libraries).
- **`go.work`** lists **multiple module roots** so local replaces work without publishing (`use ./foo`, `use ./bar`).

## Try yourself first

1. Create a new module, add a tiny dependency (e.g. a popular small package), run `go mod tidy`, inspect `go.mod` and `go.sum`.
2. Run **`go list -m all`** and read the output.
3. (Optional) Create two sibling modules and a **`go.work`** that `use`s both; import one from the other.

## Reference snippets

### `go.mod` (example)

```text
module example.com/learn

go 1.22

require (
	github.com/google/uuid v1.6.0
)
```

### Adding a dependency (commands, not Go code)

```text
go get github.com/google/uuid@v1.6.0
go mod tidy
```

### Minimal `go.work` (two local modules)

```text
go 1.22

use (
	./service-a
	./service-b
)
```

### `replace` directive in `go.mod` (point at local path)

```text
replace example.com/lib => ../lib
```

## Gotchas

- **Version `latest`** is not always what you want in production — pin versions in application repos.
- **`GOPATH` mode** is legacy; modern Go is **module mode** by default.

## Compare

After `go mod tidy`, `go.mod` should list only modules you actually import; `go.sum` should grow with checksum lines.
