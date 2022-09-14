package main

// https://leetcode.com/problems/integer-to-roman/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	indexes := map[*Node]int{}

	current := head
	copy := []Node{}
	for current != nil {
		indexes[current] = len(copy)
		copy = append(copy, Node{
			Val: current.Val,
		})
		current = current.Next
	}

	current = head
	for current != nil {
		idx := indexes[current]

		if idx < len(copy)-1 {
			copy[idx].Next = &copy[idx+1]
		}

		if current.Random != nil {
			rdx := indexes[current.Random]
			copy[idx].Random = &copy[rdx]
		}
		current = current.Next
	}

	if len(copy) > 0 {
		return &copy[0]
	}

	return nil
}

// build map of node->index
// for each original node
//    create copy
//    add copy to array (assemble all the nodes)
// for each copied node
//    initialize next, random (wire the nodes)
