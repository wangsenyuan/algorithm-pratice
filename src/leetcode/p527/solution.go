package main

import (
	"fmt"
	"strconv"
)

func main() {
	dict := []string{"like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion"}
	abbr := wordsAbbreviation(dict)
	fmt.Println(abbr)
}

func wordsAbbreviation(dict []string) []string {
	abbrMap := process(dict, "")
	res := make([]string, len(dict))

	for i := 0; i < len(dict); i++ {
		res[i] = abbrMap[dict[i]]
	}
	return res
}

func process(dict []string, prefix string) map[string]string {
	am := make(map[string][]string)

	for _, word := range dict {
		a := abbr(word, prefix)
		if l, found := am[a]; found {
			am[a] = append(l, word)
		} else {
			am[a] = []string{word}
		}
	}

	res := make(map[string]string)

	for abbr, lst := range am {
		if len(lst) == 1 {
			res[prefix+lst[0]] = abbr
			continue
		}

		tail := make([]string, len(lst))
		for i := 0; i < len(lst); i++ {
			tail[i] = lst[i][1:]
		}
		subRes := process(tail, prefix+string(lst[0][0]))
		for sw, sa := range subRes {
			res[sw] = sa
		}
	}

	return res
}

func abbr(word string, prefix string) string {
	if len(word) <= 3 {
		return prefix + word
	}
	cnt := strconv.Itoa(len(word) - 2)
	return prefix + string(word[0]) + cnt + string(word[len(word)-1])
}
