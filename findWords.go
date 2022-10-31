package main

// https://leetcode.com/problems/word-search-ii/

import "fmt"

func dfs(word string, wordIdx, m, n int, charMap map[byte][][]int, board [][]byte, visited map[string]bool, result []byte) []byte {
	if wordIdx >= len(word) {
		return result
	}

	char := board[m][n]
	if _, ok := charMap[char]; !ok {
		return result
	}

	if char != word[wordIdx] {
		return result
	}

	result = append(result, word[wordIdx])

	if m-1 >= 0 {
		nn := fmt.Sprintf("%d,%d", m-1, n)
		if _, ok := visited[nn]; !ok {
			visited[nn] = true
			nrs := dfs(word, wordIdx+1, m-1, n, charMap, board, visited, result)
			if string(nrs) == word {
				return nrs
			}
			delete(visited, nn)
		}
	}

	if n-1 >= 0 {
		nn := fmt.Sprintf("%d,%d", m, n-1)
		if _, ok := visited[nn]; !ok {
			visited[nn] = true
			nrs := dfs(word, wordIdx+1, m, n-1, charMap, board, visited, result)
			if string(nrs) == word {
				return nrs
			}
			delete(visited, nn)
		}
	}

	if m+1 < len(board) {
		nn := fmt.Sprintf("%d,%d", m+1, n)
		if _, ok := visited[nn]; !ok {
			visited[nn] = true
			nrs := dfs(word, wordIdx+1, m+1, n, charMap, board, visited, result)
			if string(nrs) == word {
				return nrs
			}
			delete(visited, nn)
		}
	}

	if n+1 < len(board[0]) {
		nn := fmt.Sprintf("%d,%d", m, n+1)
		if _, ok := visited[nn]; !ok {
			visited[nn] = true
			nrs := dfs(word, wordIdx+1, m, n+1, charMap, board, visited, result)
			if string(nrs) == word {
				return nrs
			}
			delete(visited, nn)
		}
	}

	return result
}

func findWords(board [][]byte, words []string) []string {
	chars := map[byte][][]int{}
	for m := range board {
		for n := range board[m] {
			if _, ok := chars[board[m][n]]; ok {
				chars[board[m][n]] = append(chars[board[m][n]], []int{m, n})
			} else {
				chars[board[m][n]] = [][]int{{m, n}}
			}
		}
	}

	result := []string{}
	for i := range words {
		word := words[i]

		if _, ok := chars[word[0]]; ok {
			indices := chars[word[0]]

			for s := range indices {
				initial := indices[s]
				nn := fmt.Sprintf("%d,%d", initial[0], initial[1])
				visited := map[string]bool{nn: true}
				res := dfs(word, 0, initial[0], initial[1], chars, board, visited, []byte{})
				if string(res) == word {
					result = append(result, word)
					break
				}
			}
		}
	}

	return result
}
