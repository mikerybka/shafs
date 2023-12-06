package shafs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
)

type Server struct {
	Dir string
}

func (s *Server) Save(b []byte) (string, error) {
	hash := util.SHA256(b)
	path := filepath.Join(s.Dir, hash)
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(path, b, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return hash, nil
}

func (s *Server) Load(hash string) ([]byte, error) {
	path := filepath.Join(s.Dir, hash)
	return os.ReadFile(path)
}
