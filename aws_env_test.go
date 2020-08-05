package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvironmentConfig(t *testing.T) {
	for scenario, expected := range map[string]struct {
		profile          string
		bucket           string
		region           string
		credentials_file string
	}{

		"eu-stg":  {stg_profile, "stg-bucket-europe", euRegion, ""},
		"us-stg":  {stg_profile, "stg-bucket-usa", usRegion, ""},
		"eu-prod": {prd_profile, "prd-bucket-europe", euRegion, ""},
		"us-prod": {prd_profile, "prd-bucket-usa", usRegion, ""},
	} {

		senv, err := getStaticEnvInfo(scenario, "")

		require.NoError(t, err)
		require.NotEqual(t, nil, senv)

		require.Equal(t, senv.profile, expected.profile)
		require.Equal(t, senv.bucket, expected.bucket)
		require.Equal(t, *senv.region, expected.region)
		require.Equal(t, senv.credentials_file, expected.credentials_file)
	}
}
