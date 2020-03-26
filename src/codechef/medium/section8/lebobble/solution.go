package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	var tc int
	bs, _ := reader.ReadBytes('\n')

	readInt(bs, 0, &tc)

	for tc > 0 {
		tc--

		N, D := readTwoNums(reader)
		B := readNNums(reader, N)
		P := readNNums(reader, N)
		fmt.Printf("%.4f\n", solve(N, D, B, P))
	}
}

func solve(N, D int, B []int, P []int) float64 {
	nums := make([]int, 2*N)
	for i := 0; i < N; i++ {
		nums[2*i] = B[i]
		nums[2*i+1] = B[i] + D
	}
	sort.Ints(nums)

	ii := make(map[int]int)
	var j int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ii[nums[i]] = j
		j++
	}

	m := len(ii)

	arr := make([]int64, m+1)

	update := func(pos int, v int64) {
		pos++
		for pos <= m {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int64 {
		var res int64
		pos++
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}
		return res
	}

	query := func(x int) int64 {
		return get(m-1) - get(x)
	}

	var res int64

	for i := 0; i < N; i++ {
		x := query(ii[B[i]])
		res += x * int64(100-P[i])
		y := query(ii[B[i]+D])
		res += y * int64(P[i])

		update(ii[B[i]], 100-int64(P[i]))
		update(ii[B[i]+D], int64(P[i]))
	}

	return float64(res) / 10000
}

func solve1(N, D int, B []int, P []int) float64 {
	var res float64

	p := make([]float64, N)
	for i := 0; i < N; i++ {
		p[i] = float64(P[i]) / 100
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			x, y := B[i], B[j]

			if x > y {
				res += (1-p[i])*(1-p[j]) + p[i]*p[j]
			}

			if x+D > y {
				res += p[i] * (1 - p[j])
			}

			if x > y+D {
				res += (1 - p[i]) * p[j]
			}
		}
	}

	return res
}
