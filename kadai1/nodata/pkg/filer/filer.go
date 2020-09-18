package filer

import (
	"fmt"
	"os"
	"path/filepath"
)

type Filer struct {
	Path string
}

func (f *Filer) SetPath(path string) {
	f.Path = path
}

func (f *Filer) Filer() {
	err := filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			fmt.Printf("dir name: %+v \n", info.Name())
		} else {
			fmt.Printf("file name: %+v \n", info.Name())
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error:%s", err)
	}

}
