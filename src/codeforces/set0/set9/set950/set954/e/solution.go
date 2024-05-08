package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, T := readTwoNums(reader)
	v := readNNums(reader, n)
	x := readNNums(reader, n)

	res := solve(x, v, T)

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

const oo = 1 << 50
const eps = 1e-7

func solve(x []int, v []int, T int) float64 {
	n := len(x)
	cold := make([][]int, 0, n)
	hot := make([][]int, 0, n)
	sum := make([]int, 2)
	var res int
	for i := 0; i < n; i++ {
		t := x[i] - T
		if t < 0 {
			cold = append(cold, []int{-t, v[i]})
			sum[0] += -t * v[i]
		} else if t > 0 {
			hot = append(hot, []int{t, v[i]})
			sum[1] += t * v[i]
		} else {
			res += v[i]
		}
	}

	if len(cold) == 0 || len(hot) == 0 {
		// 无法转换
		return float64(res)
	}
	slices.SortFunc(cold, func(a, b []int) int {
		return a[0] - b[0]
	})
	slices.SortFunc(hot, func(a, b []int) int {
		return a[0] - b[0]
	})

	if sum[0] > sum[1] {
		sum[0], sum[1] = sum[1], sum[0]
		cold, hot = hot, cold
	}
	for _, cur := range cold {
		res += cur[1]
	}
	// 温差越大的，最后使用，所使用的量就要越多
	var ans float64
	for _, cur := range hot {
		tmp := cur[0] * cur[1]
		if tmp < sum[0] {
			sum[0] -= tmp
			res += cur[1]
		} else {
			ans = float64(res) + float64(sum[0])/float64(cur[0])
			break
		}
	}
	return ans
}
