# Part 3 — 10 — Reflection

**Goal:** Inspect types and values at runtime with `reflect`.

## Problem

Some libraries must work with unknown struct types, such as JSON encoders, validators, ORMs, and dependency injection tools.

## Solution

Use `reflect.TypeOf` and `reflect.ValueOf` to inspect runtime shape. Keep reflection at boundaries, not inside normal business logic.

## Code snippet

```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Ada", Age: 36}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, v.Field(i))
	}
}
```

## Conclusion

Reflection trades compile-time clarity for runtime flexibility. Use it deliberately.

## What improves if you use this

You can build generic tooling around arbitrary user-defined types.

## Final things to know

**Q: Is reflection fast?**  
A: It is slower and more complex than direct code.

**Q: Can reflection modify values?**  
A: Yes, but only addressable and settable values.

**Q: Where is reflection common?**  
A: Serialization, validation, testing helpers, and frameworks.
