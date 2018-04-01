package p811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	cnt := make(map[string]int)

	for _, cpdomain := range cpdomains {
		tmp := strings.Split(cpdomain, " ")
		a, b := tmp[0], tmp[1]
		x, _ := strconv.Atoi(a)

		for i := 0; i < len(b); i++ {
			if i == 0 {
				cnt[b] += x
			} else if b[i] == '.' {
				cnt[b[i+1:]] += x
			}
		}
	}

	res := make([]string, len(cnt))
	var idx int
	for k, v := range cnt {
		res[idx] = fmt.Sprintf("%d %s", v, k)
		idx++
	}
	return res
}
