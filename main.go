package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ch := make(chan int)
	go func() {
		defer close(ch)

		for _, v := range []int{1, 2, 3, 4, 5} {
			ch <- v
		}
	}()

	for v := range ch {
		fmt.Fprintln(file, v)
	}
}
