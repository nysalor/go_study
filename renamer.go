package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		rename(os.Args[1], os.Args[2])
	} else if len(os.Args) > 1 {
		rename(os.Args[1])
	} else {
		rename(".")
	}

}

func rename(params ...string) {
	ext := "*"

	if len(params) > 1 {
		ext = params[1]
	}
	pattern := params[0] + "/*." + ext

	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		names := strings.Split(file, ".")
		renamed := fmt.Sprint(i) + "." + names[1]

		if err := os.Rename(file, renamed); err != nil {
			fmt.Printf("renamed %v -> %v \n", file, renamed)
		}
	}
}
