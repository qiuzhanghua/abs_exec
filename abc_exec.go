package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	folderPath, err1 := os.Executable()
	if err1 != nil {
		os.Exit(-1)
	}
	absPath, err2 := filepath.EvalSymlinks(folderPath)
	if err2 != nil {
		os.Exit(-2)
	}
	dir := path.Dir(absPath)
	fmt.Print(dir)
	os.Exit(0)
}
