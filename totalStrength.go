package main

// https://leetcode.com/problems/sum-of-total-strength-of-wizards/

const modulus = (10*10*10*10*10*10*10*10*10 + 7)

// teamPower function: iterates through the subarray, maintains a min while summing the elements, returns power
func teamPower(team []int) int64 {
	minPower := int64(team[0])
	totalPower := int64(team[0])

	for i := 1; i < len(team); i++ {
		val := int64(team[i])

		if val < minPower {
			minPower = val
		}
		totalPower += val
	}

	return minPower * totalPower
}

// algorithm
/***********

// 1 <= strength.length <= 105
// 1 <= strength[i] <= 109

// it's almost a power set, but order must be maintained
// slide a window across the array
//   loop that iterates over array
//   subloop idx = current loop idx >> expands the window through the array
//     within the subloop, calculate the power
//     totalPower += teamPower(i, j)
************/

func totalStrength(strength []int) int {
	totalPower := int64(0)
	for i := 0; i < len(strength); i++ {
		minTeamPower := int64(strength[i])
		teamPower := minTeamPower
		totalPower += minTeamPower * teamPower
		for j := i + 1; j < len(strength); j++ {
			currPower := int64(strength[j])
			if currPower < minTeamPower {
				minTeamPower = currPower
			}
			teamPower += currPower
			totalPower += minTeamPower * teamPower
		}
	}

	return int(totalPower % modulus)
}
