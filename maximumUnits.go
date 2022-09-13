package main

import "sort"

// https://leetcode.com/problems/maximum-units-on-a-truck/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	boxIdx := 0
	capacity := truckSize
	totalUnits := 0
	for capacity > 0 && boxIdx < len(boxTypes) {
		boxes := boxTypes[boxIdx][0]

		if boxes > capacity {
			boxes = capacity
		}

		totalUnits += boxTypes[boxIdx][1] * boxes

		capacity -= boxes
		boxIdx++
	}

	return totalUnits
}
