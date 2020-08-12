package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/endpoints"
)

// Represents static config of AWS resources and corresponding application environment
type awsEnvrionmentInfo struct {
	// path to the shared credentials file
	credentialsFile string
	// Profile name that must be used in AWS SDK clients
	profile string
	// Name of the root bucket available for a given application environment
	bucket string
	// AWS region
	region *string
}

const (
	// Profile name used in the staging application environments
	stgProfile = "staging-profile-name"

	// Profile name used in production application environments
	prdProfile = "production-profile-name"
)

// AWS regions that are in use
var euRegion = endpoints.EuWest1RegionID
var usRegion = endpoints.UsEast1RegionID

// Static configuration of AWS bucket, client profile, and associated application environment
var envInfo = map[string]struct {
	profile string
	bucket  string
	region  string
}{
	"eu-stg":  {stgProfile, "stg-bucket-europe", euRegion},
	"us-stg":  {stgProfile, "stg-bucket-usa", usRegion},
	"eu-prod": {prdProfile, "prd-bucket-europe", euRegion},
	"us-prod": {prdProfile, "prd-bucket-usa", usRegion},
}

// Returns AWS environment info (ie, sdk client profile name, aws region name, S3 bucket name)
func getStaticEnvInfo(env string, cred string) (*awsEnvrionmentInfo, error) {
	var res awsEnvrionmentInfo
	if senv, ok := envInfo[env]; ok {
		res.credentialsFile = cred
		res.profile = senv.profile
		res.bucket = senv.bucket
		res.region = &senv.region
		return &res, nil
	}
	return nil, fmt.Errorf("Unsupport application environment: '%s'", env)
}
