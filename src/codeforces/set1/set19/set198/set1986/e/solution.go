package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, m)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(a []int, k int) int {

	rem := make(map[int][]int)

	for _, v := range a {
		rem[v%k] = append(rem[v%k], v)
	}

	var res int

	pairty := len(a) % 2

	for _, vs := range rem {
		sort.Ints(vs)
		if len(vs)%2 == 1 {
			if pairty == 0 {
				// no answer
				return -1
			}
			// pairty == 1
			var pref []int
			var sum int
			for i, v := range vs {
				if i%2 == 1 {
					sum += v - vs[i-1]
				}
				pref = append(pref, sum)
			}
			best := sum
			sum = 0
			// 0, 1, 2, 3, 4
			for i := len(vs) - 1; i >= 0; i-- {
				if i%2 == 0 {
					best = min(best, sum+pref[i])
				} else {
					sum += vs[i+1] - vs[i]
				}
			}
			res += best
			pairty--
		} else {
			for i := 1; i < len(vs); i += 2 {
				res += vs[i] - vs[i-1]
			}
		}
	}

	return res / k
}

func solve1(a []int, k int) int {
	rem := make(map[int][]int)

	for _, num := range a {
		r := num % k
		rem[r] = append(rem[r], num)
	}
	n := len(a)
	var cnt int
	for _, v := range rem {
		cnt += len(v) & 1
	}
	if n%2 == 1 && cnt != 1 || n%2 == 0 && cnt > 0 {
		// 只能有一个是奇数的
		return -1
	}

	process := func(xs []int) int {
		sort.Ints(xs)
		// xs对k求余都是相同的 = r
		// 两个一组，改成相同的，把小的改变成大的
		// 但是如果是处于中间的那个，可以不用修改
		// 但是这个不一定是最大的，而可以是任何一个处于奇数位置的
		// 要算pref， suf
		if len(xs)%2 == 0 {
			var res int
			for i := 1; i < len(xs); i += 2 {
				res += (xs[i] - xs[i-1]) / k
			}
			return res
		}
		m := len(xs)
		pref := make([]int, m)
		for i := 1; i < m; i++ {
			pref[i] = pref[i-1]
			if i%2 == 1 {
				pref[i] += (xs[i] - xs[i-1]) / k
			}
		}
		best := 1 << 30
		var suf int
		for i := m - 1; i >= 0; i-- {
			if i%2 == 0 {
				best = min(best, suf+pref[i])
			} else {
				suf += (xs[i+1] - xs[i]) / k
			}
		}
		return best
	}

	var res int
	// posisition doesn't matter, only value maters
	for _, v := range rem {
		res += process(v)
	}

	return res
}
