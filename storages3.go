package storages3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/caddyserver/caddy/caddytls"
	"github.com/igk1972/magicstorage"
	"github.com/mholt/certmagic"
)

const (
	envNameEndpoint = "AWS_S3_ENDPOINT"
	envNameRegion   = "AWS_REGION"
	envNameBucket   = "AWS_BUCKET"
	envNamePath     = "AWS_PATH"
	envValueBucket  = "caddy"
	envValuePath    = "certmagic"
)

func init() {
	caddytls.RegisterClusterPlugin("s3", constructS3ClusterPlugin)
}

func constructS3ClusterPlugin() (certmagic.Storage, error) {
	return newS3Storage(), nil
}

func newS3Storage() *magicstorage.S3Storage {
	cfg := aws.NewConfig()
	if endpoint := os.Getenv(envNameEndpoint); endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
	}
	if region := os.Getenv(envNameRegion); region != "" {
		cfg.Region = aws.String(region)
	}
	var awsBucket = envValueBucket
	if bucket := os.Getenv(envNameBucket); bucket != "" {
		awsBucket = bucket
	}
	var awsPath = envValuePath
	if path := os.Getenv(envNamePath); path != "" {
		awsPath = path
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	store := &magicstorage.S3Storage{
		Bucket: aws.String(awsBucket),
		Svc:    svc,
		Path:   awsPath,
	}
	return store
}
