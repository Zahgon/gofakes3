package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3afero"
)

const usage = `
`

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type fakeS3Flags struct {
	host            string
	backendKind     string
	initialBucket   string
	fixedTimeStr    string
	noIntegrity     bool
	hostBucket      bool
	hostBucketBases HostList
	autoBucket      bool
	insecureCORS    bool
	quiet           bool

	boltDb              string
	directFsPath        string
	directFsMeta        string
	directFsBucket      string
	directFsCreatePaths bool

	fsPath        string
	fsMeta        string
	fsCreatePaths bool

	debugCPU  string
	debugHost string
}

func (f *fakeS3Flags) attach(flagSet *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Logging

// Backend specific:

// Debugging:

// Deprecated:

func (f *fakeS3Flags) fsPathFlags() (flags s3afero.FsFlags) {
	_ = "STUB: not implemented"
	return *new(s3afero.FsFlags)
}

func (f *fakeS3Flags) directFsPathFlags() (flags s3afero.FsFlags) {
	_ = "STUB: not implemented"
	return *new(s3afero.FsFlags)
}

func (f *fakeS3Flags) timeOptions() (source gofakes3.TimeSource, skewLimit time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(gofakes3.TimeSource), *new(time.Duration), nil
}

func debugServer(host string) { _ = "STUB: not implemented"; return }

func run() error { _ = "STUB: not implemented"; return nil }

func listenAndServe(addr string, handler http.Handler) error { _ = "STUB: not implemented"; return nil }

func profile(values fakeS3Flags) (func(), error) { _ = "STUB: not implemented"; return nil, nil }

type HostList struct {
	Values []string
}

func (sl HostList) String() string { _ = "STUB: not implemented"; return "" }

func (sl HostList) Type() string { _ = "STUB: not implemented"; return "" }

func (sl *HostList) Set(s string) error { _ = "STUB: not implemented"; return nil }
