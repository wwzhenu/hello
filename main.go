package main

import (
	"fmt"
	"wwzhenu/sort/libs"
)

func main() {
	data := []int{9, 8, 7, 9, 7, 4, 5, 2, 1}
	// libs.BubbleSort(data)
	// libs.SelectionSort(data)
	fmt.Println(libs.InsertionSort(data))
}
