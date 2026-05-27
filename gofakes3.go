package gofakes3

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"time"
)

// GoFakeS3 implements HTTP handlers for processing S3 requests and returning
// S3 responses.
//
// Logic is delegated to other components, like Backend or uploader.
type GoFakeS3 struct {
	requestID uint64

	storage   Backend
	versioned VersionedBackend

	wrapCORS                func(h http.Handler) http.Handler // WithInsecureCORS
	timeSource              TimeSource                        // WithTimeSource
	timeSkew                time.Duration                     // WithTimeSkewLimit
	metadataSizeLimit       int                               // WithMetadataSizeLimit
	integrityCheck          bool                              // WithIntegrityCheck
	failOnUnimplementedPage bool                              // WithUnimplementedPageError
	hostBucket              bool                              // WithHostBucket
	hostBucketBases         []string                          // WithHostBucketBase
	autoBucket              bool                              // WithAutoBucket
	uploader                MultipartBackend
	log                     Logger
}

// New creates a new GoFakeS3 using the supplied Backend. Backends are pluggable.
// Several Backend implementations ship with GoFakeS3, which can be found in the
// gofakes3/backends package.
func New(backend Backend, options ...Option) *GoFakeS3 { _ = "STUB: not implemented"; return nil }

// versioned MUST be set before options as one of the options disables it:

func (g *GoFakeS3) nextRequestID() uint64 { _ = "STUB: not implemented"; return 0 }

// Create the AWS S3 API
func (g *GoFakeS3) Server() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func (g *GoFakeS3) timeSkewMiddleware(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// hostBucketMiddleware forces the server to use VirtualHost-style bucket URLs:
// https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html
func (g *GoFakeS3) hostBucketMiddleware(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// hostBucketBaseMiddleware forces the server to use VirtualHost-style bucket URLs:
// https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html
func (g *GoFakeS3) hostBucketBaseMiddleware(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (g *GoFakeS3) httpError(w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

// FIXME: request id

func (g *GoFakeS3) listBuckets(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// S3 has two versions of this API, both of which are close to identical. We manage that
// jank in here so the Backend doesn't have to with the following tricks:
//
// - Hiding the NextMarker inside the ContinuationToken for V2 calls
// - Masking the Owner in the response for V2 calls
//
// The wrapping response objects are slightly different too, but the list of
// objects is pretty much the same.
//
// - https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGET.html
// - https://docs.aws.amazon.com/AmazonS3/latest/API/v2-RESTBucketGET.html
func (g *GoFakeS3) listBucket(bucketName string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// We have observed (though not yet confirmed) that simple clients
// tend to work fine if you simply ignore pagination, so the
// default if this is not implemented is to retry without it. If
// you care about this performance impact for some weird reason,
// you'll need to handle it yourself.

// From the S3 docs: "This element is returned only if you specify
// a delimiter request parameter." Dunno why. This hack has been moved
// into GoFakeS3 to spare backend implementers the trouble.

// We are just cheating with these continuation tokens; they're just the NextMarker
// from v1 in disguise! That may change at any time and should not be relied upon
// though.

// On the topic of "fetch-owner", the AWS docs say, in typically vague style:
// "If you want the owner information in the response, you can specify
// this parameter with the value set to true."
//
// What does the bare word 'true' mean when we're talking about a query
// string parameter, which can only be a string? Does it mean the word
// 'true'? Does it mean 'any truthy string'? Does it mean only the key
// needs to be present (i.e. '?fetch-owner'), which we are assuming
// for now? This is why you need proper technical writers.
//
// Probably need to hit up the s3assumer at some point, but until then, here's
// another FIXME!

func (g *GoFakeS3) getBucketLocation(bucketName string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// S300006

func (g *GoFakeS3) listBucketVersions(bucketName string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// S300004:

// S300004: S3 ignores everything if you pass an empty key marker so
// let's hide that bit of ugliness from Backend.

// S300005: S3 returns the _string_ 'null' for the version ID if the
// bucket has never had versioning enabled. GoFakeS3 backend
// implementers should be able to simply return the empty string;
// GoFakeS3 itself should handle this particular bit of jank once and
// once only.

// CreateBucket creates a new S3 bucket in the BoltDB storage.
func (g *GoFakeS3) createBucket(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteBucket deletes the bucket in the underlying backend, if and only if it
// contains no items.
func (g *GoFakeS3) deleteBucket(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Support for Minio's DeleteBucket with force-delete header.

// HeadBucket checks whether a bucket exists.
func (g *GoFakeS3) headBucket(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// GetObject retrievs a bucket object.
func (g *GoFakeS3) getObject(
	bucket, object string,
	versionID VersionID,
	w http.ResponseWriter,
	r *http.Request,
) error {
	_ = "STUB: not implemented"
	return nil
}

// get object from backend

// Writes Content-Length, and Content-Range if applicable:

// writeGetOrHeadObjectResponse contains shared logic for constructing headers for
// a HEAD and a GET request for a /bucket/object URL.
func (g *GoFakeS3) writeGetOrHeadObjectResponse(obj *Object, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	// "If the current version of the object is a delete marker, Amazon S3
	// behaves as if the object was deleted and includes x-amz-delete-marker:
	// true in the response."
	return nil
}

// headObject retrieves only meta information of an object and not the whole.
func (g *GoFakeS3) headObject(
	bucket, object string,
	versionID VersionID,
	w http.ResponseWriter,
	r *http.Request,
) error {
	_ = "STUB: not implemented"
	return nil
}

// createObjectBrowserUpload allows objects to be created from a multipart upload initiated
// by a browser form.
func (g *GoFakeS3) createObjectBrowserUpload(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// maximum amount of memory before temp files are used

// FIXME: how does Content-MD5 get sent when using the browser? does it?

// CreateObject creates a new S3 object.
func (g *GoFakeS3) createObject(bucket, object string, w http.ResponseWriter, r *http.Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Parse conditional headers for PutObject

// XXX: no code for this, according to s3tests

// Satisfies s3tests

// XXX: no code for this, according to s3tests

// hashingReader is still needed to get the ETag even if integrityCheck
// is set to false:

// CopyObject copies an existing S3 object
func (g *GoFakeS3) copyObject(bucket, object string, meta map[string]string, w http.ResponseWriter, r *http.Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// XXX No support for versionId subresource

// XXX No support for delete marker
// "If the current version of the object is a delete marker, Amazon S3
// behaves as if the object was deleted."

// merge metadata, ACL is not preserved

func (g *GoFakeS3) deleteObject(bucket, object string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) deleteObjectVersion(bucket, object string, version VersionID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteMulti deletes multiple S3 objects from the bucket.
// https://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html
func (g *GoFakeS3) deleteMulti(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) initiateMultipartUpload(bucket, object string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// From the docs:
//
//	A part number uniquely identifies a part and also defines its position
//	within the object being created. If you upload a new part using the same
//	part number that was used with a previous part, the previously uploaded part
//	is overwritten. Each part must be at least 5 MB in size, except the last
//	part. There is no size limit on the last part of your multipart upload.
func (g *GoFakeS3) putMultipartUploadPart(bucket, object string, uploadID UploadID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Satisfies s3tests

func (g *GoFakeS3) abortMultipartUpload(bucket, object string, uploadID UploadID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) completeMultipartUpload(bucket, object string, uploadID UploadID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) listMultipartUploads(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) listMultipartUploadParts(bucket, object string, uploadID UploadID, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) getBucketVersioning(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// S300007

func (g *GoFakeS3) putBucketVersioning(bucket string, w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// S300007

// We only need to respond that this is not implemented if there's an
// attempt to enable it. If we receive a request to disable it, or an
// empty request, that matches the current state and has no effect so
// we can accept it.

func (g *GoFakeS3) ensureBucketExists(bucket string) error { _ = "STUB: not implemented"; return nil }

func (g *GoFakeS3) xmlEncoder(w http.ResponseWriter) *xml.Encoder {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoFakeS3) xmlDecodeBody(rdr io.ReadCloser, into interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func formatHeaderTime(t time.Time) string {
	_ = "STUB: not implemented"
	// .Format("Mon, 2 Jan 2006 15:04:05 MST")
	return ""
}

func metadataSize(meta map[string]string) int { _ = "STUB: not implemented"; return 0 }

func metadataHeaders(headers map[string][]string, at time.Time, sizeLimit int) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listBucketPageFromQuery(query url.Values) (page ListBucketPage, rerr error) {
	_ = "STUB: not implemented"
	return *new(ListBucketPage), nil
}

// List Objects V1 uses marker only:

// List Objects V2 uses continuation-token preferentially, or
// start-after if continuation-token is missing. continuation-token is
// an opaque value that looks like this: 1ueGcxLPRx1Tr/XYExHnhbYLgveDs2J/wm36Hy4vbOwM=.
// This just looks like base64 junk so we just cheat and base64 encode
// the next marker and hide it in a continuation-token.

// FIXME: log
// FIXME: confirm for sure what AWS does here

// List Objects V2 uses start-after if continuation-token is missing:

func listBucketVersionsPageFromQuery(query url.Values) (page ListBucketVersionsPage, rerr error) {
	_ = "STUB: not implemented"
	return *new(ListBucketVersionsPage), nil
}

// parsePutConditions extracts conditional headers from HTTP request headers
// and returns a PutConditions struct, or nil if no conditional headers are present.
func parsePutConditions(headers http.Header) (*PutConditions, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Check If-Match header
}

// Check If-None-Match header
