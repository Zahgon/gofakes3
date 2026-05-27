package main

// Here we are checking to see if a bucket that has never had versioning enabled
// will respond successfully to ListVersions; turns out it will.
//
// Discoveries:
//
//   - S3 responds to ListVersions even if the bucket has never been versioned.
//     We will probably need to fake this in in the GoFakeS3 struct to use the
//     normal bucket listing API when the backend does not implement versioning,
//     but this will probably need to wait until we implement proper pagination.
//
//   - The API returns the _string_ 'null' for the version ID, which the Go SDK
//     happily returns as the *string* value 'null' (yecch!). GoFakeS3 backend
//     implementers should be able to simply return the empty string; GoFakeS3
//     itself should handle this particular bit of jank once and once only.
type S300005ListVersionsWithNeverVersionedBucket struct{}

func (s S300005ListVersionsWithNeverVersionedBucket) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check version length

// We have uploaded two keys, two versions per key, but as we have never
// enabled versioning, we only expect two results

// This is pretty bad... the AWS SDK returns the *STRING* 'null' if there's no version ID.
