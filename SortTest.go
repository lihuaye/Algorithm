package main

import (
	"Algorithm/sort"
	"fmt"
)

func main() {
	n := 100000000
	arr := mysort.GenerateRandomArray(n, 0, n)
	//arr := []int{7, 9, 1, 3, 5, 7, 89, 123, 453, 675, 765, 865, 324, 4645, 76573, 234, 34, 54353, 433, 12}
	//mysort.TestSort("SelectionSort", arr, mysort.SelectionSort)
	//mysort.TestSort("InsertionSort", arr, mysort.InsertionSort)
	//utils.WriteRandom2File("hello.txt", arr)
	//arr := utils.ReadRandom("hello.txt", 500000)
	//mysort.TestSort("BubbleSort", arr, mysort.BubbleSort)
	//mysort.TestSort("SelectionSort", arr, mysort.SelectionSort)
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	fmt.Println(arr)
	mysort.TestSort("ShellSort", arr, mysort.ShellSort)
	//mysort.TestSort("InsertionSort", arr1, mysort.InsertionSort)
	//mysort.TestSort("SelectionSort", arr1, mysort.SelectionSort)
	mysort.TestSort("MergeSort", arr, mysort.MergeSort)
	//fmt.Println("=====")
	//arr1 := mysort.GenerateBasicOrderArray(n, 100)
	//mysort.TestSort("InsertionSort", arr1, mysort.InsertionSort)
}
