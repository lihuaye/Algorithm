package mysort

// 希尔排序
func ShellSort(arr []int) {
	size := len(arr)
	// 动态定义增量大小
	gap := 1
	for gap < size/3 {
		gap = gap*3 + 1
	}

	for ; gap > 0; gap /= 3 {
		// i=gap 是因为在gap前都是每组的第一个元素
		// eg.
		// arr = [3,1,4,5,2,1]
		// 当gap=2时
		// 第一组 [3,4,2]
		// 第二组 [1,5,1]
		// 插排，第一个元素默认有序，使用从gap下标开始进行插排
		for i := gap; i < size; i++ {
			// 插排
			temp := arr[i]
			j := i
			for ; j >= gap; j -= gap {
				if temp < arr[j-gap] {
					arr[j] = arr[j-gap]
				} else {
					break
				}
			}
			arr[j] = temp
		}
	}
}
