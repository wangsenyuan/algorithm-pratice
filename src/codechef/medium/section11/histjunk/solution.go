package main

import (
	"bufio"
	"bytes"
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
		n, m := readTwoNums(reader)
		var cnt int
		var res [][]int
		if n >= 4 {
			for i := 0; i < m; i++ {
				reader.ReadBytes('\n')
			}
			cnt, res = solve1(n, m)
		} else if n < 3 || n == 3 && m == 3 {
			// complete graph
			for i := 0; i < m; i++ {
				reader.ReadBytes('\n')
			}
			cnt, res = solve2(n, m)
		} else {
			E := make([][]int, m)
			for i := 0; i < m; i++ {
				E[i] = readNNums(reader, 2)
			}
			cnt, res = solve3(n, m, E)
		}
		fmt.Printf("%d %d\n", cnt, len(res))
		var buf bytes.Buffer
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
		}
		fmt.Print(buf.String())
		// fmt.Println("======")
	}
}

func solve1(n, m int) (int, [][]int) {
	// n >= 4
	a, b, c, d := n+1, n+2, n+3, n+4

	res := make([][]int, 0, 2*n+2)
	res = append(res, []int{a, b})
	res = append(res, []int{c, d})

	for i := 1; i <= n; i++ {
		res = append(res, []int{a, i})
		res = append(res, []int{c, i})
	}

	return 4, res
}

func solve2(n, m int) (int, [][]int) {
	// n < 3 || n == 3 && m == 3
	return 0, nil
}

func solve3(n, m int, E [][]int) (int, [][]int) {
	// n == 3 && m == 2
	cnt := make([]int, 4)

	for _, e := range E {
		cnt[e[0]]++
		cnt[e[1]]++
	}

	X := []int{0, 0, 0}

	for i := 1; i <= 3; i++ {
		if cnt[i] == 2 {
			X[1] = i
		} else if X[0] == 0 {
			X[0] = i
		} else {
			X[2] = i
		}
	}
	a, b := 4, 5
	res := make([][]int, 0, 3)
	res = append(res, []int{a, X[0]})
	res = append(res, []int{a, X[2]})
	res = append(res, []int{b, X[1]})
	return 2, res
}
