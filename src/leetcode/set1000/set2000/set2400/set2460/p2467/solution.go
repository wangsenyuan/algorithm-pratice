package p2467

import (
	"fmt"
	"strconv"
)

func splitMessage(message string, limit int) []string {
	n := len(message)

	id_len := make([]int, n)

	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		id_len[i-1] = len(s)
	}

	// can split in <=x groups
	split := func(x int) bool {
		s := strconv.Itoa(x)
		m := len(s)
		// < / and >
		m += 3
		var id int
		for i := 0; i < n; {
			tmp_len := id_len[id]
			if tmp_len+m >= limit {
				// equals will not advance to next position
				return false
			}
			i += limit - (tmp_len + m)
			id++
		}

		return id <= x
	}

	l, r := 1, n+1

	for l < r {
		mid := (l + r) / 2
		if split(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r == n+1 {
		return nil
	}
	var buf []string
	s := "<%d/%d>"
	for i, j := 0, 0; i < n; j++ {
		x := fmt.Sprintf(s, (j + 1), r)
		ll := limit - len(x)
		buf = append(buf, message[i:min(n, i+ll)]+x)
		i += ll
	}

	return buf
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
