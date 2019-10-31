// +build s3

package storages3

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/mholt/certmagic"
	"github.com/securityclippy/magicstorage"
	"github.com/stretchr/testify/assert"
)

// const TestPrefix = "s3tlstest"

// these tests needs a running S3 server
func setupS3Env(t *testing.T) *S3Storage {

	// os.Setenv(EnvNamePrefix, TestPrefix)
	// os.Setenv(HTTPTokenEnvName, "2f9e03f8-714b-5e4d-65ea-c983d6b172c4")
	//
	cs, err := NewS3Storage()
	assert.NoError(t, err)

	// _, err = cs.S3Client.DeleteTree(TestPrefix, nil)
	// assert.NoError(t, err)
	return cs
}

func TestS3Storage_Store(t *testing.T) {
	cs := setupS3Env(t)

	err := cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.crt"), []byte("crt data"))
	assert.NoError(t, err)
}

func TestS3Storage_Exists(t *testing.T) {
	cs := setupS3Env(t)

	key := path.Join("acme", "example.com", "sites", "example.com", "example.com.crt")

	err := cs.Store(key, []byte("crt data"))
	assert.NoError(t, err)

	exists := cs.Exists(key)
	assert.True(t, exists)
}

func TestS3Storage_Load(t *testing.T) {
	cs := setupS3Env(t)

	key := path.Join("acme", "example.com", "sites", "example.com", "example.com.crt")
	content := []byte("crt data")

	err := cs.Store(key, content)
	assert.NoError(t, err)

	contentLoded, err := cs.Load(key)
	assert.NoError(t, err)

	assert.Equal(t, content, contentLoded)
}

func TestS3Storage_Delete(t *testing.T) {
	cs := setupS3Env(t)

	key := path.Join("acme", "example.com", "sites", "example.com", "example.com.crt")
	content := []byte("crt data")

	err := cs.Store(key, content)
	assert.NoError(t, err)

	err = cs.Delete(key)
	assert.NoError(t, err)

	exists := cs.Exists(key)
	assert.False(t, exists)

	contentLoaded, err := cs.Load(key)
	assert.Nil(t, contentLoaded)

	_, ok := err.(certmagic.ErrNotExist)
	assert.True(t, ok)
}

func TestS3Storage_Stat(t *testing.T) {
	cs := setupS3Env(t)

	key := path.Join("acme", "example.com", "sites", "example.com", "example.com.crt")
	content := []byte("crt data")

	err := cs.Store(key, content)
	assert.NoError(t, err)

	info, err := cs.Stat(key)
	assert.NoError(t, err)

	assert.Equal(t, key, info.Key)
}

func TestS3Storage_List(t *testing.T) {
	cs := setupS3Env(t)

	err := cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.crt"), []byte("crt"))
	assert.NoError(t, err)
	err = cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.key"), []byte("key"))
	assert.NoError(t, err)
	err = cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.json"), []byte("meta"))
	assert.NoError(t, err)

	keys, err := cs.List(path.Join("acme", "example.com", "sites", "example.com"), true)
	assert.NoError(t, err)
	assert.Len(t, keys, 3)
	assert.Contains(t, keys, path.Join("acme", "example.com", "sites", "example.com", "example.com.crt"))
}

func TestS3Storage_ListNonRecursive(t *testing.T) {
	cs := setupS3Env(t)

	err := cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.crt"), []byte("crt"))
	assert.NoError(t, err)
	err = cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.key"), []byte("key"))
	assert.NoError(t, err)
	err = cs.Store(path.Join("acme", "example.com", "sites", "example.com", "example.com.json"), []byte("meta"))
	assert.NoError(t, err)

	keys, err := cs.List(path.Join("acme", "example.com", "sites"), false)
	assert.NoError(t, err)

	assert.Len(t, keys, 1)
	assert.Contains(t, keys, path.Join("acme", "example.com", "sites", "example.com"))
}

func TestS3Storage_LockUnlock(t *testing.T) {
	cs := setupS3Env(t)
	lockKey := path.Join("acme", "example.com", "sites", "example.com", "lock")

	err := cs.Lock(lockKey)
	assert.NoError(t, err)

	err = cs.Unlock(lockKey)
	assert.NoError(t, err)
}

func TestS3Storage_TwoLocks(t *testing.T) {
	cs := setupS3Env(t)
	cs2 := setupS3Env(t)
	lockKey := path.Join("acme", "example.com", "sites", "example.com", "lock")

	err := cs.Lock(lockKey)
	assert.NoError(t, err)

	go time.AfterFunc(5*time.Second, func() {
		err = cs.Unlock(lockKey)
		assert.NoError(t, err)
	})

	err = cs2.Lock(lockKey)
	assert.NoError(t, err)

	err = cs2.Unlock(lockKey)
	assert.NoError(t, err)
}
