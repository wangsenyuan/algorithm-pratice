package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	p1, p2, _, _, _ := process(reader)

	fmt.Println(p1, p2)
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

func process(reader *bufio.Reader) (p1 int, p2 int, friends [][]int, prices []int, pizzas [][]int) {
	n, m := readTwoNums(reader)
	friends = make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var k, pos int
		pos = readInt(s, 0, &k) + 1
		friends[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &friends[i][j]) + 1
		}
	}
	prices = make([]int, m)
	pizzas = make([][]int, m)
	for i := range m {
		s, _ := reader.ReadBytes('\n')
		var k, pos int
		pos = readInt(s, 0, &prices[i]) + 1
		pos = readInt(s, pos, &k) + 1
		pizzas[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &pizzas[i][j]) + 1
		}
	}
	p1, p2 = solve(friends, prices, pizzas)
	return
}

const D = 9

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func getState(v []int) int {
	var res int
	for _, i := range v {
		res |= 1 << (i - 1)
	}
	return res
}

func solve(friends [][]int, prices []int, pizzas [][]int) (int, int) {
	freq := make([]int, 1<<D)

	for _, f := range friends {
		s := getState(f)
		freq[s]++
	}

	dp := make([][]pair, 1<<D)
	for i := 0; i < 1<<D; i++ {
		dp[i] = make([]pair, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = pair{inf, -1}
		}
	}
	n := len(prices)
	for i := range n {
		s := getState(pizzas[i])

		if prices[i] <= dp[s][0].first {
			dp[s][1] = dp[s][0]
			dp[s][0] = pair{prices[i], i}
		} else if prices[i] <= dp[s][1].first {
			dp[s][1] = pair{prices[i], i}
		}
	}

	getCost := func(a, b int) (res int, j1 int, j2 int) {
		res = dp[a][0].first
		j1 = dp[a][0].second + 1
		if a != b {
			res += dp[b][0].first
			j2 = dp[b][0].second + 1
		} else {
			res += dp[a][1].first
			j2 = dp[a][1].second + 1
		}
		return
	}
	best := []int{0, inf, 0, 0}

	for s1 := 1; s1 < 1<<D; s1++ {
		for s2 := 1; s2 < 1<<D; s2++ {
			cost, j1, j2 := getCost(s1, s2)
			if cost >= inf {
				continue
			}
			s := s1 | s2
			var cnt int
			for x := s; x > 0; x = (x - 1) & s {
				cnt += freq[x]
			}
			if cnt > best[0] || cnt == best[0] && cost < best[1] {
				best[0] = cnt
				best[1] = cost
				best[2] = j1
				best[3] = j2
			}
		}
	}

	return best[2], best[3]
}

func solve1(friends [][]int, prices []int, pizzas [][]int) (int, int) {
	// fp[state] = 能够使用state的pizza能够取悦的人数

	freq := make([]int, 1<<D)

	for _, f := range friends {
		var state int
		for _, i := range f {
			state |= 1 << (i - 1)
		}
		freq[state]++
	}

	fp := make([]int, 1<<D)
	for state := 1; state < 1<<D; state++ {
		i := bits.OnesCount(uint(state))
		for s1 := (state - 1) & state; s1 > 0; s1 = (s1 - 1) & state {
			j := bits.OnesCount(uint(s1))
			if (i-j)%2 == 1 {
				fp[state] += fp[s1]
			} else {
				fp[state] -= fp[s1]
			}
		}
		fp[state] += freq[state]
	}
	// dp[state] = 这个状态下的最小price
	dp := make([]pair, 1<<D)
	for i := range 1 << D {
		dp[i] = pair{inf, -1}
	}
	n := len(prices)
	for i := range n {
		var state int
		for _, j := range pizzas[i] {
			state |= 1 << (j - 1)
		}
		if dp[state].first > prices[i] {
			dp[state].first = prices[i]
			dp[state].second = i
		}
	}

	// now try to find the answer
	best := []int{0, inf, 0, 0}

	for s1 := 0; s1 < 1<<D; s1++ {
		for s2 := 0; s2 < 1<<D; s2++ {
			cnt := fp[s1|s2]
			cost := dp[s1].first + dp[s2].first

			if cost >= inf {
				continue
			}

			if cnt > best[0] || cnt == best[0] && cost < best[1] {
				best[0] = cnt
				best[1] = cost
				best[2] = dp[s1].second + 1
				best[3] = dp[s2].second + 1
			}
		}
	}

	if best[2] == 0 || best[2] == best[3] {
		best[2] = best[3]
		best[3] = 0
	}
	if best[2] == 0 {
		for i := range len(prices) {
			if best[2] == 0 || prices[i] < prices[best[2]-1] {
				best[2] = i + 1
			}
		}
	}
	if best[3] == 0 {
		// 这里再去找到最便宜的那个
		for i := range len(prices) {
			if best[2] == i+1 {
				continue
			}
			if best[3] == 0 || prices[i] < prices[best[3]-1] {
				best[3] = i + 1
			}
		}
	}

	return best[2], best[3]
}
