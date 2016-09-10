package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("ab", "aa"))
}

func isIsomorphic(s string, t string) bool {
	as := make(map[byte]byte)
	bs := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if c, mapped := as[a]; mapped {
			if b != c {
				return false
			}
		} else {
			as[a] = b
		}

		if c, mapped := bs[b]; mapped {
			if a != c {
				return false
			}
		} else {
			bs[b] = a
		}
	}

	return true
}
