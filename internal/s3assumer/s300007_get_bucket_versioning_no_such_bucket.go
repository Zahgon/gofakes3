package main

// Does GetBucketVersioning return ErrNoSuchBucket when a nonexistent bucket is used?
// Does PutBucketVersioning return ErrNoSuchBucket when a nonexistent bucket is used?
type S300007BucketVersioningNoSuchBucket struct{}

func (s S300007BucketVersioningNoSuchBucket) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBucketVersioning

// PutBucketVersioning
