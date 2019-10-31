package storages3

import (
	"fmt"
	"net"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/caddyserver/caddy/caddytls"
	"github.com/mholt/certmagic"
	"github.com/securityclippy/magicstorage"
)

const (
	EnvNameEndpoint = "AWS_ENDPOINT"
	EnvNameBucket   = "AWS_BUCKET"
	EnvNamePath     = "AWS_PATH"
	EnvValueBucket  = "caddy"
	EnvValuePath    = "certmagic"
)

func init() {
	caddytls.RegisterClusterPlugin("s3", constructS3ClusterPlugin)
}

func constructS3ClusterPlugin() (certmagic.Storage, error) {
	return magicstorage.NewS3Storage()
}

func NewS3Storage() *S3Storage {
	cfg := aws.NewConfig()
	if region := os.Getenv(EnvNameRegion); region != "" {
		cfg.Region = aws.String(region)
	}
	if region := os.Getenv(EnvNameRegion); region != "" {
		cfg.Endpoint = aws.String(endpoint)
	}
	var awsBucket = EnvValueBucket
	if bucket := os.Getenv(EnvNameBucket); busket != "" {
		awsBucket = aws.String(bucket)
	}
	var awsPath = EnvValuePath
	if path := os.Getenv(EnvNamePath); path != "" {
		awsPath = aws.String(path)
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	store := &S3Storage{
		bucket: awsBucket,
		svc:    svc,
		Path:   awsPath,
	}
	return store
}
