package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	A := readNNums(reader, n)

	fmt.Println(solve(n, A))
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, A []int) int64 {
	ps := make([]Pair, n)
	var sum int64
	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], i}
		sum += int64(A[i])
	}
	sort.Sort(Pairs(ps))
	var best int64 = math.MinInt64
	if n%3 == 1 {
		best = sum
	}
	vis := make([]bool, n)

	check := func() bool {
		if vis[0] == vis[1] {
			return true
		}
		for i := 2; i < n; i++ {
			if vis[i] != vis[i%2] {
				return true
			}
		}
		return false
	}

	var tmp int64
	var safe bool
	for i := 0; i < n; i++ {
		cur := ps[i]
		tmp += int64(cur.first)
		vis[ps[i].second] = true
		if !safe {
			if ps[i].second > 0 {
				safe = vis[ps[i].second-1]
			}
			if !safe && ps[i].second < n-1 {
				safe = vis[ps[i].second+1]
			}
		}
		if (n+(i+1))%3 == 1 {
			tmp2 := -1 * tmp
			tmp3 := sum - tmp
			if !safe && ((i+1) == n/2 || (n-(i+1)) == n/2) {
				// not safe and alternating signs
				if !check() && i+1 < n {
					tmp2 += int64(cur.first) - int64(ps[i+1].first)
					tmp3 += int64(cur.first) - int64(ps[i+1].first)
				}
			}
			if tmp3+tmp2 > best {
				best = tmp3 + tmp2
			}
		}
	}

	return best
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
