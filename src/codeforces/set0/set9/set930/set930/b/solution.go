package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	fmt.Printf("%.9f\n", res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(s string) float64 {
	n := len(s)

	pos := make([][]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = make([]int, 0, 1)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		pos[x] = append(pos[x], i)
	}
	var win int

	cnt := make([]int, 26)
	for _, xs := range pos {

		var best int
		for d := 0; d < n; d++ {
			for _, i := range xs {
				j := (i + d) % n
				cnt[int(s[j]-'a')]++
			}
			var tmp int
			for i := 0; i < 26; i++ {
				if cnt[i] == 1 {
					tmp++
				}
				cnt[i] = 0
			}
			best = max(best, tmp)
		}
		win += best
	}

	return float64(win) / float64(n)
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
