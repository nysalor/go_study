package main

import (
    "fmt"
    "os"
    "path/filepath"
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

func rename(strings...string) {
	ext := "*"

	if len(strings) > 1 {
		ext = strings[1]
	}
	pattern := strings[0] + "/*." + ext

	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
