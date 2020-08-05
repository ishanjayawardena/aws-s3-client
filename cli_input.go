package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

// Parses command line arguments and builds S3 request information
func buildS3RequestInfo(senv *awsEnvrionmentInfo) (*S3RequestInfo, error) {
	var req S3RequestInfo

	req.bucket = string(senv.bucket)

	flag.StringVar(&req.key, "k", "", "Object key name.")
	flag.DurationVar(&req.timeout, "d", 10*time.Second, "S3 read timeout.")
	flag.Parse()

	if req.key == "" {
		return nil, fmt.Errorf("S3 key not specified")
	}

	req.ctx = context.Background()
	if req.timeout > 0 {
		req.ctx, req.cancelFn = context.WithTimeout(req.ctx, req.timeout)
	}

	return &req, nil
}
