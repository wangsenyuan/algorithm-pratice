package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		S := make([]Pair, n+2)
		for i := 1; i <= n; i++ {
			line, _ := reader.ReadBytes('\n')
			x := line[0]
			var cnt int
			readInt(line, 2, &cnt)
			S[i] = Pair{x, cnt}
		}
		fmt.Println(solve(n, S))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

type Pair struct {
	first  byte
	second int
}

// S content from 1...n
func solve(n int, S []Pair) int64 {
	S[0] = Pair{'#', 0}
	S[n+1] = Pair{'@', 1}

	p := make([]int, n+2)
	var center int
	var mx int
	for i := 1; i <= n; i++ {
		p[i] = 1
		if mx > i {
			p[i] = min(p[2*center-i], mx-i)
		}

		for S[i+p[i]] == S[i-p[i]] {
			p[i]++
		}
		if i+p[i] > mx {
			center = i
			mx = i + p[i]
		}
	}

	sum := make([]int64, n+2)
	for i := 1; i <= n; i++ {
		sum[i] = int64(S[i].second) + sum[i-1]
	}
	var res int64

	for i := 1; i <= n; i++ {
		cnt := int64(S[i].second)
		res += cnt * (cnt + 1) / 2
		if p[i] > 1 {
			res -= (cnt + 1) / 2
			u := i - p[i] + 1
			v := i + p[i] - 1
			res += (sum[v] - sum[u-1] + 1) / 2
			if S[u-1].first == S[v+1].first {
				res += int64(min(S[u-1].second, S[v+1].second))
			}
		} else {
			if S[i-1].first == S[i+1].first {
				res += int64(min(S[i-1].second, S[i+1].second))
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
