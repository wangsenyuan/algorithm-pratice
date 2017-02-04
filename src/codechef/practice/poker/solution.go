package main

import (
	"fmt"
	"sort"
)

func main() {
	var rank = make(map[byte]int)

	rank['A'] = 1

	for i := 1; i < 9; i++ {
		rank[byte(i + '1')] = i + 1
	}

	rank['T'] = 10
	rank['J'] = 11
	rank['Q'] = 12
	rank['K'] = 13

	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var a, b, c, d, e string
		fmt.Scanf("%s %s %s %s %s", &a, &b, &c, &d, &e)

		res := play(a, b, c, d, e, rank)

		fmt.Println(res)

		t--
	}
}

func play(a, b, c, d, e string, rank map[byte]int) string {
	ranks := make([]int, 5)
	ranks[0] = rank[a[0]]
	ranks[1] = rank[b[0]]
	ranks[2] = rank[c[0]]
	ranks[3] = rank[d[0]]
	ranks[4] = rank[e[0]]

	sort.Ints(ranks)

	x := kindType(ranks)
	y := suiteType(a, b, c, d, e)
	if x == 0 && y == 0 {
		return "royal flush"
	}

	if x == 1 && y == 0 {
		return "straight flush"
	}

	if x == 7 {
		return "four of a kind"
	}

	if x == 6 {
		return "full house"
	}

	if y == 0 {
		return "flush"
	}

	if x == 1 {
		return "straight"
	}

	if x == 4 {
		return "three of a kind"
	}

	if x == 5 {
		return "two pairs"
	}

	if x == 3 {
		return "pair"
	}

	return "high card"
}
func suiteType(a string, b string, c string, d string, e string) int {
	if a[1] == b[1] && b[1] == c[1] && c[1] == d[1] && d[1] == e[1] {
		return 0
	}
	return 1
}

func kindType(ranks []int) int {
	if ranks[0] == 1 && ranks[1] == 10 && ranks[2] == 11 && ranks[3] == 12 && ranks[4] == 13 {
		return 0
	}

	if ranks[1]+1 == ranks[2] && ranks[2]+1 == ranks[3] && ranks[3]+1 == ranks[4] &&
		(ranks[0]+1 == ranks[1] || ranks[4]+1 == ranks[0]) {
		return 1
	}

	cnt := make(map[int]int)

	for i := 0; i < len(ranks); i++ {
		cnt[ranks[i]]++
	}

	if len(cnt) == 5 {
		//all different
		return 2
	}

	if len(cnt) == 4 {
		//one pair, 2, 1, 1, 1
		return 3
	}
	if len(cnt) == 3 {
		//3, 1, 1
		for _, c := range cnt {
			if c == 3 {
				return 4
			}
		}
		//2, 2, 1
		return 5
	}

	if len(cnt) == 2 {
		//3, 2
		for _, c := range cnt {
			if c == 3 {
				return 6
			}
		}

		//4, 1
		return 7
	}
	return -1
}
