package main

import "container/heap"

// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	//return kSmallestPairs_Custom(nums1, nums2, k)
	return kSmallestPairs_builtin(nums1, nums2, k)
}

// Code Below uses "containers/heap"
type minHeap [][]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i][0]+h[i][1] < h[j][0]+h[j][1]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(value interface{}) {
	*h = append(*h, value.([]int))
}

func (h *minHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func kSmallestPairs_builtin(nums1 []int, nums2 []int, k int) [][]int {
	data := &minHeap{}

	// 3rd value is index from nums2
	for i := range nums1 {
		heap.Push(data, []int{nums1[i], nums2[0], 0})
	}
	ret := make([][]int, 0, k)
	for ; k > 0 && data.Len() > 0; k-- {
		min := heap.Pop(data).([]int)
		ret = append(ret, []int{min[0], min[1]})
		// if index+1 exists, add to heap
		if min[2]+1 < len(nums2) {
			heap.Push(data, []int{min[0], nums2[min[2]+1], min[2] + 1})
		}
	}
	return ret
}

// Code Above uses "containers/heap"

// Code Below Defines a Heap
func kSmallestPairs_Custom(nums1 []int, nums2 []int, k int) [][]int {
	heap := NewHeap()
	for i, num := range nums1 {
		if i >= k {
			break
		}
		heap.Push(Data{i, 0, num + nums2[0]})
	}
	ans := make([][]int, 0, k)
	for j := k; j > 0 && len(heap) > 0; j-- {
		data := heap.Pop()
		ans = append(ans, []int{nums1[data.i], nums2[data.j]})
		if data.j+1 < len(nums2) {
			heap.Push(Data{data.i, data.j + 1, nums1[data.i] + nums2[data.j+1]})
		}
	}
	return ans
}

type Data struct {
	i, j int
	sum  int
}

type Heap []Data

func NewHeap() Heap {
	return Heap(make([]Data, 0))
}

func (h *Heap) Push(data Data) {
	n := len(*h)
	*h = append(*h, data)
	h.heapifyUp(n)
}

func (h *Heap) Pop() Data {
	ans := (*h)[0]
	n := len(*h)
	(*h)[0], (*h)[n-1] = (*h)[n-1], (*h)[0]
	*h = (*h)[:n-1]
	h.heapifyDown(0)
	return ans
}

func (h *Heap) heapifyUp(i int) {
	parent := (i - 1) / 2
	if parent >= 0 && (*h)[parent].sum > (*h)[i].sum {
		(*h)[parent], (*h)[i] = (*h)[i], (*h)[parent]
		h.heapifyUp(parent)
	}
}

func (h *Heap) heapifyDown(i int) {
	smallest := i
	left, right := 2*i+1, 2*i+2
	if left < len(*h) && (*h)[left].sum < (*h)[smallest].sum {
		smallest = left
	}
	if right < len(*h) && (*h)[right].sum < (*h)[smallest].sum {
		smallest = right
	}
	if smallest != i {
		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		h.heapifyDown(smallest)
	}
}

// Code Above Defines a Heap
