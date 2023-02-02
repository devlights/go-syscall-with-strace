package main

import (
	"fmt"
	"os"

	"github.com/devlights/go-syscall-with-strace/pkg/enumerable"
)

func main() {
	file, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for v := range enumerable.Range(1, 10) {
		fmt.Fprintln(file, v)
	}
}
