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

	n, w, h := readThreeNums(reader)
	envs := make([][]int, n)
	for i := 0; i < n; i++ {
		envs[i] = readNNums(reader, 2)
	}
	res := solve(n, w, h, envs)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func solve(n, w, h int, envs [][]int) []int {
	items := make([]Envelope, n)
	var H int
	for i := 0; i < n; i++ {
		items[i] = Envelope{envs[i][0], envs[i][1], i}
		H = max(H, envs[i][1])
	}

	sort.Sort(Envelopes(items))
	H++
	arr := make([]int, 2*H)

	update := func(p int, v int) {
		p += H
		arr[p] = v
		for p > 0 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		var res int
		l += H
		r += H

		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	dp := make([]int, n+1)
	var best int
	for i := 0; i < n; i++ {
		cur := items[i]
		if cur.width <= w {
			break
		}

		if cur.height <= h {
			continue
		}

		cnt := get(cur.height+1, H)
		dp[i] = cnt + 1
		update(cur.height, cnt+1)
		best = max(best, cnt+1)
	}

	if best == 0 {
		return nil
	}

	cur := -1
	for i := n - 1; i >= 0; i-- {
		if dp[i] == best {
			cur = i
			break
		}
	}
	res := make([]int, 0, best)
	res = append(res, items[cur].index+1)

	for i := cur - 1; i >= 0; i-- {
		if dp[i]+1 == best && items[i].width > items[cur].width && items[i].height > items[cur].height {
			best--
			cur = i
			res = append(res, items[cur].index+1)
		}
	}
	//reverse(res)

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Envelope struct {
	width  int
	height int
	index  int
}

type Envelopes []Envelope

func (ps Envelopes) Len() int {
	return len(ps)
}

func (ps Envelopes) Less(i, j int) bool {
	return ps[i].width > ps[j].width || ps[i].width == ps[j].width && ps[i].height < ps[j].height
}

func (ps Envelopes) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
