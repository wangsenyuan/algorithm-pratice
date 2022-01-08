package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		q := readNum(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, A, Q)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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

func solve(n int, A []int, Q [][]int) []int {
	pairs := make(map[Pair]int)

	for _, cur := range Q {
		p := Pair{cur[0], cur[1]}
		pairs[p]++
	}

	rvals := make(map[int][]int)
	lvals := make(map[int][]int)

	add := func(mem map[int][]int, k, v int) {
		if _, found := mem[k]; !found {
			mem[k] = make([]int, 0, 2)
		}
		mem[k] = append(mem[k], v)
	}

	for p := range pairs {
		l, r := p.first, p.second
		add(rvals, r, l)
		add(lvals, l, r)
	}

	fp := make(map[int]int)

	BLOCK := int(math.Sqrt(float64(len(Q))))

	big := make(map[int]bool)

	for r, v := range rvals {
		if len(v) > BLOCK {
			big[r] = true
			fp[r] = 1
		}
	}
	bigVals := make(map[int][]int)
	for l, v := range lvals {
		for _, r := range v {
			if big[r] {
				add(bigVals, l, r)
			}
		}
	}

	recent := make(map[int]int)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	prev := make([]int, n)

	prevIndForVal := make(map[int]int)

	for i := 0; i < n; i++ {
		prev[i] = -1
		x := A[i]
		if big[x] {
			dp[i] = fp[x]
			if y, found := prevIndForVal[x]; found {
				prev[i] = y
			}
		} else {
			for _, l := range rvals[x] {
				if li, ok := recent[l]; ok && dp[i] < dp[li]+1 {
					dp[i] = dp[li] + 1
					prev[i] = li
				}
			}
		}

		recent[x] = i

		for _, r := range bigVals[x] {
			if fp[r] < dp[i]+1 {
				fp[r] = dp[i] + 1
				prevIndForVal[r] = i
			}
		}
	}

	var pos int
	var best int
	for i := 0; i < n; i++ {
		if best < dp[i] {
			best = dp[i]
			pos = i
		}
	}

	var arr []int

	for pos >= 0 {
		arr = append(arr, pos+1)
		pos = prev[pos]
	}

	reverse(arr)

	return arr
}

type Pair struct {
	first, second int
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
