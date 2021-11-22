package unzip

import (
	"fmt"
	"os"
	gopath "path"
	"strings"
	"unzipR/pkg/utils"

	"github.com/mholt/archiver"
)

func TarLz4(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewTarLz4()
	archive.Unarchive(path, out)
}

func TarSz(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewTarSz()
	archive.Unarchive(path, out)
}

func TarXz(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewTarXz()
	archive.Unarchive(path, out)
}

func Lz4(path string) {
	out := utils.OutPath(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1))

	if !utils.IsExist(gopath.Dir(out)) {
		_ = os.MkdirAll(gopath.Dir(out), 0777)
	}
	dist, _ := os.Create(out)
	defer dist.Close()

	src, _ := os.Open(path)
	defer src.Close()

	archive := archiver.NewLz4()
	archive.Decompress(src, dist)
}

func Xz(path string) {
	out := utils.OutPath(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1))

	if !utils.IsExist(gopath.Dir(out)) {
		_ = os.MkdirAll(gopath.Dir(out), 0777)
	}
	dist, _ := os.Create(out)
	defer dist.Close()

	src, _ := os.Open(path)
	defer src.Close()

	archive := archiver.NewXz()
	archive.Decompress(src, dist)
}
