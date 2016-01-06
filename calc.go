package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	parse(arg[1])
}

func parse(formula string) {
	fmt.Printf("%v\n", formula)
}
