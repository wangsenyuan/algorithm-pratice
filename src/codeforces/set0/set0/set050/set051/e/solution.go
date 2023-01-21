package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)

	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, E)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, E [][]int) int64 {
	a := make([][]int64, n)
	a2 := make([][]int64, n)
	a3 := make([][]int64, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int64, n)
		a2[i] = make([]int64, n)
		a3[i] = make([]int64, n)
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		a[u][v]++
		a[v][u]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k <= i; k++ {
				a2[i][k] += a[i][j] * a[j][k]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a2[i][j] = a2[j][i]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k <= i; k++ {
				a3[i][k] += a2[i][j] * a[j][k]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a3[i][j] = a3[j][i]
		}
	}

	var ans int64

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += a2[i][j] * a3[j][i]
		}
	}

	for i := 0; i < n; i++ {
		ans -= 5 * a2[i][i] * a3[i][i]
	}

	for i := 0; i < n; i++ {
		ans += 5 * a3[i][i]
	}

	return ans / 10
}
