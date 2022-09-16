package main

import (
	"sort"
)

// https://leetcode.com/problems/combination-sum/

/*
result := []int{250} // fixed size array, zero-filled

func buildResult(start, candidates, result, results, target)
  for idx = start; idx < len(candidates); idx++
    // try adding cand
	cand = candidates[idx]
    result += cand
    subtract from target
	// exact match
	if target == 0
	  add result to results
	  // copy result into new array of correct (trimmed) size
	  return
	// failed resut
	else if target < candidate
	  return
	// continue building
    else
      return buildResult(idx, candidates, result, results, target)
	// undo cand
	target += cand
	result -= cand
*/

// algorithm
/***********

sort the candidate array

candidateSum(candidates, target, result) (results, newTarget)
	iterate through candidates
		newTarget = target - current candidate
		// 8 - 2 = 6
		if newTarget == 0
			save results
		else if newTarget > current candidate
			add current candidate to results
			// [2]
			results, newTarget = duplicateCandidate(candidate, newTarget, results)
			// [2, 2, 2, 2], 0
			if newTarget == 0
				save results
			if nextCandidate exists, while newTarget < nextCandidate and len(results) > 0
				add value from last element of results to newTarget
				remove lastElement of results
				// [], 8
	return results arrays

duplicateCandidate(candidate, target, results) (results, newTarget)
// 2, 6, [2]
// 2, 4, [2, 2]
// 2, 0, [2, 2, 2]
	newTarget = target - current candidate
	// 6 - 2 = 4
	// 4 - 2 = 2
	// 2 - 2 = 0
	if newTarget < currentCandidate
		return results, newTarget

	add currentCandidate to results
	// [2, 2]
	// [2, 2, 2]
	// [2, 2, 2]
	if newTarget == 0
		return results, newTarget

	return duplicateCandidate(candidate, newTarget, results)
	// 2, 4, [2, 2]
	// 2, 2, [2, 2, 2]


************/

func handleCandidate(candidate, target int, result []int) ([]int, int) {
	if target < candidate {
		return result, target
	}

	target -= candidate

	result = append(result, candidate)
	if target == 0 {
		return result, target
	}

	return handleCandidate(candidate, target, result)
}

/*
func buildResult(start, candidates, result, results, target)
  for idx = start; idx < len(candidates); idx++
    // try adding cand
	cand = candidates[idx]
    result += cand
    subtract from target
	// exact match
	if target == 0
	  add result to results
	  return
	// failed resut
	else if target < candidate
	  return
	// continue building
    else
      return buildResult(idx, candidates, result, results, target)
	// undo cand
	target += cand
	result -= cand
*/

// 2, 3, 5

func buildResult(idx, target int, candidates, result []int, results [][]int) [][]int {
	if target == 0 {
		cpy := make([]int, len(result))
		copy(cpy, result)

		results = append(results, cpy)

		return results
	} else if target < 0 {
		return results
	}

	for j := idx; j < len(candidates); j++ {
		candidate := candidates[j]

		result = append(result, candidate)
		// we should have preserved target
		results = buildResult(j, target-candidate, candidates, result, results)
		result = result[:len(result)-1]

		/*target -= candidate

		if target < candidate {
			if target == 0 {
				cpy := make([]int, len(result))
				copy(cpy, result)

				results = append(results, cpy)
			}
			//target += candidate
			result = result[:len(result)-1]
			fmt.Printf("\n\n>>[%d] [%d] [%d] [%d] %+v\n\n\n", j, target, candidate, target+candidate, result)
			return result, results
			//return buildResult(j, target, candidates, result, results)
		}

		fmt.Printf("\n\n<<[%d] [%d] [%d] [%d] %+v\n\n\n", j, target, candidate, target+candidate, result)
		result, results = buildResult(j, target, candidates, result, results)
		fmt.Printf("\n\n[%d] [%d] next iteration %+v\n\n\n", j, target, result)*/
	}

	return results
}

// [2]
//   [2 2]
//     [2 2 2]
//   [2 2]
//     [2 2 3]
//   [2 2]
//     [2 2 6]
// [2]
//   [2 3]

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	return buildResult(0, target, candidates, []int{}, [][]int{})
}
