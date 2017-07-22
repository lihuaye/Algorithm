package mysort

// 选择排序
func SelectionSort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if i != minIndex {
			Swap(&arr[i], &arr[minIndex])
		}
	}
}
