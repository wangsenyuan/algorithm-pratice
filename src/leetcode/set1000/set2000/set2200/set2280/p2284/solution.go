package p2284

import "strings"

func largestWordCount(messages []string, senders []string) string {
	cnt := make(map[string]int)

	for i := 0; i < len(messages); i++ {
		s := senders[i]
		j := strings.Split(messages[i], " ")
		cnt[s] += len(j)
	}
	var res string

	for k, v := range cnt {
		if v > cnt[res] || v == cnt[res] && k > res {
			res = k
		}
	}
	return res
}
