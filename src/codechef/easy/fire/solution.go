package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n, s, k int
		scanner.Scan()
		scanner.Scan()
		bs := scanner.Bytes()
		x := readInt(bs, 0, &n)
		x = readInt(bs, x+1, &s)
		s--
		x = readInt(bs, x+1, &k)
		neighbors := make([][]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			bs = scanner.Bytes()
			var m int
			x = readInt(bs, 0, &m)
			neighbors[j] = make([]int, m)
			for a := 0; a < m; a++ {
				x = readInt(bs, x+1, &neighbors[j][a])
				neighbors[j][a]--
			}
		}
		important := make([]int, k)
		x = -1
		scanner.Scan()
		bs = scanner.Bytes()
		for j := 0; j < k; j++ {
			x = readInt(bs, x+1, &important[j])
			important[j]--
		}
		res := solve(n, s, neighbors, important)
		if res {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func solve(n int, s int, neighbors [][]int, important []int) bool {
	//fmt.Printf("[debug] %d %d %v %v", n, s, neighbors, important)
	special := make([]bool, n)
	for _, x := range important {
		special[x] = true
	}

	var f func(p int, v int) bool
	f = func(p int, v int) bool {
		if special[v] {
			return false
		}

		cc := 0
		for _, w := range neighbors[v] {
			if w == p {
				continue
			}
			cc++
			if f(v, w) {
				return true
			}
		}
		return cc <= 1
	}

	return f(-1, s)
}
