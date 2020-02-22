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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)

		W := make([]int, n)
		P := make([]int, n)
		V := make([]int, n)

		for i := 0; i < n; i++ {
			line := readNNums(scanner, 3)
			W[i] = line[0]
			P[i] = line[1]
			V[i] = line[2]
		}

		res := solve(n, P, W, V)

		fmt.Println(res)
	}
}

func solve(n int, P []int, W []int, V []int) int64 {
	if n < 20 {
		return solve1(n, P, W, V)
	}

	// return solve2(n, P, W, V)
	return solve3(n, P, W, V)
}

func solve1(n int, P []int, W []int, V []int) int64 {
	N := 1 << uint(n)

	dp := make([]int64, N)
	dp[1] = int64(V[0])

	res := dp[1]

	for i := 1; i < n; i++ {
		ii := 1 << uint(i)

		for state := 0; state < ii; state++ {
			var took int
			for j := 0; j < i && took <= W[i]; j++ {
				if state&(1<<uint(j)) > 0 {
					took += P[j]
				}
			}

			if took > W[i] {
				continue
			}

			cur := state | ii
			dp[cur] = max(dp[cur], dp[state]+int64(V[i]))

			res = max(res, dp[cur])
		}
	}

	return res
}

func solve2(n int, P []int, W []int, V []int) int64 {
	dp := make(map[int64]int64)
	dp[0] = 0

	for i := 0; i < n; i++ {
		p := int64(P[i])
		w := int64(W[i])
		v := int64(V[i])
		fp := make(map[int64]int64)
		for k, x := range dp {
			fp[k] = x

			if k <= w {
				fp[k+p] = max(fp[k+p], dp[k]+v)
			}
		}
		dp = fp
	}
	var res int64

	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

const INF = math.MaxInt64

func solve3(n int, P []int, W []int, V []int) int64 {
	half := n / 2
	N := 1 << uint(half)

	var res int64

	vec := make([]Pair, 0, N)

	for state := 0; state < N; state++ {
		ok := true
		var took, earn int64
		for i := 0; i < half; i++ {
			if state&(1<<uint(i)) > 0 {
				if took > int64(W[i]) {
					ok = false
					break
				}
				took += int64(P[i])
				earn += int64(V[i])
			}
		}
		if !ok {
			continue
		}
		vec = append(vec, Pair{took, earn})
		res = max(res, earn)
	}

	sort.Sort(Pairs(vec))

	ph := make([]Pair, len(vec))

	var p int

	for i := 0; i < len(vec); i++ {
		if p == 0 || ph[p-1].second < vec[i].second {
			ph[p] = vec[i]
			p++
		}
	}

	ph = ph[:p]

	secondHalf := n - half

	N = 1 << uint(half)

	for state := 0; state < N; state++ {
		var took, earn int64
		ok := true
		var f int64 = INF

		for i := 0; i < secondHalf; i++ {
			if state&(1<<uint(i)) == 0 {
				continue
			}
			if took > int64(W[i+half]) {
				ok = false
				break
			}
			f = min(f, int64(W[i+half])-took)
			took += int64(P[i+half])
			earn += int64(V[i+half])
		}

		if !ok {
			continue
		}

		j := sort.Search(p, func(j int) bool {
			return ph[j].first > f
		})
		j--
		res = max(res, ph[j].second+earn)
	}

	return res
}

type Pair struct {
	first, second int64
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second > ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
