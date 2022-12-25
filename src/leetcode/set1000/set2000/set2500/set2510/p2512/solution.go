package p2512

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	pos := arrayToSet(positive_feedback)
	neg := arrayToSet(negative_feedback)

	n := len(report)

	scores := make([]int, n)

	for i := 0; i < n; i++ {
		var sc int

		words := strings.Split(report[i], " ")

		for _, w := range words {
			if pos[w] {
				sc += 3
			} else if neg[w] {
				sc--
			}
		}

		scores[i] = sc
	}

	type Pair struct {
		first  int
		second int
	}

	students := make([]Pair, n)

	for i := 0; i < n; i++ {
		students[i] = Pair{scores[i], student_id[i]}
	}
	sort.Slice(students, func(i, j int) bool {
		if students[i].first > students[j].first {
			return true
		}
		if students[i].first == students[j].first {
			return students[i].second < students[j].second
		}
		return false
	})
	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = students[i].second
	}

	return res
}

func arrayToSet(arr []string) map[string]bool {
	res := make(map[string]bool)

	for _, s := range arr {
		res[s] = true
	}

	return res
}
