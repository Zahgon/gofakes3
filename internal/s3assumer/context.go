package main

import (
	"context"
	"math/rand"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Context struct {
	context.Context
	config Config
	rand   *rand.Rand
}

func (c *Context) Config() Config           { _ = "STUB: not implemented"; return *new(Config) }
func (c *Context) Rand() *rand.Rand         { _ = "STUB: not implemented"; return nil }
func (c *Context) RandString(sz int) string { _ = "STUB: not implemented"; return "" }

func (c *Context) RandBytes(sz int) []byte { _ = "STUB: not implemented"; return nil }

func (c *Context) S3Client() *s3.Client { _ = "STUB: not implemented"; return nil }

func (c *Context) EnsureVersioningEnabled(client *s3.Client, bucket string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) EnsureVersioningNeverEnabled(client *s3.Client, bucket string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) getBucketVersioning(client *s3.Client, bucket string) (*s3.GetBucketVersioningOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Logger struct{}

func (l Logger) Log(vs ...interface{}) { _ = "STUB: not implemented"; return }
