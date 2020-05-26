package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(reader, m)
		}

		res := solve(n, m, G)
		if !res {
			fmt.Println("-1")
			continue
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Printf("%d ", G[i][j])
			}
			fmt.Println()
		}
	}
}

func solve(n, m int, G [][]int) bool {
	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt[G[i][j]]++
		}
	}
	odds := make([]*Pair, 0, len(cnt))
	evens := make([]*Pair, 0, len(cnt))

	for k, v := range cnt {
		if v&1 == 1 {
			odds = append(odds, &Pair{k, v})
		} else {
			evens = append(evens, &Pair{k, v})
		}
	}

	if m&1 == 1 {
		if len(odds) > n {
			return false
		}

		if (n-len(odds))&1 == 1 {
			return false
		}

		// len(odds) <= n

		for i := 0; i < len(odds); i++ {
			G[i][m/2] = odds[i].first
			if odds[i].second > 1 {
				evens = append(evens, &Pair{odds[i].first, odds[i].second - 1})
			}
		}

		if len(odds) < n {
			var j int
			for i := len(odds); i+1 < n; i += 2 {
				for j < len(evens) && evens[j].second == 0 {
					j++
				}
				if j == len(evens) {
					return false
				}
				cur := evens[j]

				G[i][m/2] = cur.first
				G[i+1][m/2] = cur.first
				cur.second -= 2
			}
		}
	} else {
		if len(odds) > 0 {
			return false
		}
	}

	var p int
	for i := 0; i < n; i++ {
		for j, k := 0, m-1; j < k; j, k = j+1, k-1 {
			for evens[p].second == 0 {
				p++
			}
			if p == len(evens) {
				return false
			}
			cur := evens[p]

			G[i][j] = cur.first
			G[i][k] = cur.first
			cur.second -= 2
		}
	}

	return true
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
