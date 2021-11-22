package unzip

import (
	"fmt"
	"strings"

	"github.com/kawakatz/unzipR/pkg/utils"

	"github.com/mholt/archiver"
)

// Unzip decompress .zip.
func Unzip(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewZip()
	archive.Unarchive(path, out)
}
