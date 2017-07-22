package mysort

import (
	"math/rand"
	"time"
	"fmt"
)

// 生成随机数数组
func GenerateRandomArray(size int, rangeStart int, rangeEnd int) []int {
	arr := make([]int, size)
	// 创建随机种子
	seed := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(seed)
	for i := 0; i < size; i++ {
		// 生成随机数
		next := r.Int()
		// 让随机数范围在[rangeStart,rangeEnd]中
		next = next%(rangeEnd-rangeStart+1) + rangeStart
		arr[i] = next
	}
	return arr
}

// 生成基本有序的数组
func GenerateBasicOrderArray(size int, shuffle int) []int {
	arr := make([]int, size)

	seed := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(seed)

	for i := 0; i < size; i++ {
		arr[i] = i
	}

	for i := 0; i < shuffle; i++ {
		aIndex := r.Intn(size)
		bIndex := r.Intn(size)
		Swap(&arr[aIndex], &arr[bIndex])
	}
	return arr
}

// 交换2个数
func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

// 验证是否排序完成
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// 测试排序算法
func TestSort(sortName string, arr []int, sort func([]int)) {
	begin := time.Now().UnixNano()
	sort(arr)
	end := time.Now().UnixNano()
	isSuccess := IsSorted(arr)
	spendSecond := float64(end-begin) / 1000000

	fmt.Println("当前排序算法：" + sortName)
	fmt.Printf("耗时：%f ms\n", spendSecond)
	if isSuccess {
		fmt.Println("排序成功")
	} else {
		fmt.Println("排序失败")
		fmt.Println(arr)
	}
}
