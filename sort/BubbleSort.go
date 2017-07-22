package mysort

// 冒泡排序
func BubbleSort(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}