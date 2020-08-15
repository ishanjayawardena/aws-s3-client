package main

import "log"

// Reads content of an S3 object and displays that content in the OS pager
func run() {
	log.SetPrefix("aws-s3-client: ")
	log.SetFlags(0)

	// Set up AWS operation/request
	// determine the AWS region, bucket, and client profile for the given application environment
	senv, err := getStaticEnvInfo(Environment, "")

	if err != nil {
		log.Fatal("failed to create static application config: ", err)
	}

	// create s3 request information from program arguments
	req, err := buildS3RequestInfo(senv)
	if err != nil {
		log.Fatal("failed to create S3 request info: ", err)
	}

	log.Printf("displaying content of file='%s/%s'\n", req.bucket, req.key)
	// -- End of setting up AWS operation/request

	// Set up S3 service client
	// Create a reusable AWS session that can be shared between multiple service clients
	s := createClientSession(senv)
	// Create an S3 service client from the session
	var s3 = createS3Client(s)
	// Make sure we cancel the S3 request after the timeout
	// -- End of setting up service client

	// Request S3 object/file and build the content
	defer req.cancelFn()

	content, err := readS3ObjectContent(s3, req)
	if err != nil {
		log.Fatal("failed to read content from S3: ", err)
	}
	// -- End of user request processing

	// Display content to the user in the default $PAGER of the operating system
	if err = displayInPager(content.String()); err != nil {
		log.Fatal("failed to display output: ", err)
	}
}
