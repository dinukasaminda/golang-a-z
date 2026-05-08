# Part 2 — 09 — `encoding/json`

**Goal:** Marshal structs to JSON and unmarshal JSON into Go values using **struct tags**.

## What you are learning

- **`json.Marshal`** → `[]byte` + `error`; **`json.MarshalIndent`** for pretty printing.
- **`json.Unmarshal`** fills structs, maps, slices, or **`any`**.
- Struct tag **`json:"field_name"`** renames keys; **`json:"-"`** omits; **`json:"name,omitempty"`** omits zero values.
- **`json.Number`** preserves numbers without float64 loss when decoding into **`map[string]any`**.

## Try yourself first

1. Define a struct with `json` tags and marshal it to a string (print).
2. Unmarshal JSON with unknown field — observe extra fields are **ignored** by default on structs.
3. Decode `{"count": 42}` into a `map[string]any` and read the number.

## Reference snippets

### Marshal / Unmarshal struct

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	u := User{Name: "Ada", Age: 0}
	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var v User
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}
	fmt.Println(v)
}
```

### `omitempty` and ignored unknown fields

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Host string `json:"host"`
}

func main() {
	raw := `{"host":"localhost","extra":true}`
	var c Config
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
}
```

### `map[string]any` and `json.Number`

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	dec := json.NewDecoder(bytes.NewReader([]byte(`{"n":1.5,"i":42}`)))
	dec.UseNumber()
	var m map[string]any
	if err := dec.Decode(&m); err != nil {
		panic(err)
	}
	i, err := m["i"].(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
```

## Gotchas

- Only **exported** struct fields participate in JSON.
- **`time.Time`** marshals to RFC3339 string by default when embedded in structs.

## Compare

With `Age: 0` and `omitempty`, marshalled JSON should often **omit** `"age"` (unless you set Age non-zero).
