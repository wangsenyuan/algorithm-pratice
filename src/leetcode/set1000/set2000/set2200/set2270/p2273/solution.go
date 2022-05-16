package p2273

import "sort"

func removeAnagrams(words []string) []string {
	var res []string

	for i := 0; i < len(words); {
		cur := anagrame(words[i])
		res = append(res, words[i])
		j := i + 1

		for j < len(words) && anagrame(words[j]) == cur {
			j++
		}
		i = j
	}

	return res
}

func anagrame(word string) string {
	buf := []byte(word)
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})
	return string(buf)
}
