package libs

import "fmt"

/**
* 冒泡排序
 */
func BubbleSort(data []int) []int {
	l := len(data)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

/**
* 选择排序
 */
func SelectionSort(data []int) []int {
	l := len(data)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if data[min] > data[j] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
		fmt.Println(data)
	}
	return data
}

func InsertionSort(data []int) []int {
	for i := range data {
		preIndex := i - 1
		current := data[i]
		for preIndex >= 0 && current < data[preIndex] {
			data[preIndex+1] = data[preIndex]
			preIndex -= 1
		}
		data[preIndex+1] = current
	}
	return data
}
