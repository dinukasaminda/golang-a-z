package main

import "fmt"

type Metadata struct {
	ID   int
	Tags []string
}

// Promoted Fields
type Document struct {
	Metadata
	Title string
}

// Promoted Method
type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("Log: ", msg)
}

type Service struct {
	Logger
	Name string
}

// Name clash (Outer wins; inner must be qualified)

type Inner struct{ Name string }
type Outer struct {
	Inner
	Name string
}

func main() {
	// Promoted Fields
	doc := &Document{
		Title: "Document 01",
		Metadata: Metadata{
			ID:   10001,
			Tags: []string{"Tag1", "Tag2"},
		},
	}
	doc.ID = 20001
	fmt.Println(doc)
	fmt.Println(doc.ID)

	// Promoted Method
	s := Service{
		Name:   "API",
		Logger: Logger{},
	}
	s.Log("Hello!")

	// Name clash
	o := Outer{Inner: Inner{Name: "inner"}, Name: "outer"}
	fmt.Println(o.Name)
	fmt.Println(o.Inner.Name)

	//You should access embedded fields both as d.Metadata.ID and d.ID when unambiguous.
}
