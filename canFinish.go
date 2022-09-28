package main

// https://leetcode.com/problems/course-schedule/

func canFinish(numCourses int, prerequisites [][]int) bool {
	reqs := make([]int, numCourses)
	courses := make([][]int, numCourses)

	for i := range prerequisites {
		reqs[prerequisites[i][0]]++
		courses[prerequisites[i][1]] = append(courses[prerequisites[i][1]], prerequisites[i][0])
	}

	queue := []int{}
	for i := range reqs {
		if reqs[i] == 0 {
			queue = append(queue, i)
		}
	}

	results := []int{}
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		results = append(results, course)

		for i := range courses[course] {
			reqs[courses[course][i]]--
			if reqs[courses[course][i]] == 0 {
				queue = append(queue, courses[course][i])
			}
		}
	}

	return len(results) == numCourses
}
