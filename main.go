package main

import "fmt"

func swap(nums []int, firstIdx, secondIdx int) {
	temp := nums[firstIdx]
	nums[firstIdx] = nums[secondIdx]
	nums[secondIdx] = temp
}

func Partition(nums []int, start, end int) int {
	pivot := nums[end]
	pIndex := start
	for i := start; i < end; i++ {
		if nums[i] <= pivot {
			swap(nums, i, pIndex)
			pIndex += 1
		}
	}

	swap(nums, pIndex, end)
	return pIndex
}

func Quicksort(nums []int, start, end int) {
	if start < end {
		pivot := Partition(nums, start, end)
		Quicksort(nums, start, pivot-1)
		Quicksort(nums, pivot+1, end)
	}

}

func main() {
	nums := []int{8, 10, 2, 100, 21}
	Quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)

}
