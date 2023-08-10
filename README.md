# go1.21-example-data-race

OOF

Works fine with Go 1.20 but fails on `gotip` (effectively Go 1.21.0).

# Stacktrace

```go
=== RUN   Example1
==================
WARNING: DATA RACE
Read at 0x0001004526e8 by goroutine 7:
  fmt.Printf()
      /Users/me/sdk/gotip/src/fmt/print.go:233 +0x44
  go1_21.Example1.func1()
      /Users/me/code/oleg/go1.21-example-data-race/example_test.go:12 +0x20

Previous write at 0x0001004526e8 by main goroutine:
  testing.runExample.func2()
      /Users/me/sdk/gotip/src/testing/run_example.go:55 +0xb4
  runtime.deferreturn()
      /Users/me/sdk/gotip/src/runtime/panic.go:477 +0x34
  testing.runExamples()
      /Users/me/sdk/gotip/src/testing/example.go:44 +0x1d0
  testing.(*M).Run()
      /Users/me/sdk/gotip/src/testing/testing.go:1927 +0x9c4
  main.main()
      _testmain.go:49 +0x294

Goroutine 7 (running) created at:
  go1_21.Example1()
      /Users/me/code/oleg/go1.21-example-data-race/example_test.go:11 +0x7c
  testing.runExample()
      /Users/me/sdk/gotip/src/testing/run_example.go:63 +0x314
  testing.runExamples()
      /Users/me/sdk/gotip/src/testing/example.go:44 +0x1d0
  testing.(*M).Run()
      /Users/me/sdk/gotip/src/testing/testing.go:1927 +0x9c4
  main.main()
      _testmain.go:49 +0x294
==================
--- PASS: Example1 (0.00s)
str: "abc"
=== RUN   Example2
--- PASS: Example2 (0.10s)
FAIL
FAIL    go1_21  0.341s
FAIL
```
