package selectSort

func SelectSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		index := i // 假设元素最小
		for j := i + 1; j < n-1; j++ {
			if array[index] > array[j+1] {
				index = j + 1 // 记录元素最小的索引
			}
		}
		// 将最小的元素排序到左边
		array[i], array[index] = array[index], array[i]
		//i++,已经将最小的元素排序到最左边，然后在找第二大的元素
	}
	return array
}
