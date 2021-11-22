package decompress

import (
	"io/ioutil"
	"strings"

	"unzipR/pkg/unzip"

	"github.com/h2non/filetype"
)

func Do(path string) {
	buf, _ := ioutil.ReadFile(path)
	kind, _ := filetype.Match(buf)
	mime := kind.MIME.Value

	switch mime {
	case "application/zip":
		unzip.Unzip(path)
	case "application/x-tar":
		unzip.Tar(path)
	case "application/x-7z-compressed":
		unzip.Sevenzip(path)
	case "application/vnd.rar":
		unzip.Unrar(path)
	case "application/gzip":
		if strings.HasSuffix(path, ".tar.gz") || strings.HasSuffix(path, ".tgz") {
			unzip.TarGzip(path)
		} else {
			unzip.Gunzip(path)
		}
	case "application/x-bzip2":
		if strings.HasSuffix(path, ".tar.bz2") {
			unzip.TarBz2(path)
		} else {
			unzip.Bzip2(path)
		}
	case "application/x-lzip":
		if strings.HasSuffix(path, ".tar.lz4") {
			unzip.TarLz4(path)
		} else {
			unzip.Lz4(path)
		}
	case "application/x-xz":
		if strings.HasSuffix(path, ".tar.xz") {
			unzip.TarXz(path)
		} else {
			unzip.Xz(path)
		}
	}
}
