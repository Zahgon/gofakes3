package main

// Does GetBucketLocation return ErrNoSuchBucket when a nonexistent bucket is used?
type S300006GetBucketLocationNoSuchBucket struct{}

func (s S300006GetBucketLocationNoSuchBucket) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check version length
