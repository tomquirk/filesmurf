package filesmurf

import (
	"errors"
	"os"
	"strings"
)

// Move moves a file by creating dest directory path and renaming the file
func Move(srcPath string, dstPath string) error {
	pathTree := strings.Split(dstPath, "/")
	os.MkdirAll(strings.Join(pathTree[:len(pathTree)-1], "/"), os.ModePerm)
	err := os.Rename(srcPath, dstPath)

	if err != nil {
		return errors.New("Error moving file")
	}
	return nil
}
