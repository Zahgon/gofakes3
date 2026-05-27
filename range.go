package gofakes3

import (
	"net/http"
)

type ObjectRange struct {
	Start, Length int64
}

func (o *ObjectRange) writeHeader(sz int64, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type ObjectRangeRequest struct {
	Start, End int64
	FromEnd    bool
}

const RangeNoEnd = -1

func (o *ObjectRangeRequest) Range(size int64) (*ObjectRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If no end is specified, range extends to end of the file.

// If no start is specified, end specifies the range start relative
// to the end of the file.

// parseRangeHeader parses a single byte range from the Range header.
//
// Amazon S3 doesn't support retrieving multiple ranges of data per GET request:
// https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html
func parseRangeHeader(s string) (*ObjectRangeRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
