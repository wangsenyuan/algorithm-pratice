package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	email2Name := make(map[string][]int)

	n := len(accounts)

	set := make([]int, n)

	for i := 0; i < n; i++ {
		set[i] = i
		sort.Strings(accounts[i][1:])
		k := 1
		j := 1
		for j < len(accounts[i]) && k < len(accounts[i]) {
			accounts[i][j] = accounts[i][k]
			for k < len(accounts[i]) && accounts[i][j] == accounts[i][k] {
				k++
			}
			j++
		}
		accounts[i] = accounts[i][:j]
	}

	var find func(x int) int

	find = func(x int) int {
		if x == set[x] {
			return x
		}
		set[x] = find(set[x])
		return set[x]
	}

	for j, account := range accounts {
		for i := 1; i < len(account); i++ {
			pj := find(j)
			email := account[i]
			if _, found := email2Name[email]; !found {
				email2Name[email] = make([]int, 0, 10)
				email2Name[email] = append(email2Name[email], j)
			} else {
				email2Name[email] = append(email2Name[email], j)
				k := email2Name[email][0]
				pk := find(k)
				set[pj] = pk
			}
		}
	}

	merge := func(a, b []string) []string {
		c := make([]string, 0, len(a)+len(b))
		c = append(c, a[0])

		i, j := 1, 1

		for i < len(a) && j < len(b) {
			x, y := a[i], b[j]
			if x < y {
				c = append(c, x)
				i++
			} else if x > y {
				c = append(c, y)
				j++
			} else {
				c = append(c, x)
				i++
				j++
			}
		}

		for i < len(a) {
			c = append(c, a[i])
			i++
		}

		for j < len(b) {
			c = append(c, b[j])
			j++
		}

		return c
	}

	result := make([][]string, 0, n)
	resultMap := make(map[int]int)
	for i := 0; i < n; i++ {
		pi := find(i)

		if _, found := resultMap[pi]; !found {
			result = append(result, accounts[pi])
			resultMap[pi] = len(result) - 1
		}

		j := resultMap[pi]
		result[j] = merge(result[j], accounts[i])
	}

	return result
}

func main() {
	//test2()
	test3()
}

func test1() {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}

	result := accountsMerge(accounts)

	fmt.Println(result)
}

func test2() {
	accounts := [][]string{{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
		{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
		{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
		{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
	}
	result := accountsMerge(accounts)

	fmt.Println(result)
}

func test3() {
	accounts := [][]string{{"David", "David0@m.co", "David1@m.co"},
		{"David", "David3@m.co", "David4@m.co"},
		{"David", "David4@m.co", "David5@m.co"},
		{"David", "David2@m.co", "David3@m.co"},
		{"David", "David1@m.co", "David2@m.co"},
	}
	result := accountsMerge(accounts)

	fmt.Println(result)
}
