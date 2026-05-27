package gofakes3

import (
	"net/http"
	"strings"
)

var (
	corsHeaders = []string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"cache-control",
		"Content-Disposition",
		"Content-Encoding",
		"Content-Length",
		"Content-Type",
		"X-Amz-Date",
		"X-Amz-User-Agent",
		"X-CSRF-Token",
		"x-amz-acl",
		"x-amz-content-sha256",
		"x-amz-meta-filename",
		"x-amz-meta-from",
		"x-amz-meta-private",
		"x-amz-meta-to",
		"x-amz-security-token",
		"x-requested-with",
	}
	corsHeadersString = strings.Join(corsHeaders, ", ")
)

type withCORS struct {
	handler http.Handler

	methods string
	origin  string
	headers string
	expose  string
}

func wrapCORS(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func wrapInsecureCORS(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *withCORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
