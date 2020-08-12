package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvironmentConfig(t *testing.T) {
	for scenario, expected := range map[string]struct {
		profile         string
		bucket          string
		region          string
		credentialsFile string
	}{

		"eu-stg":  {stgProfile, "stg-bucket-europe", euRegion, ""},
		"us-stg":  {stgProfile, "stg-bucket-usa", usRegion, ""},
		"eu-prod": {prdProfile, "prd-bucket-europe", euRegion, ""},
		"us-prod": {prdProfile, "prd-bucket-usa", usRegion, ""},
	} {

		senv, err := getStaticEnvInfo(scenario, "")

		require.NoError(t, err)
		require.NotEqual(t, nil, senv)

		require.Equal(t, senv.profile, expected.profile)
		require.Equal(t, senv.bucket, expected.bucket)
		require.Equal(t, *senv.region, expected.region)
		require.Equal(t, senv.credentialsFile, expected.credentialsFile)
	}
}
