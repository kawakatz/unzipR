package main

import (
	"flag"
	"os"
	"sync"
	"unzipR/pkg/decompress"
	"unzipR/pkg/utils"
)

func main() {
	flag.Parse()
	cmdArgs := flag.Args()
	if len(cmdArgs) == 0 {
		utils.Usage()
		os.Exit(0)
	}

	target := cmdArgs[0]
	if !utils.IsExist(target) {
		utils.Usage()
		os.Exit(0)
	}

	if utils.IsDir(target) {
		files := utils.LsR(target)

		var wg sync.WaitGroup
		pathChan := make(chan string)
		for i := 0; i < 20; i++ {
			wg.Add(1)

			go func() {
				defer recover()

				for path := range pathChan {
					decompress.Do(path)
				}
				wg.Done()
			}()
		}

		for _, each := range files {
			pathChan <- each
		}
		close(pathChan)

		wg.Wait()
	} else {
		decompress.Do(target)
	}
}
