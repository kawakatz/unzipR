package unzip

import (
	"fmt"
	"os"
	gopath "path"
	"strings"
	"unzipR/pkg/utils"

	"github.com/mholt/archiver"
)

// Gunzip decompress .gz.
func Gunzip(path string) {
	out := utils.OutPath(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1))

	if !utils.IsExist(gopath.Dir(out)) {
		_ = os.MkdirAll(gopath.Dir(out), 0777)
	}
	dist, _ := os.Create(out)
	defer dist.Close()

	src, _ := os.Open(path)
	defer src.Close()

	archive := archiver.NewGz()
	archive.Decompress(src, dist)
}

// TarGzip decompress .tar.gz.
func TarGzip(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewTarGz()
	archive.Unarchive(path, out)
}
