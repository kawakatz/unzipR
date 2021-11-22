# unzip Recursively🤐
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/kawakatz/unzipR/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/kawakatz/unzipR"><img src="https://goreportcard.com/badge/github.com/kawakatz/unzipR"></a>
<a href="https://www.codefactor.io/repository/github/kawakatz/unzipR/badge"><img src="https://www.codefactor.io/repository/github/kawakatz/unzipR/badge"></a>
<a href="https://twitter.com/kawakatz"><img src="https://img.shields.io/twitter/follow/kawakatz.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a>  •
  <a href="#-todo">Todo</a>  •
  <a href="#acknowledgement">Acknowledgement</a>
</p>

unzipR decompress all the compressed files in the directory.

unzipR can decompress files in the following compressed formats:
- zip
- tar
- 7zip
- rar
- gzip (including .tar.gz and .tgz)
- bzip2 (including .tar.bz2)
- lzip4 (including .tar.lz4)
- xz (including .tar.xz)

## Installation
```sh
# install unzipR
➜  ~ go install -v github.com/kawakatz/unzipR/cmd/unzipR@latest
```

## Usage
```sh
➜  ~ unzipR <path>
```

`<path>` can be a path of a file or directory.<br>
If a directory path is specified, unzipR will **recursively** decompress all the files in the directory.<br>

## Todo
- Support for password-protected zip
- Decompressing a file after it has been unzipped if it contains a compressed file.

## Acknowledgement
This README.md format is inspired by  [@projectdiscovery](https://github.com/projectdiscovery/)🙇‍♂️