package admin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	root, _         = os.Getwd()
	globalViewPaths []string
	globalAssetFSes []AssetFSInterface
)

func init() {
	if path := os.Getenv("WEB_ROOT"); path != "" {
		root = path
	}
}

// RegisterViewPath register view path for all assetfs
func RegisterViewPath(pth string) {
	globalViewPaths = append(globalViewPaths, pth)

	for _, assetFS := range globalAssetFSes {
		if assetFS.RegisterPath(filepath.Join(root, "vendor", pth)) != nil {
			for _, gopath := range strings.Split(os.Getenv("GOPATH"), ":") {
				if assetFS.RegisterPath(filepath.Join(gopath, "src", pth)) == nil {
					break
				}
			}
		}
	}
}

type AssetFSInterface interface {
	RegisterPath(path string) error
	Asset(name string) ([]byte, error)
	Glob(pattern string) (matches []string, err error)
	Compile() error
}

type AssetFileSystem struct {
	Paths []string
}

func (fs *AssetFileSystem) RegisterPath(pth string) error {
	if _, err := os.Stat(pth); !os.IsNotExist(err) {
		var existing bool
		for _, p := range fs.Paths {
			if p == pth {
				existing = true
				break
			}
		}
		if !existing {
			fs.Paths = append(fs.Paths, pth)
		}
		return nil
	}
	return errors.New("not found")
}

func (fs *AssetFileSystem) Asset(name string) ([]byte, error) {
	for _, pth := range fs.Paths {
		if _, err := os.Stat(filepath.Join(pth, name)); err == nil {
			return ioutil.ReadFile(filepath.Join(pth, name))
		}
	}
	return []byte{}, fmt.Errorf("%v not found", name)
}

func (fs *AssetFileSystem) Glob(pattern string) (matches []string, err error) {
	for _, pth := range fs.Paths {
		if results, err := filepath.Glob(filepath.Join(pth, pattern)); err == nil {
			for _, result := range results {
				matches = append(matches, strings.TrimPrefix(result, pth))
			}
		}
	}
	return
}

func (fs *AssetFileSystem) Compile() error {
	return nil
}
