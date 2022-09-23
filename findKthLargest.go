package main

import (
	"math/rand"
	"time"
)

func partition(nums []int, low, high, pivotIdx int) ([]int, int) {
	pivotVal := nums[pivotIdx]
	nums[pivotIdx], nums[high] = nums[high], nums[pivotIdx]

	i := low
	for j := low; j < high; j++ {
		if nums[j] < pivotVal {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[high], nums[i] = nums[i], nums[high]

	return nums, i
}

func quickSelect(nums []int, low, high, currentK int) int {
	if low == high {
		return nums[low]
	}

	pivotIdx := rand.Intn(high-low) + low
	nums, pivotIdx = partition(nums, low, high, pivotIdx)

	if currentK == pivotIdx {
		return nums[currentK]
	} else if currentK < pivotIdx {
		return quickSelect(nums, low, pivotIdx-1, currentK)
	}

	return quickSelect(nums, pivotIdx+1, high, currentK)
}

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())

	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

/*
import (
	"math/rand"
	"sort"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/

func addNewValue_v1(kNums []int, new, k int) []int {
	nk := len(kNums)
	insert := sort.Search(nk, func(i int) bool {
		return kNums[i] <= new
	})
	if nk < k {
		kNums = append(kNums, 0)
	} else {
		nk--
	}

	copy(kNums[insert+1:nk+1], kNums[insert:nk])
	kNums[insert] = new

	return kNums
}

func findKthLargest_v1(nums []int, k int) int {
	kNums := []int{}
	for i := range nums {
		num := nums[i]

		if len(kNums) < k || num >= kNums[len(kNums)-1] {
			kNums = addNewValue_v1(kNums, num, k)
		}
	}

	return kNums[len(kNums)-1]
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func addNewValue_make(kNums []int, nk, new, k int) ([]int, int) {
	insert := sort.Search(nk, func(i int) bool {
		return kNums[i] <= new
	})
	if nk == k {
		nk--
	}

	copy(kNums[insert+1:nk+1], kNums[insert:nk])
	kNums[insert] = new

	return kNums, nk + 1
}

func findKthLargest_v2(nums []int, k int) int {
	nk := 0
	kNums := make([]int, k)
	kNums[len(kNums)-1] = MinInt
	for i := range nums {
		num := nums[i]

		if num >= kNums[len(kNums)-1] {
			kNums, nk = addNewValue_make(kNums, nk, num, k)
		}
	}

	return kNums[len(kNums)-1]
}

func partition(nums []int, low, high, pivotIdx int) ([]int, int) {
	pivotVal := nums[pivotIdx]
	nums[pivotIdx], nums[high] = nums[high], nums[pivotIdx]

	i := low
	for j := low; j < high; j++ {
		if nums[j] < pivotVal {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[high], nums[i] = nums[i], nums[high]

	return nums, i
}

func quickSelect(nums []int, low, high, currentK int) int {
	if low == high {
		return nums[low]
	}

	pivotIdx := rand.Intn(high-low) + low
	nums, pivotIdx = partition(nums, low, high, pivotIdx)

	if currentK == pivotIdx {
		return nums[currentK]
	} else if currentK < pivotIdx {
		return quickSelect(nums, low, pivotIdx-1, currentK)
	}

	return quickSelect(nums, pivotIdx+1, high, currentK)
}

func findKthLargest_works(nums []int, k int) int {
	low := 0
	high := len(nums) - 1

	currentK := len(nums) - k

	return quickSelect(nums, low, high, currentK)
}

func findKthLargest(nums []int, k int) int {
	low := 0
	high := len(nums) - 1
	currentK := len(nums) - k

	pivotIdx := -1
	for currentK != pivotIdx {
		if low == high {
			return nums[low]
		}

		pivotIdx = rand.Intn(high-low) + low
		nums, pivotIdx = partition(nums, low, high, pivotIdx)
		if currentK < pivotIdx {
			high = pivotIdx - 1
		} else {
			low = pivotIdx + 1
		}
	}

	return nums[k-1]
}
*/
