package main

import "fmt"

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	table := make(map[string]string)

	for n > 0 {
		n--
		var s, t string
		fmt.Scanf("%s %s", &s, &t)
		table[s] = t
	}

	for q > 0 {
		q--
		var s string
		fmt.Scanf("%s", &s)
		i := len(s)
		for i > 0 && s[i - 1] != '.' {
			i--
		}
		if i == 0 {
			fmt.Println("unknown")
			continue
		}
		extension := s[i:]
		if typ, found := table[extension]; found {
			fmt.Println(typ);
		} else {
			fmt.Println("unknown")
		}
	}
}
