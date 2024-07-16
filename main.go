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

func Merge(arr, left, right []int) {
	k := 0
	posL := 0
	posR := 0
	lastIndexLeft := len(left) - 1
	lastIndexRight := len(right) - 1
	for {
		if posL > lastIndexLeft || posR > lastIndexRight {
			break
		}

		if left[posL] < right[posR] {
			arr[k] = left[posL]
			posL += 1
		} else if right[posR] <= left[posL] {
			arr[k] = right[posR]
			posR += 1
		}
		k += 1
	}

	for {
		if posL > lastIndexLeft {
			break
		}
		arr[k] = left[posL]
		k += 1
		posL += 1
	}

	for {
		if posR > lastIndexRight {
			break
		}
		arr[k] = right[posR]
		k += 1
		posR += 1
	}

}

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2

	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	for i := 0; i < mid; i++ {
		left[i] = arr[i]
	}
	for i := mid; i < len(arr); i++ {
		right[i-mid] = arr[i]
	}

	MergeSort(left)
	MergeSort(right)
	Merge(arr, left, right)
}

func main() {
	nums := []int{8, 10, 2, 100, 21}
	Quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)

}
