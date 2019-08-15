package sort

// 选择排序
// N^2/2次比较, N次交换

// -
// 1.运行时间与输入无关
// 2.数据移动是最少的, 交换次数与数组大小为线性关系
func SelectionSortInt(param []int) {
	for i, _ := range param {
		min := i
		for j := i + 1; j < len(param) && less(param[min], param[j]) > 0; j++ {
			min = j
		}
		exch(param, i, min)
	}
}

// 插入排序
// 平均N^2/4次比较 N^2/4次交换
// 最坏N^2/2次比较 N^2/2次交换
// 最好N-1次比较   0次交换

//-
// 插入排序对于部分有序的数组很有效
//
func InsertionSort(param []int) {
	for i, _ := range param {
		for j := i; j > 0 && less(param[j], param[j-1]) < 0; j-- {
			exch(param, j, j-1)
		}
	}
}

func less(i, j int) int {
	return i - j
}

func exch(param []int, i, j int) {
	param[i], param[j] = param[j], param[i]
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
