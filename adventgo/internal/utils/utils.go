package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ROOT_DIR_NM = "advent-challenges"
)

func FindRootPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	currentDir := wd
	for {
		if filepath.Base(currentDir) == ROOT_DIR_NM {
			return filepath.Abs(currentDir)
		}
		currentDir = filepath.Dir(currentDir)
	}
}

func InitDisplay(day string, part int, inputPath string) {
	fmt.Println("")
	fmt.Printf("Solution: Day %s Problem %d \n", day, part)
	fmt.Println("==============================")
	fmt.Printf("Data File: %s", inputPath)
	fmt.Println("")
}
