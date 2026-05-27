package main

// It is not clear from the docs at
// https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html
// what happens if the object does not exist. The implication is
// that no error is returned; this is consistent with DeleteObject.
//
// This test confirms that this is indeed the case.
type S300003DeleteVersionFromNonexistentObject struct{}

func (s S300003DeleteVersionFromNonexistentObject) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to use a real version ID because they have significance on AWS even
// though the meaning is opaque:

// Delete should succeed the first time. No versions for the object remain and
// the object should no longer exist.

// Now we should get the answer about what S3 actually does when you try to delete
// a version for an object that is known not to exist:

// DEBUG: Response s3/DeleteObject Details:
// ---[ RESPONSE ]--------------------------------------
// HTTP/1.1 204 No Content
