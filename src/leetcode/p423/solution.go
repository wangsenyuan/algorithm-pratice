package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(originalDigits("owoztneoer"))
	fmt.Println(originalDigits("fviefuro"))

}

/**
--zero: 0
--one: 1
--two: 2
--three: 3
--four: 4
--five: 5
--six: 6
--seven: 7
eight: 8
nine: 9
**/
func originalDigits(s string) string {
	dict := make(map[byte]int)
	dict['x'] = 6
	dict['s'] = 7
	dict['v'] = 5
	dict['u'] = 4
	dict['z'] = 0
	dict['r'] = 3
	dict['w'] = 2
	dict['o'] = 1
	dict['n'] = 9
	dict['e'] = 8
	order := []byte{'x', 's', 'v', 'u', 'z', 'r', 'w', 'o', 'n', 'e'}
	name := make(map[byte]string)
	name['x'] = "six"
	name['s'] = "seven"
	name['v'] = "five"
	name['u'] = "four"
	name['z'] = "zero"
	name['r'] = "three"
	name['w'] = "two"
	name['o'] = "one"
	name['n'] = "nine"
	name['e'] = "eight"

	bs := []byte(s)
	sort.Sort(sortBytes(bs))

	words := make(map[byte]int)
	j := 0
	for i := 1; i <= len(bs); i++ {
		if i < len(bs) && bs[i] == bs[i-1] {
			continue
		}
		words[bs[j]] = i - j
		j = i
	}

	var res []byte
	for _, o := range order {
		for words[o] > 0 {
			res = append(res, byte(dict[o]+'0'))
			for j := 0; j < len(name[o]); j++ {
				words[name[o][j]]--
			}
		}
	}
	sort.Sort(sortBytes(res))
	return string(res)
}

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}
