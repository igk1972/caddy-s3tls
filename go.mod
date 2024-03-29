module github.com/igk1972/caddy-s3tls

go 1.13

replace github.com/mholt/certmagic => github.com/mholt/certmagic v0.6.2

require (
	github.com/aws/aws-sdk-go v1.25.23
	github.com/caddyserver/caddy v1.0.3
	github.com/igk1972/magicstorage v0.0.2
	github.com/mholt/certmagic v0.8.3
	github.com/stretchr/testify v1.4.0
)
