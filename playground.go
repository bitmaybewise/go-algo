package main

import (
	"fmt"

	"github.com/hlmerscher/go-algo/algo"
)

func main() {
	col := []int{7, 5, 1, 8, 3}
	algo.SelectionSort(col)
	fmt.Println(col)
}
