package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		A := readNNums(reader, 3)
		B := readNNums(reader, 3)
		fmt.Println(solve(A, B))
	}
}

func solve(A, B []int) int {

	var best = 3
	var loop func(a, b []int64, cnt int)

	loop = func(a, b []int64, cnt int) {
		if cnt >= best {
			return
		}
		if sameArr(a, b) {
			best = cnt
			return
		}

		if cnt >= 2 {
			return
		}
		add := make(map[int64]bool)
		for i := 0; i < len(a); i++ {
			add[b[i]-a[i]] = true
		}
		mult := make(map[int64]bool)

		for i := 0; i < len(a); i++ {
			if a[i] != 0 && b[i]%a[i] == 0 {
				mult[b[i]/a[i]] = true
			}
		}

		mult[mulFac(a[0], a[1], b[0], b[1])] = true
		mult[mulFac(a[2], a[1], b[2], b[1])] = true
		mult[mulFac(a[0], a[2], b[0], b[2])] = true
		mult[0] = true

		for mask := 0; mask < 8; mask++ {
			all := make([]int, 0, 3)
			for i := 0; i < 3; i++ {
				if mask&(1<<uint(i)) > 0 {
					all = append(all, i)
				}
			}
			aa := make([]int64, 3)
			for x := range add {
				for i := 0; i < 3; i++ {
					aa[i] = a[i]
				}
				for _, j := range all {
					aa[j] += x
				}
				loop(aa, b, cnt+1)
			}

			for x := range mult {
				for i := 0; i < 3; i++ {
					aa[i] = a[i]
				}
				for _, j := range all {
					aa[j] *= x
				}
				loop(aa, b, cnt+1)
			}
		}
	}
	x := make([]int64, 3)
	y := make([]int64, 3)
	for i := 0; i < 3; i++ {
		x[i] = int64(A[i])
		y[i] = int64(B[i])
	}

	loop(x, y, 0)

	return best
}

func mulFac(a, b, c, d int64) int64 {
	if b != a && (d-c)%(b-a) == 0 {
		return (d - c) / (b - a)
	}
	return 1
}

func sameArr(a, b []int64) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
