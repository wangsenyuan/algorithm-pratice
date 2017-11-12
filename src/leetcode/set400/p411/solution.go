package main

import (
	"fmt"
	"strconv"
)

func main() {
	//target := "abcdef"
	//target := "abc"
	/*for i := 1; i < len(target); i++ {
		fmt.Println(abbreviation(target, i))
	}*/
	//dict := []string{"ablade", "xxxxef", "abdefi"}
	//dict := []string{"abd", "acd", "acc"}
	target := "usa"
	dict := []string{"hel", "uus", "ssa", "uaa", "uss", "uua", "aaa", "uuu", "sss", "uul", "url", "uls", "uas"}
	fmt.Println(minAbbreviation(target, dict))
}

func minAbbreviation(target string, dictionary []string) string {
	for i := 1; i <= len(target); i++ {
		abbr := abbreviation(target, i)
		for _, x := range abbr {
			valid := true
			for _, dict := range dictionary {
				if validWordAbbreviation(dict, x) {
					valid = false
					break
				}
			}
			if valid {
				return x
			}
		}
	}
	return ""
}

func abbreviation(str string, k int) []string {
	if k == 0 {
		return []string{""}
	}
	var res []string

	if k == 1 {
		res = append(res, strconv.Itoa(len(str)))
		if len(str) == 1 {
			res = append(res, str)
		}
		return res
	}
	checked := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		x := str[i]
		for a := 0; a < k; a++ {
			b := k - 1 - a

			if i == 0 && a != 0 {
				continue
			}

			if i > 0 && a == 0 {
				continue
			}

			if i == len(str)-1 && b != 0 {
				continue
			}
			if i < len(str)-1 && b == 0 {
				continue
			}

			for _, lx := range abbreviation(str[:i], a) {
				for _, rx := range abbreviation(str[i+1:], b) {
					sx := lx + string(x) + rx
					if checked[sx] {
						continue
					}
					checked[sx] = true
					res = append(res, sx)
				}
			}
		}
	}
	return res
}

func validWordAbbreviation(word string, abbr string) bool {
	x := 0
	for i := 0; i < len(abbr); i++ {
		if abbr[i] >= '0' && abbr[i] <= '9' {
			if x == 0 && abbr[i] == '0' {
				return false
			}
			x = x*10 + int(abbr[i]-'0')
			continue
		}
		if x >= len(word) {
			return false
		}
		word = word[x:]
		x = 0
		if abbr[i] != word[0] {
			return false
		}
		word = word[1:]
	}

	return x == len(word)
}
