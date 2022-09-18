package main

// https://leetcode.com/problems/number-of-provinces/

func checkCity(visited *[]bool, isConnected [][]int, city int) {
	for j := range isConnected[city] {
		connection := j

		if isConnected[city][connection] == 1 && !(*visited)[connection] {
			(*visited)[connection] = true
			checkCity(visited, isConnected, connection)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	provinces := 0

	for i := range isConnected {
		idx := i

		if !visited[idx] {
			checkCity(&visited, isConnected, idx)
			provinces++
		}
	}

	return provinces
}
