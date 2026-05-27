package main

func testName(test Test) string { _ = "STUB: not implemented"; return "" }

type Test interface {
	Run(ctx *Context) error
}
