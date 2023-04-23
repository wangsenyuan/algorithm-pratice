package lcp73

import "strings"

func adventureCamp(expeditions []string) int {
	known := make(map[string]int)

	ss := strings.Split(expeditions[0], "->")

	for _, s := range ss {
		known[s]++
	}

	res := -1
	var cnt int
	for i := 1; i < len(expeditions); i++ {
		xs := strings.Split(expeditions[i], "->")
		tmp := make(map[string]int)
		for _, x := range xs {
			if len(x) > 0 && known[x] == 0 {
				tmp[x]++
			}
		}

		if len(tmp) > 0 && (res < 0 || len(tmp) > cnt) {
			cnt = len(tmp)
			res = i
		}

		for _, x := range xs {
			if len(x) > 0 {
				known[x]++
			}
		}
	}

	return res
}
