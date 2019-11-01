package storages3

import (
	"github.com/caddyserver/caddy/caddytls"
	"github.com/igk1972/magicstorage"
	"github.com/mholt/certmagic"
)

func init() {
	caddytls.RegisterClusterPlugin("s3", constructS3ClusterPlugin)
}

func constructS3ClusterPlugin() (certmagic.Storage, error) {
	return magicstorage.NewS3Storage()
}
