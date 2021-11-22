package unzip

import (
	"fmt"
	"strings"
	"unzipR/pkg/utils"

	"github.com/mholt/archiver"
)

// Tar decompress .tar.
func Tar(path string) {
	out := utils.OutDir(path)
	fmt.Println(path + " -> " + strings.Replace(out, " ", "\\ ", -1) + "/*")

	archive := archiver.NewTar()
	archive.Unarchive(path, out)
}
