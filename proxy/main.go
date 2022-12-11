package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("some_file.txt")
	if err != nil {
		panic(err)
	}

	accessorProxy := SomeObjectAccessProxy{}

	_, err = accessorProxy.access("some_file.txt")
	if err != nil {
		fmt.Println(err)
	}
}
