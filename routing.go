package gofakes3

import (
	"net/http"
)

// routeBase is a http.HandlerFunc that dispatches top level routes for
// GoFakeS3.
//
// URLs are assumed to break down into two common path segments, in the
// following format:
//
//	/<bucket>/<object>
//
// The operation for most of the core functionality is built around HTTP
// verbs, but outside the core functionality, the clean separation starts
// to degrade, especially around multipart uploads.
func (g *GoFakeS3) routeBase(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// x-amz-id-2 is 48 bytes of random stuff

// routeObject oandles URLs that contain both a bucket path segment and an
// object path segment.
func (g *GoFakeS3) routeObject(bucket, object string, w http.ResponseWriter, r *http.Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// routeBucket handles URLs that contain only a bucket path segment, not an
// object path segment.
func (g *GoFakeS3) routeBucket(bucket string, w http.ResponseWriter, r *http.Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// routeMultipartUploadBase operates on routes that contain '?uploads' in the
// query string. These routes may or may not have a value for bucket or object;
// this is validated and handled in the target handler functions.
func (g *GoFakeS3) routeMultipartUploadBase(bucket, object string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// routeVersioningBase operates on routes that contain '?versioning' in the
// query string. These routes may or may not have a value for bucket; this is
// validated and handled in the target handler functions.
func (g *GoFakeS3) routeVersioning(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// routeVersions operates on routes that contain '?versions' in the query string.
func (g *GoFakeS3) routeVersions(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// routeVersion operates on routes that contain '?versionId=<id>' in the
// query string.
func (g *GoFakeS3) routeVersion(bucket, object string, versionID VersionID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// routeMultipartUpload operates on routes that contain '?uploadId=<id>' in the
// query string.
func (g *GoFakeS3) routeMultipartUpload(bucket, object string, uploadID UploadID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func versionFromQuery(qv []string) string {
	_ = "STUB: not implemented"
	// The versionId subresource may be the string 'null'; this has been
	// observed coming in via Boto. The S3 documentation for the "DELETE
	// object" endpoint describes a 'null' version explicitly, but we don't
	// want backend implementers to have to special-case this string, so
	// let's hide it in here:
	return ""
}
