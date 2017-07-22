package mysort

// 插入排序
func InsertionSort(arr []int) {
	size := len(arr)
	for i := 1; i < size; i++ {
		temp := arr[i]
		j := i
		for ; j > 0; j-- {
			if arr[j-1] > temp {
				Swap(&arr[j], &arr[j-1])
			} else {
				break
			}
		}
		arr[j] = temp
	}
}
