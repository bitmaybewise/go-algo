package main

import (
	"fmt"

	"github.com/hlmerscher/go-algo/helper"

	"github.com/hlmerscher/go-algo/algo"
)

func main() {
	col := helper.GenRandomCollection(100)
	algo.SelectionSort(col)
	fmt.Printf("Selection sort:\n%v\n", col)
}
