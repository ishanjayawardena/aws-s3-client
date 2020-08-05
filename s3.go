package main

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 operations.

type S3RequestInfo struct {
	bucket   string
	key      string
	timeout  time.Duration
	cancelFn func()
	ctx      context.Context
}

func createS3Client(s *session.Session) *s3.S3 {
	return s3.New(s)
}

func readS3ObjectContent(s3c *s3.S3, req *S3RequestInfo) (*strings.Builder, error) {
	obj, err := requestS3Object(s3c, req)
	if err != nil {
		return nil, fmt.Errorf("S3 Read error: %v", err)
	}

	return s3ContentAsString(obj)
}

func requestS3Object(s3Client *s3.S3, req *S3RequestInfo) (*s3.GetObjectOutput, error) {
	obj, err := s3Client.GetObjectWithContext(req.ctx, &s3.GetObjectInput{
		Bucket: aws.String(req.bucket),
		Key:    aws.String(req.key),
	})
	if err == nil {
		return obj, nil
	}
	aerr, ok := err.(awserr.Error)
	if ok && aerr.Code() == s3.ErrCodeNoSuchKey {
		return nil, err
	}

	return nil, fmt.Errorf("S3 error: %v", err)
}

func s3ContentAsString(res *s3.GetObjectOutput) (*strings.Builder, error) {
	buf := new(strings.Builder)
	n, err := io.Copy(buf, res.Body)

	if err != nil {
		return nil, fmt.Errorf("S3 response read error: %v", err)
	}
	defer res.Body.Close()

	if *res.ContentLength != n {
		return nil, fmt.Errorf("S3 response read error. Content length doesn't match the read data length")
	}

	return buf, nil
}
