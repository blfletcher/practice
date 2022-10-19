package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	level := 0
	queue := [][]TreeNode{{*root}}
	for true {
		currentLevel := queue[level]
		if len(currentLevel) == 0 {
			break
		}

		for len(currentLevel) > 0 {
			currentNode := currentLevel[0]
			currentLevel = currentLevel[1:]
			if len(result) <= level {
				result = append(result, []int{})
			}

			if level%2 == 0 {
				result[level] = append(result[level], currentNode.Val)
			} else {
				result[level] = append([]int{currentNode.Val}, result[level]...)
			}

			if len(queue) <= level+1 {
				queue = append(queue, []TreeNode{})
			}

			if currentNode.Left != nil {
				queue[level+1] = append(queue[level+1], *currentNode.Left)
			}
			if currentNode.Right != nil {
				queue[level+1] = append(queue[level+1], *currentNode.Right)
			}
		}
		level++
	}

	return result
}
