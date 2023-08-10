package main

import (
	"fmt"
	"time"
)

func Example1() {
	fmt.Printf("num1: %d\n", 123)

	go func() {
		fmt.Printf("str: %q\n", "abc")
	}()

	// Output:
	// num1: 123
}

func Example2() {
	fmt.Printf("num: %d\n", 123)

	go func() {
		fmt.Printf("str: %q\n", "abc")
	}()

	time.Sleep(100 * time.Millisecond)

	// Output:
	// num: 123
	// str: "abc"
}
