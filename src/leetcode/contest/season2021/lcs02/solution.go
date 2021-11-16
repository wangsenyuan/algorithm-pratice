package lcs02

import "sort"

func halfQuestions(questions []int) int {
	sort.Ints(questions)
	var cnts []int
	for i, j := 1, 0; i <= len(questions); i++ {
		if i == len(questions) || questions[i] > questions[i-1] {
			cnts = append(cnts, i-j)
			j = i
		}
	}
	n := len(questions) / 2
	sort.Ints(cnts)
	var sum int
	p := len(cnts)
	for p > 0 && sum < n {
		sum += cnts[p-1]
		p--
	}
	return len(cnts) - p
}
