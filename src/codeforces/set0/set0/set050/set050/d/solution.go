package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	k, eps := readTwoNums(reader)

	orig := readNNums(reader, 2)

	ps := make([][]int, n)

	for i := 0; i < n; i++ {
		ps[i] = readNNums(reader, 2)
	}

	res := solve(n, k, float64(eps), orig, ps)

	fmt.Printf("%.9f\n", res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n, k int, e float64, orig []int, pos [][]int) float64 {

	e /= 1000

	f := make([]float64, n+1)

	prop := func(i int, r float64) float64 {
		dx := float64(pos[i-1][0] - orig[0])
		dy := float64(pos[i-1][1] - orig[1])
		dis := dx*dx + dy*dy

		if dis <= r*r {
			return 1.0
		}
		return math.Exp(1.0 - dis/(r*r))
	}

	check := func(r float64) bool {
		for i := 1; i <= n; i++ {
			f[i] = 0
		}

		f[0] = 1.0

		for i := 1; i <= n; i++ {
			pr := prop(i, r)

			for j := i; j > 0; j-- {
				f[j] = f[j]*(1-pr) + pr*f[j-1]
			}
			f[0] = f[0] * (1 - pr)
		}

		var sum float64

		for i := 0; i < k; i++ {
			sum += f[i]
		}

		return sum <= e
	}

	var ans float64
	var x float64 = 1e4

	for i := 0; i < 50; i, x = i+1, x/2 {
		if !check(ans + x) {
			ans += x
		}
	}

	return ans
}
