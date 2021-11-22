package unzip

import (
	"fmt"
	"os"
	gopath "path"
	"strings"
	"unzipR/pkg/utils"

	"github.com/mholt/archiver"
)

func Bzip2(path string) {
	out := utils.OutPath(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1))

	if !utils.IsExist(gopath.Dir(out)) {
		_ = os.MkdirAll(gopath.Dir(out), 0777)
	}
	dist, _ := os.Create(out)
	defer dist.Close()

	src, _ := os.Open(path)
	defer src.Close()

	archive := archiver.NewBz2()
	archive.Decompress(src, dist)
}

func TarBz2(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewTarBz2()
	archive.Unarchive(path, out)
}
