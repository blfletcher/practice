package main

func getFood(grid [][]byte) int {
	queue := [][]int{}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '*' {
				queue = append(queue, []int{i, j})
				break
			}
		}
	}

	path := 0
	for len(queue) > 0 {
		n := len(queue)
		path++
		for i := 0; i < n; i++ {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, next := range directions {
				nr := r + next[0]
				nc := c + next[1]

				if nr < len(grid) && nr >= 0 && nc < len(grid[0]) && nc >= 0 {
					if grid[nr][nc] == '#' {
						return path
					}
					if grid[nr][nc] == 'O' {
						grid[nr][nc] = 'X'
						queue = append(queue, []int{nr, nc})
					}
				}
			}
		}
	}

	return -1
}
