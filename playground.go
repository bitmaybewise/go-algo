package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Println("Index ", i, "Value ", v)
	}
}
