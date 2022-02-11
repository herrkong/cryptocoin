package heapsort

//https://zhuanlan.zhihu.com/p/271662316
func HeapSort(Input []int) (OutPut []int) {
	first := 0
	lo := first
	hi := len(Input)

	// 大根堆
	// 节点大于其叶子结点的值
	// 由下往上
	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(Input, i, hi, first)
	}
	// 6 5 3 4 2 1

	// Pop elements, largest first, into end of data.
	// 把最大的丢到末尾 再对前面的元素调整出大根堆
	for i := hi - 1; i >= 0; i-- {
		Input[first], Input[first+i] = Swap(Input[first], Input[first+i])
		siftDown(Input, lo, i, first)
	}

	OutPut = Input
	return OutPut
}

// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown(data []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		// 左节点小于右节点
		if child+1 < hi && Less(data[first+child], data[first+child+1]) {
			child++
		}
		// root大于子节点 无须再调整了
		if !Less(data[first+root], data[first+child]) {
			return
		}
		// 交换值 保证根节点最大
		data[first+root], data[first+child] = Swap(data[first+root], data[first+child])
		root = child
	}
}

func Len(d []int) int {
	return len(d)
}

func Less(i, j int) bool {
	return i <= j
}

func Swap(i, j int) (int, int) {
	return j, i
}
