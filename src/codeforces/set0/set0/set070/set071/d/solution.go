package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	for _, row := range res {
		fmt.Println(row)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (cards [][]string, res []string) {
	n, m := readTwoNums(reader)
	cards = make([][]string, n)
	for i := range n {
		s := readString(reader)
		cards[i] = strings.Split(s, " ")[:m]
	}
	res = solve(cards)
	return
}

const SUITE = "SDHC"
const NUM = "A23456789TJQK"

func parse(card string) (rank int, suite int) {
	if card == "J1" {
		suite = 1
	} else if card == "J2" {
		suite = 2
	} else {
		rank = strings.IndexByte(NUM, card[0]) + 1
		suite = strings.IndexByte(SUITE, card[1])
	}
	return
}

func solve(cards [][]string) (res []string) {

	n := len(cards)
	m := len(cards[0])
	cur := make([][]int, n)

	pos := make([][]int, 14)
	for i := range 14 {
		pos[i] = make([]int, 4)
		for j := range 4 {
			pos[i][j] = -1
		}
	}

	for i, row := range cards {
		cur[i] = make([]int, m)
		for j, x := range row {
			r, s := parse(x)
			pos[r][s] = i*m + j
			if r == 0 {
				// joker 的位置需要特殊处理
				cur[i][j] = -s
			} else {
				cur[i][j] = r*4 + s
			}
		}
	}

	var free []int

	for i := 1; i <= 13; i++ {
		for j := range 4 {
			if pos[i][j] < 0 {
				free = append(free, i*4+j)
			}
		}
	}

	// joker都被交换了，看是否能成功
	check := func() [][]int {
		var good [][]int
		for i := 0; i+3 <= n; i++ {
			for j := 0; j+3 <= m; j++ {
				var rank int
				var suite int
				for u := i; u < i+3; u++ {
					for v := j; v < j+3; v++ {
						rank |= 1 << (cur[u][v] / 4)
						suite |= 1 << (cur[u][v] % 4)
					}
				}
				if bits.OnesCount(uint(rank)) == 9 || bits.OnesCount(uint(suite)) == 1 {
					for _, cur := range good {
						if cur[0]+3 <= i || cur[1]+3 <= j || j+3 <= cur[1] {
							return [][]int{cur, {i, j}}
						}
					}
					good = append(good, []int{i, j})
				}
			}
		}
		return nil
	}

	replace := make([]int, 3)
	for j := range 3 {
		replace[j] = -1
	}

	var dfs func(i int, flag int) bool

	dfs = func(i int, flag int) bool {
		if i == 3 {
			arr := check()
			if len(arr) > 0 {
				res = append(res, "Solution exists.")

				if replace[1] < 0 && replace[2] < 0 {
					res = append(res, "There are no jokers.")
				} else if replace[1] >= 0 && replace[2] >= 0 {
					res = append(res, fmt.Sprintf("Replace J1 with %c%c and J2 with %c%c.", NUM[replace[1]/4-1], SUITE[replace[1]%4], NUM[replace[2]/4-1], SUITE[replace[2]%4]))
				} else if replace[1] >= 0 {
					res = append(res, fmt.Sprintf("Replace J1 with %c%c.", NUM[replace[1]/4-1], SUITE[replace[1]%4]))
				} else if replace[2] >= 0 {
					res = append(res, fmt.Sprintf("Replace J2 with %c%c.", NUM[replace[2]/4-1], SUITE[replace[2]%4]))
				}
				res = append(res, fmt.Sprintf("Put the first square to (%d, %d).", arr[0][0]+1, arr[0][1]+1))
				res = append(res, fmt.Sprintf("Put the second square to (%d, %d).", arr[1][0]+1, arr[1][1]+1))

				return true
			}
			return false
		}

		if pos[0][i] < 0 {
			return dfs(i+1, flag)
		}
		// let's swap
		r, c := pos[0][i]/m, pos[0][i]%m
		// 把joker换出来
		for j, card := range free {
			if (flag>>j)&1 == 0 {
				cur[r][c] = card
				replace[i] = card
				if dfs(i+1, flag|(1<<j)) {
					return true
				}
				replace[i] = -1
				cur[r][c] = -i
			}
		}
		return false
	}

	if dfs(1, 0) {
		return res
	}
	res = append(res, "No solution.")
	return res
}
