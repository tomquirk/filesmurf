package filesmurf

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// MatchFunc function defines the set of rules that determine which files will be actioned
type MatchFunc func(string) bool

// ActionFunc function defines the action to be applied to a matched file
type ActionFunc func(string)

// Run runs the things
func Run(rootPath string, match MatchFunc, action ActionFunc) {
	startTime := time.Now()
	var c int

	err := filepath.Walk(rootPath, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}

		if match(path) {
			c++
			action(path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("walk error [%v]\n", err)
	}

	elapsed := time.Since(startTime)
	log.Printf("modified %d files in %s", c, elapsed)
}
