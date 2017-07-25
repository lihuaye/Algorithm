package main

import (
	"Algorithm/sort"
)

func main() {
	n := 100000
	//arr := mysort.GenerateRandomArray(n, 0, n)
	arr := mysort.GenerateBasicOrderArray(n, 0)
	//arr := []int{7, 9, 1, 3, 5, 7, 89, 123, 453, 675, 765, 865, 324, 4645, 76573, 234, 34, 54353, 433, 12}
	//mysort.TestSort("SelectionSort", arr, mysort.SelectionSort)
	//mysort.TestSort("InsertionSort", arr, mysort.InsertionSort)
	//utils.WriteRandom2File("hello.txt", arr)
	//arr := utils.ReadRandom("hello.txt", 500000)
	//fmt.Println(arr)
	//mysort.TestSort("BubbleSort", arr, mysort.BubbleSort)
	//fmt.Println(arr)
	//mysort.TestSort("BubbleSort", arr, mysort.BubbleSort)
	//mysort.TestSort("SelectionSort", arr, mysort.SelectionSort)
	//arr1 := make([]int, len(arr))
	//copy(arr1, arr)
	//mysort.TestSort("ShellSort", arr, mysort.ShellSort)
	//mysort.TestSort("InsertionSort", arr1, mysort.InsertionSort)
	//mysort.TestSort("MergeSort", arr1, mysort.MergeSort)
	mysort.TestSort("QuickSort", arr, mysort.QuickSort)
	//fmt.Println("=====")
	//arr1 := mysort.GenerateBasicOrderArray(n, 100)
	//mysort.TestSort("InsertionSort", arr1, mysort.InsertionSort)
}
