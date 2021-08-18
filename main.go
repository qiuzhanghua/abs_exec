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
	fmt.Println(os.Args[0])
	p1, _ := filepath.EvalSymlinks(os.Args[0])
	fmt.Println(p1)
	p2, _ := filepath.Abs(p1)
	fmt.Println(p2)
	p3 := filepath.Dir(p2)
	fmt.Println(p3)
	os.Exit(0)
}

func AbsPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return ".", err
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
