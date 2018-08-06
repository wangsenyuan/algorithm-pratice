package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		P, Q := readTwoNums(scanner)
		people := readNNums(scanner, Q)
		res := solve(P, Q, people)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

func solve(P, Q int, people []int) int {
	sort.Ints(people)

	var loop func(a, b int) int
	cache := make(map[int]map[int]int)

	loop = func(a, b int) int {
		if _, founda := cache[a]; founda {
			if v, foundb := cache[a][b]; foundb {
				return v
			}
		}
		if cache[a] == nil {
			cache[a] = make(map[int]int)
		}
		j := sort.Search(Q, func(j int) bool {
			return people[j] > b
		})
		// Q[j] > b
		i := sort.Search(Q, func(i int) bool {
			return people[i] >= a
		})
		if i >= j {
			//no one release on this range
			cache[a][b] = 0
			return 0
		}
		// Q[i] >= a
		best := math.MaxInt32
		for k := i; k < j; k++ {
			c := people[k]
			tmp := (b - a) + loop(a, c-1) + loop(c+1, b)
			if tmp < best {
				best = tmp
			}
		}
		cache[a][b] = best
		return best
	}

	return loop(1, P)
}
