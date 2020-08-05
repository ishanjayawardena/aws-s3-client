package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func createClientSession(env *awsEnvrionmentInfo) *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:      env.region,
		Credentials: credentials.NewSharedCredentials(env.credentialsFile, env.profile),
	}))
}
