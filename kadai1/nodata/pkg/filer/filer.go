package filer

import (
	"fmt"
	"image/jpeg"
	"image/png"
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
		if !info.IsDir() {
			file, err := os.Open("./" + path)
			if err != nil {
				return err
			}
			defer file.Close()
			img, err := png.Decode(file)
			if err != nil {
				fmt.Printf("not image: %+v  \n", path)
			}
			fmt.Printf("file byte: %+v \n", img)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error:%s", err)
	}

}
