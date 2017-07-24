package mysort

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
	//基点下标
	l := 0
	//小于基点值的最后一个下标
	minIndex := 0
	//大于基点值的最后一个下标
	maxIndex := 0
	//基点值
	point := arr[l]
	//重新排序,确定基点位置
	for i := 0; i < size; i++ {
		//arr[i]小于基点值时
		if arr[i] < point {
			//小于基准点的下标+1,用来容纳新值
			minIndex ++
			//交换值
			if maxIndex != 0 {
				Swap(&arr[minIndex], &arr[i])
				maxIndex++
			}
		} else {
			maxIndex++
		}
	}
	Swap(&arr[l], &arr[minIndex])
	return minIndex
}
