package gofakes3

import (
	"io"
)

func parseClampedInt(in string, defaultValue, min, max int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadAll is a fakeS3-centric replacement for ioutil.ReadAll(), for use when
// the size of the result is known ahead of time. It is considerably faster to
// preallocate the entire slice than to allow growslice to be triggered
// repeatedly, especially with larger buffers.
//
// It also reports S3-specific errors in certain conditions, like
// ErrIncompleteBody.
func ReadAll(r io.Reader, size int64) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
