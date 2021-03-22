package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--

		s, _ := reader.ReadBytes('\n')
		var n int
		for n < len(s) && s[n] >= '0' && s[n] <= '9' {
			n++
		}
		res := solve(n, s)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(n int, S []byte) int {
	var mx int
	for i := 0; i < n; i++ {
		mx = max(mx, int(S[i]-'0'))
	}
	k := mx + 1
	vis := make([]bool, n*k)
	que := make([]int, n*k)
	dist := make([]int, n*k)
	for i := 0; i < len(dist); i++ {
		dist[i] = n*k + 1
	}

	var front, end int
	for i := 0; i < n; i++ {
		u := i*k + int(S[i]-'0')
		que[end] = u
		end++
		vis[u] = true
		dist[u] = 0
	}

	cp := make([][]int, n)
	cf := make([]int, n)
	cb := make([]int, n)

	for i := 0; i < n; i++ {
		cp[i] = make([]int, k)
		cp[i][cb[i]] = int(S[i] - '0')
		cb[i]++
	}

	for front < end {
		cur := que[front]
		front++
		pos := cur / k
		val := cur % k
		for i := pos - val; i <= pos+val; i++ {
			if i < 0 || i >= n {
				continue
			}
			to := i*k + val
			if vis[to] {
				continue
			}
			d := abs(i - pos)
			for cf[i]+1 < cb[i] && dist[i*k+cp[i][cf[i]+1]] <= dist[cur] {
				cf[i]++
			}
			if cp[i][cf[i]] <= val-d {
				vis[to] = true
				dist[to] = dist[cur] + 1
				que[end] = to
				end++
				cp[i][cb[i]] = val
				cb[i]++
			}
		}
	}

	var ret int

	for i := 0; i < n; i++ {
		ret = max(ret, dist[i*k+mx])
	}

	return ret
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
