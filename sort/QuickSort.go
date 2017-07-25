package mysort

import "math/rand"

//快速排序
func QuickSort(arr []int) {
	size := len(arr)
	if size <= 1 {
		return
	}
	p := partition(arr, size)
	QuickSort(arr[0:p])
	QuickSort(arr[p+1:])
}

//确定基点的位置
func partition(arr []int, size int) int {
	//随机选取一个下标来做基点值
	randIndex := rand.Intn(size)
	//跟第一个元素进行交换
	Swap(&arr[0], &arr[randIndex])
	//小于基点值的最后一个下标
	minIndex := 0
	//基点值
	point := arr[0]
	//重新排序,确定基点位置
	for i := 1; i < size; i++ {
		//arr[i]小于基点值时
		if arr[i] < point {
			//小于基准点的下标+1,用来容纳新值
			minIndex ++
			Swap(&arr[minIndex], &arr[i])
		}
	}
	//将最后一位小于基点值的值和基点值进行交换构成 arr[0:minIndex] < 基点值 <arr[minIndex+1:]的数组
	Swap(&arr[0], &arr[minIndex])
	//返回基点应该在排序完成的数组中的下标
	return minIndex
}
