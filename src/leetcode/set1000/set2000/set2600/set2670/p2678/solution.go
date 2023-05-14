package p2678

import "strconv"

func countSeniors(details []string) int {
	var res int
	n := len(details)

	for i := 0; i < n; i++ {
		res += check(details[i])
	}

	return res
}

func check(detail string) int {
	s := detail[11:13]
	age, _ := strconv.Atoi(s)
	if age > 60 {
		return 1
	}
	return 0
}
