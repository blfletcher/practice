package main

// https://leetcode.com/problems/reorder-data-in-log-files/

import (
	"sort"
	"strconv"
	"strings"
)

type wordEntry struct {
	key   string
	value string
}

func reorderLogFiles(logs []string) []string {
	wordEntries := []wordEntry{}
	digitLog := []string{}

	for i := range logs {
		lg := logs[i]

		lgArr := strings.Split(lg, " ")
		_, err := strconv.ParseInt(lgArr[1], 10, 64)
		if err == nil || strings.Contains(err.Error(), "value out of range") {
			digitLog = append(digitLog, lg)
		} else {
			wordEntries = append(wordEntries, wordEntry{
				key:   lgArr[0],
				value: strings.Join(lgArr[1:], " "),
			})
		}
	}

	sort.Slice(wordEntries, func(i, j int) bool {
		if wordEntries[i].value == wordEntries[j].value {
			return wordEntries[i].key < wordEntries[j].key
		}
		return wordEntries[i].value < wordEntries[j].value
	})

	result := []string{}
	for i := range wordEntries {
		word := wordEntries[i]

		result = append(result, word.key+" "+word.value)
	}

	result = append(result, digitLog...)

	return result
}
