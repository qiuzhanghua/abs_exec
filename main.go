package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//folderPath, err1 := os.Executable()
	//if err1 != nil {
	//	os.Exit(-1)
	//}
	//absPath, err2 := filepath.EvalSymlinks(folderPath)
	//if err2 != nil {
	//	os.Exit(-2)
	//}
	//dir := path.Dir(absPath)
	//fmt.Print(dir)
	absPath3, err3 := filepath.EvalSymlinks(os.Args[0])
	if err3 != nil {
		os.Exit(-1)
	}
	absPath4, err4 := filepath.Abs(absPath3)
	if err4 != nil {
		os.Exit(-2)
	}
	fmt.Print(path.Dir(absPath4))
	os.Exit(0)
}
