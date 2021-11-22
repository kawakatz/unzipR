package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/djherbis/times"
	"github.com/fatih/color"
)

func Usage() {
	fmt.Println("usage: unzipR <path>")
}

func GrepSlice(s []string, key string) bool {
	for _, v := range s {
		if strings.Contains(strings.ToLower(v), strings.ToLower(key)) {
			return true
		}
	}

	return false
}

func GrepColor(str string, key string) string {
	index := strings.Index(strings.ToLower(str), strings.ToLower(key))
	colored := str[:index] + color.GreenString(str[index:index+len(key)]) + str[index+len(key):]

	return colored
}

func LsR(dir string) []string {
	files, _ := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, LsR(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func IsDir(path string) bool {
	fi, _ := os.Stat(path)
	return fi.IsDir()
}

func OutDir(path string) string {
	out := strings.TrimSuffix(path, filepath.Ext(path))
	if _, err := os.Stat(out); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(out, 0777)
		return out
	}

	t, _ := times.Stat(path)
	out += "-" + strconv.FormatInt(t.ChangeTime().UnixNano(), 10)[:4]

	if _, err := os.Stat(out); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(out, 0777)
		return out
	} else {
		i := 2
		for {
			tmp := out + " " + strconv.Itoa(i)
			if _, err := os.Stat(tmp); errors.Is(err, os.ErrNotExist) {
				_ = os.MkdirAll(tmp, 0777)
				return tmp
			}

			i++
		}
	}
}

func OutPath(path string) string {
	var out string
	if strings.HasSuffix(path, ".gz") {
		out = path[:len(path)-3]
	} else if strings.HasSuffix(path, ".bz2") {
		out = path[:len(path)-4]
	} else if strings.HasSuffix(path, ".lz4") {
		out = path[:len(path)-4]
	} else if strings.HasSuffix(path, ".xz") {
		out = path[:len(path)-3]
	} else {
		out = path + "-decompressed"
	}

	if _, err := os.Stat(out); errors.Is(err, os.ErrNotExist) {
		return out
	} else {
		i := 2
		for {
			tmp := out + " " + strconv.Itoa(i)
			if _, err := os.Stat(tmp); errors.Is(err, os.ErrNotExist) {
				return tmp
			}

			i++
		}
	}
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
