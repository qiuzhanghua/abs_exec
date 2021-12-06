package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	absPath, err := AbsPath()
	if err != nil {
		panic(err)
	}
	fmt.Println(absPath)
	os.Exit(0)
}

func AbsPath() (string, error) {
	// os.Executable() not work as expected
	//execPath, err := os.Executable()
	//if err != nil {
	//	return ".", err
	//}
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	absPath, err := filepath.EvalSymlinks(execPath)
	//	absPath, err := filepath.EvalSymlinks(os.Args[0])
	if err != nil {
		return ".", err
	}
	absPath, err = filepath.Abs(absPath)
	if err != nil {
		return ".", err
	}
	absPath, err = filepath.Abs(path.Dir(absPath))
	if err != nil {
		return ".", err
	}
	return absPath, nil
}
