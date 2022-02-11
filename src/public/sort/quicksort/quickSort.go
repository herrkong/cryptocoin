package quicksort

func sortArray(nums []int) []int {
	return sort(nums, 0, len(nums)-1)
}

func sort(nums []int, lo, hi int) []int {
	if lo == hi {
		return []int{nums[lo]}
	}
	mid := int(uint(lo+hi) >> 1)
	left := sort(nums, lo, mid)
	right := sort(nums, mid+1, hi)
	return merge(left, right)
}

//合并两个有序数组
func merge(left []int, right []int) []int {
	ans := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}
	if i != len(left) {
		ans = append(ans, left[i:]...)
	}
	if j != len(right) {
		ans = append(ans, right[j:]...)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
