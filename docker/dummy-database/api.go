package main

import (
	"fmt"
	"os"
)

func main() {

	lines, err := os.Stat("api.go")

	fmt.Println(lines, err)
}
