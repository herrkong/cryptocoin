package main

import (
	"cryptocoin/src/public/sort/mergeSort"
	"cryptocoin/src/public/sort/selectSort"
	"fmt"
)

func main() {
	Input := []int{5, 4, 3, 6, 2, 1}
	OutPut := mergeSort.MergeSort(Input)
	selectSortOutPut := selectSort.SelectSort(Input)
	fmt.Println(OutPut)
	fmt.Println(selectSortOutPut)


}
