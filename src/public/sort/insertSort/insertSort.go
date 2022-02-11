package insertSort

import "fmt"

func InsertSort(Input []int) (OutPut []int) {
	a := 0
	b := len(Input)
	for i := a + 1; i < b; i++ {
		for j := i; j > a && Less(Input[j], Input[j-1]); j-- {
			Input[j], Input[j-1] = Swap(Input[j], Input[j-1])
		}
	}
	fmt.Println(Input)
	OutPut = Input
	return OutPut
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
