package main

// It is not clear from the docs at
// https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html
// what happens if the version or object does not exist. The implication is
// that no error is returned; this is consistent with DeleteObject.
//
// This test confirms that this is indeed the case.
//
// It also highlighted that the delete methods return a 204 status, not a 200.
type S300002DeleteNonexistentVersion struct{}

func (s S300002DeleteNonexistentVersion) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Create two versions so we guarantee that even though we are deleting one of
// them, the object still exists:

// We need to use a real version ID because they have significance on AWS even
// though the meaning is opaque:

// Delete should succeed the first time. An object version will remain after this.

// Now we should get the answer about what S3 actually does when you try to delete
// a version that is known not to exist!

// DEBUG: Response s3/DeleteObject Details:
// ---[ RESPONSE ]--------------------------------------
// HTTP/1.1 204 No Content
