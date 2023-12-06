package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/mikerybka/shafs"
	"github.com/mikerybka/util"
)

func main() {
	in, out := os.Args[1], os.Args[2]
	err := mkSHAFS(in, out)
	if err != nil {
		panic(err)
	}
}

func mkSHAFS(in, out string) error {
	inFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer inFile.Close()
	fi, err := inFile.Stat()
	if err != nil {
		return err
	}
	if fi.IsDir() {
		err := os.MkdirAll(out, os.ModePerm)
		if err != nil {
			return err
		}
		entries, err := inFile.ReadDir(0)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			err = mkSHAFS(filepath.Join(in, entry.Name()), filepath.Join(out, entry.Name()))
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		b, err := io.ReadAll(inFile)
		if err != nil {
			return err
		}
		f, err := shafs.Save(b)
		if err != nil {
			return err
		}
		return util.WriteJSONFile(out, f)
	}
}
