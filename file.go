package shafs

import (
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"

	"github.com/mikerybka/util"
)

func Save(b []byte) (*File, error) {
	h := sha256.New()
	size, err := h.Write(b)
	if err != nil {
		return nil, err
	}
	hash := hex.EncodeToString(h.Sum(nil))
	path := filepath.Join(util.HomeDir(), "data/shafs/cache", hash)
	err = util.WriteFile(path, b)
	if err != nil {
		return nil, err
	}
	return &File{
		Size:   size,
		SHA256: hash,
	}, nil
}

type File struct {
	Size   int
	SHA256 string
}
