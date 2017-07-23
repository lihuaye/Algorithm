package mysort

// 归并排序
func MergeSort(arr []int) {
	size := len(arr)
	if size == 1 {
		return
	}
	//递归左半部分
	MergeSort(arr[:size/2])
	//递归右半部分
	MergeSort(arr[size/2:])
	//拷贝数组
	copyArr := make([]int, size)
	copy(copyArr, arr)
	//左起始点下标
	l := 0
	//右起始点下标
	r := size / 2
	// 分割点[l,m),[m,size)
	m := size / 2
	//合并排序好的左半部分和右半部分
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
