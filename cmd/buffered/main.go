package main

import (
	"bufio"
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

	bufW := bufio.NewWriter(file)
	for v := range enumerable.Range(1, 10) {
		fmt.Fprintln(bufW, v)
	}
	defer bufW.Flush()
}
