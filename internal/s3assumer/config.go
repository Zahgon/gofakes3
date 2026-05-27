package main

type Config struct {
	S3Endpoint         string
	S3Region           string
	S3PathStyle        bool
	S3TestBucketPrefix string
	Verbose            bool
}

func (c Config) BucketStandard() string { _ = "STUB: not implemented"; return "" }

func (c Config) BucketUnversioned() string { _ = "STUB: not implemented"; return "" }
