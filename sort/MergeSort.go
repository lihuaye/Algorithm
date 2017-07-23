package mysort

// 归并排序
func MergeSort(arr []int) {
	size := len(arr)
	if size == 1 {
		return
	}
	MergeSort(arr[:size/2])
	MergeSort(arr[size/2:])

	copyArr := make([]int, size)
	copy(copyArr, arr)

	l := 0
	r := size / 2
	// 分割点[l,m),[m,size)
	m := size / 2
	for k := 0; k < size; k++ {
		if l < m && r < size {
			if copyArr[l] > copyArr[r] {
				arr[k] = copyArr[r]
				r++
			} else {
				arr[k] = copyArr[l]
				l++
			}
		} else if l == m {
			arr[k] = copyArr[r]
			r++
		} else if r == size {
			arr[k] = copyArr[l]
			l++
		}
	}
}
