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

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		ans := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

const MAX_K = 4000010
const MAX_S = 145

var spf [MAX_K]int
var ans [MAX_K]int64

var res []int
var bak []int

func init() {
	for x := 2; x < MAX_K; x++ {
		if spf[x] == 0 {
			for y := x; y < MAX_K; y += x {
				if spf[y] == 0 {
					spf[y] = x
				}
			}
		}
	}

	res = make([]int, MAX_S)
	bak = make([]int, MAX_S)

	count := make([]int, MAX_S)
	for k := 1; 4*k+1 < MAX_K; k++ {
		fact := factors(4*k + 1)
		var sum int64
		for i := len(fact) - 1; i >= 0; i-- {
			count[i] = ((4*k+1)/fact[i] + 1) / 2
			for j := i + 1; j < len(fact); j++ {
				if fact[j]%fact[i] == 0 {
					count[i] -= count[j]
				}
			}
			sum += int64(fact[i]) * int64(count[i])
		}
		ans[k] = sum - 1
	}
}

func factors(x int) []int {
	res[0] = 1
	pos := 1
	for x > 1 {
		p := spf[x]
		var cnt int
		for x%p == 0 {
			x /= p
			cnt++
		}

		copy(bak, res[:pos])
		for pw, cur := 1, p; pw <= cnt; pw, cur = pw+1, cur*p {
			for i := 0; i < pos; i++ {
				bak[pw*pos+i] = res[i] * cur
			}
		}
		pos = pos * (cnt + 1)

		copy(res, bak[:pos])
	}

	sort.Ints(res[:pos])

	return res[:pos]
}

func solve(n int) int64 {
	return ans[n]
}
