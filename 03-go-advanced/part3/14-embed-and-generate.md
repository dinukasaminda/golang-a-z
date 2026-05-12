# Part 3 — 14 — `embed` and `go:generate`

**Goal:** Embed static files and automate code generation steps.

## Problem

Programs often need templates, SQL files, migrations, or generated code. Keeping these steps manual causes missing files and stale output.

## Solution

Use `//go:embed` for static assets and `//go:generate` comments to document repeatable generation commands.

## Code snippet

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed templates/*
var templates embed.FS

func main() {
	b, err := templates.ReadFile("templates/welcome.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
```

Generation comment example:

```go
//go:generate stringer -type=Status
```

## Conclusion

`embed` makes deployment simpler; `go:generate` makes build-adjacent steps discoverable.

## What improves if you use this

Your binaries can carry required assets, and generation workflows become repeatable.

## Final things to know

**Q: Does `go generate` run automatically during `go build`?**  
A: No. You run it explicitly.

**Q: Can `embed` include files outside the package directory?**  
A: No.

**Q: What imports `embed` if only directives use it?**  
A: Use `_ "embed"` when you do not reference `embed.FS` directly.
