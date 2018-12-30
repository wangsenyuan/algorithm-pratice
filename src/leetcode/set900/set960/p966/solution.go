package p966

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	org := make(map[string]int)
	low := make(map[string]int)
	vow := make(map[string]int)
	for i := 0; i < len(wordlist); i++ {
		org[wordlist[i]] = i + 1
		s := strings.ToLower(wordlist[i])
		if _, found := low[s]; !found {
			low[s] = i + 1
		}
		o := replaceVow(s)
		if _, found := vow[o]; !found {
			vow[o] = i + 1
		}
	}
	res := make([]string, len(queries))
	for j, query := range queries {
		if i, found := org[query]; found {
			//rule 1
			res[j] = wordlist[i-1]
			continue
		}
		s := strings.ToLower(query)
		if i, found := low[s]; found {
			res[j] = wordlist[i-1]
			continue
		}

		o := replaceVow(s)

		if i, found := vow[o]; found {
			res[j] = wordlist[i-1]
		}

	}
	return res
}

func replaceVow(s string) string {
	buf := []byte(s)
	for i := 0; i < len(buf); i++ {
		if isVow(buf[i]) {
			buf[i] = '*'
		}
	}
	return string(buf)
}

func isVow(a byte) bool {
	return a == 'a' || a == 'o' || a == 'e' || a == 'i' || a == 'u'
}
