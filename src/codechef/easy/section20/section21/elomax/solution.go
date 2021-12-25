package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		R := readNNums(reader, n)
		C := make([][]int, n)
		for i := 0; i < n; i++ {
			C[i] = readNNums(reader, m)
		}
		fmt.Println(solve(n, m, R, C))
	}
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

func solve(n int, m int, R []int, C [][]int) int {
	bestRatingAt := make([]int, n)
	for i := 0; i < n; i++ {
		C[i][0] += R[i]
	}

	for j := 1; j < m; j++ {
		for i := 0; i < n; i++ {
			C[i][j] += C[i][j-1]
			if C[i][j] > C[i][bestRatingAt[i]] {
				bestRatingAt[i] = j
			}
		}
	}
	curRanking := make([]int, n)
	bestRanking := make([]int, n)
	for i := 0; i < n; i++ {
		bestRanking[i] = n
	}
	pp := make([]Pair, n)

	fillRanking := func(j int) {
		for i := 0; i < n; i++ {
			pp[i] = Pair{C[i][j], i}
		}
		sort.Sort(PP(pp))
		curRanking[pp[0].second] = 1
		for i := 1; i < n; i++ {
			curRanking[pp[i].second] = curRanking[pp[i-1].second]
			if pp[i].first < pp[i-1].first {
				curRanking[pp[i].second] = i + 1
			}
		}
	}

	for j := 0; j < m; j++ {
		fillRanking(j)
		for i := 0; i < n; i++ {
			bestRanking[i] = min(bestRanking[i], curRanking[i])
		}
	}

	var cnt int
	for j := 0; j < m; j++ {
		fillRanking(j)

		for i := 0; i < n; i++ {
			if curRanking[i] == bestRanking[i] {
				if bestRatingAt[i] == j {
					cnt++
				}
				bestRanking[i] = n + 1
			}
		}
	}
	return n - cnt
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type PP []Pair

func (pp PP) Len() int {
	return len(pp)
}

func (pp PP) Less(i, j int) bool {
	return pp[i].first > pp[j].first
}

func (pp PP) Swap(i, j int) {
	pp[i], pp[j] = pp[j], pp[i]
}
