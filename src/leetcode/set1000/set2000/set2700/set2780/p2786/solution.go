package p2786

import "sort"

var vowels map[byte]int

func init() {
	vowels = make(map[byte]int)
	vowels['A'] = 1
	vowels['E'] = 2
	vowels['I'] = 3
	vowels['O'] = 4
	vowels['U'] = 5
	vowels['a'] = 6
	vowels['e'] = 7
	vowels['i'] = 8
	vowels['o'] = 9
	vowels['u'] = 10
}

const words = "AEIOUaeiou"

func sortVowels(s string) string {
	var vw []int
	for i := 0; i < len(s); i++ {
		if j, ok := vowels[s[i]]; ok {
			vw = append(vw, j)
		}
	}
	sort.Ints(vw)

	popVowel := func() byte {
		j := vw[0]
		vw = vw[1:]
		return words[j-1]
	}

	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		if _, ok := vowels[s[i]]; ok {
			buf[i] = popVowel()
		}
	}

	return string(buf)
}
