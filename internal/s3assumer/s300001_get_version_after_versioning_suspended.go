package main

// It's not clear from the docs what S3 does when versioning has been enabled,
// then suspended, then you request a version ID that exists.
//
// Turns out it continues to work just fine.
//
// This script also revealed that a bucket that has never had versioning will
// return empty strings for Status and MFADelete.
type S300001GetVersionAfterVersioningSuspended struct{}

func (t *S300001GetVersionAfterVersioningSuspended) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: defer delete object
