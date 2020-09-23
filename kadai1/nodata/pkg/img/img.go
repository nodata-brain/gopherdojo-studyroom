package imgconv

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Path string
}

func (f *Config) SetPath(path string) {
	f.Path = path
}

func (f *Config) Conv() {
	err := filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			img, err := jpeg.Decode(file)
			if err != nil {
				fmt.Printf("not image: %+v  \n", path)
				return nil
			}
			p := strings.Replace(path, ".jpg", ".png", -1)
			m, err := os.Create(p)
			err = png.Encode(m, img)
			if err != nil {
				return err
			}
			fmt.Printf("create: %s  \n", path)

		}

		return nil
	})
	if err != nil {
		fmt.Printf("error:%s", err)
	}
}
