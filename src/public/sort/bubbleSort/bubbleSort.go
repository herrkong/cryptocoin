package bubbleSort

func BubleSort(array []int) []int {
	length := len(array)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			// j <= length-i-1 这个是关键，每次 i ，少比较最后一位数组
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
	return array
}
