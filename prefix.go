package gofakes3

import (
	"net/url"
)

type Prefix struct {
	HasPrefix bool
	Prefix    string

	HasDelimiter bool
	Delimiter    string
}

func prefixFromQuery(query url.Values) Prefix { _ = "STUB: not implemented"; return *new(Prefix) }

func NewPrefix(prefix, delim *string) (p Prefix) { _ = "STUB: not implemented"; return *new(Prefix) }

func NewFolderPrefix(prefix string) (p Prefix) { _ = "STUB: not implemented"; return *new(Prefix) }

// FilePrefix returns the path portion, then the remaining portion of the
// Prefix if the Delimiter is "/". If the Delimiter is not set, or not "/",
// ok will be false.
//
// For example:
//
//	/foo/bar/  : path: /foo/bar  remaining: ""
//	/foo/bar/b : path: /foo/bar  remaining: "b"
//	/foo/bar   : path: /foo      remaining: "bar"
func (p Prefix) FilePrefix() (path, remaining string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// PrefixMatch checks whether key starts with prefix. If the prefix does not
// match, nil is returned.
//
// It is a best-effort attempt to implement the prefix/delimiter matching found
// in S3.
//
// To check whether the key belongs in Contents or CommonPrefixes, compare the
// result to key.
func (p Prefix) Match(key string, match *PrefixMatch) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// If there is no prefix in the search, the match is the prefix:

// If the request does not contain a delimiter, prefix matching is a
// simple string prefix:

// Delimited + Prefix matches, for example:
//	 $ aws s3 ls s3://my-bucket/
//	                            PRE AWSLogs/
//	 $ aws s3 ls s3://my-bucket/AWSLogs
//	                            PRE AWSLogs/
//	 $ aws s3 ls s3://my-bucket/AWSLogs/
//	                            PRE 260839334643/
//	 $ aws s3 ls s3://my-bucket/AWSLogs/2608
//	                            PRE 260839334643/

// If the key exactly matches the prefix, but only up to a delimiter,
// AWS appends the delimiter to the result:
//	 $ aws s3 ls s3://my-bucket/AWSLogs
//	                            PRE AWSLogs/

func (p Prefix) String() string { _ = "STUB: not implemented"; return "" }

type PrefixMatch struct {
	// Input key passed to PrefixMatch.
	Key string

	// CommonPrefix indicates whether this key should be returned in the bucket
	// contents or the common prefixes part of the "list bucket" response.
	CommonPrefix bool

	// The longest matched part of the key.
	MatchedPart string
}

func (match *PrefixMatch) AsCommonPrefix() CommonPrefix {
	_ = "STUB: not implemented"
	return *new(CommonPrefix)
}
