package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func find(id int) error {
	if id < 0 {
		return fmt.Errorf("find %d: %w", id, ErrNotFound)
	}
	return nil
}

type HTTPError struct {
	Code int
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http %d", e.Code)
}

func doRequest() error {
	return fmt.Errorf("request: %w", &HTTPError{
		Code: 404,
	})
}

func main() {
	// errors.Is
	err := find(-1)
	fmt.Println(err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("handled not found")
	}

	// errors.As
	err = doRequest()
	var he *HTTPError
	if errors.As(err, &he) {
		fmt.Println("status", he.Code)
	}

	// Unwrap manually
	errx := fmt.Errorf("outer: %w", errors.New("inner"))
	fmt.Println(errors.Unwrap(errx))

	// We can use %w to wrap inner error correctly with chaining outer error
	// we can unwrap check inner errror by using Is, As, Unwrap functions

}
