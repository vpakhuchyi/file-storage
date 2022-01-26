package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"file-storage/config"
)

// Client is a struct that shall be used for interactions with Storage.
type Client struct {
	Config config.Storage
}

// New is a constructor of Storage Client.
func New(cfg config.Storage) Client {
	return Client{
		Config: cfg,
	}
}

// Save saves given content as a file with a given name.
func (c Client) Save(fileName string, content []byte) error {
	return os.WriteFile(c.Config.FileDir+fileName, content, 0666)
}

// Download writes a file to dest based on its fileName.
func (c Client) Download(fileName string, dest io.Writer) error {
	f, err := os.OpenFile(c.Config.FileDir+fileName, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(dest, f); err != nil {
		return err
	}

	return nil
}

// List returns a list of met-information about existing files.
func (c Client) List() ([]string, error) {
	files := map[string]int64{}
	err := filepath.WalkDir(c.Config.FileDir, func(path string, d os.DirEntry, err error) error {
		fi, err := os.Lstat(path)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		files[fi.Name()] = fi.Size()

		return nil
	})
	if err != nil {
		return nil, err
	}

	meta := make([]string, len(files))
	for name, size := range files {
		meta = append(meta, fmt.Sprintf("file name: [%s] | file size: [%d bytes]\n", name, size))
	}

	meta[len(meta)-1] = truncateLastChar(meta[len(meta)-1]) // This is needed for redundant new line deletion.

	return meta, nil
}

func truncateLastChar(s string) string {
	return s[:len(s)-1]
}
