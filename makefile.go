//go:build tools
// +build tools

// Run this script like so:
//
//	go run makefile.go <cmd> <args>...
//

package main

import (
	"log"
	"os/exec"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func runBuildDocker() error { _ = "STUB: not implemented"; return nil }

func runRelease(args []string) error { _ = "STUB: not implemented"; return nil }

func runBuildRelease() error { _ = "STUB: not implemented"; return nil }

// runCover collects true code coverage for all packages in gofakes3.
// It does so by running 'go test' for each child package (enumerated by
// 'go list ./...') with the '-coverpkg' flag, populated with the same
// 'go list'.
func runCover(args []string) error { _ = "STUB: not implemented"; return nil }

func command(name string, args ...string) *exec.Cmd { _ = "STUB: not implemented"; return nil }

func goList() (pkgs []string) { _ = "STUB: not implemented"; return nil }
