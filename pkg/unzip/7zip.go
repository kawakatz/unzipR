package unzip

import (
	"fmt"
	"io"
	"os"
	gopath "path"
	"path/filepath"
	"strings"

	"github.com/kawakatz/unzipR/pkg/utils"

	"github.com/saracen/go7z"
)

// Sevenzip decompress .7z.
func Sevenzip(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	sz, _ := go7z.OpenReader(path)
	defer sz.Close()

	for {
		hdr, err := sz.Next()
		if err == io.EOF {
			break
		}

		if hdr.IsEmptyStream && !hdr.IsEmptyFile {
			if err := os.MkdirAll(hdr.Name, os.ModePerm); err != nil {
				panic(err)
			}
			continue
		}

		outPath := filepath.Join(out, hdr.Name)
		if !utils.IsExist(gopath.Dir(outPath)) {
			_ = os.MkdirAll(gopath.Dir(outPath), 0777)
		}
		f, err := os.Create(outPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := io.Copy(f, sz); err != nil {
			panic(err)
		}
	}
}
