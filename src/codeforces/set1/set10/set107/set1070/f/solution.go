package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	people := make([][]int, n)
	for i := 0; i < n; i++ {
		people[i] = readNNums(reader, 2)
	}

	res := solve(people)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(people [][]int) int {
	var res int
	var alice []int
	var bob []int
	var other []int

	n := len(people)
	var m int
	for i := 0; i < n; i++ {
		cur := people[i]
		if cur[0] == 11 {
			res += cur[1]
			m++
		} else if cur[0] == 10 {
			alice = append(alice, cur[1])
		} else if cur[0] == 1 {
			bob = append(bob, cur[1])
		} else {
			other = append(other, cur[1])
		}
	}

	sort.Ints(alice)
	reverse(alice)
	sort.Ints(bob)
	reverse(bob)

	var it int
	for it < len(alice) && it < len(bob) {
		res += alice[it]
		res += bob[it]
		it++
	}

	sort.Ints(other)
	reverse(other)
	// 如果选中了x支持alice,
	// m + x >= y + z
	// m + y >= x + z
	// 如果z <= m
	// x - y >= z - m
	// x - y <= m - z
	// 选择支持bob的人越多，那么other的人就要越少
	// 在没有y的情况下，x + z <= m

	addOther := func(a []int, b []int) int {
		// x keeps
		x, y := it, it
		var i, z int
		stack := make([]int, len(b))
		var best int
		var sum int
		for y < len(a) || i < len(b) {
			if i == len(b) || y < len(a) && a[y] >= b[i] {
				sum += a[y]
				y++
			} else {
				sum += b[i]
				stack[z] = b[i]
				i++
				z++
			}

			for (m+x < y+z || m+y < x+z) && z > 0 {
				sum -= stack[z-1]
				z--
			}
			if m+x >= y+z && m+y >= x+z {
				best = max(best, sum)
			}
		}

		return best
	}

	if len(alice) > it {
		res += addOther(alice, other)
	} else {
		res += addOther(bob, other)
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
