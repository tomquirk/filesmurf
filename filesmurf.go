package filesmurf

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GetDstPath returns the appropriate destination directory of a given source filepath
type GetDstPath func(string) string

// Run runs the things
func Run(rootPath string, getDstPath GetDstPath) {
	startTime := time.Now()
	var c int

	err := filepath.Walk(rootPath, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}
		dstPath := getDstPath(path)

		if dstPath != "" {
			c++
			pathTree := strings.Split(dstPath, "/")
			os.MkdirAll(strings.Join(pathTree[:len(pathTree)-1], "/"), os.ModePerm)
			os.Rename(path, getDstPath(path))
		}

		return nil
	})

	if err != nil {
		log.Fatalf("walk error [%v]\n", err)
	}

	elapsed := time.Since(startTime)
	log.Printf("moved %d files in %s", c, elapsed)
}
