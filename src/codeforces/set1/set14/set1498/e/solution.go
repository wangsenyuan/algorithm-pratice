package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	K := readNNums(reader, n)

	ask := func(a, b int) string {
		a++
		b++
		fmt.Printf("? %d %d\n", a, b)
		res, _ := reader.ReadString('\n')
		return res[:len(res)-1]
	}

	res := solve(n, K, ask)

	fmt.Printf("! %d %d\n", res[0]+1, res[1]+1)

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

func solve(n int, K []int, ask func(int, int) string) []int {
	// if k[a] < k[b] then b is reachable from a any way
	// so we just need to check whether a is reachable from b
	W := make([]Triple, 0, n*n/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if K[i] > K[j] {
				W = append(W, Triple{K[i] - K[j], i, j})
			} else {
				W = append(W, Triple{K[j] - K[i], j, i})
			}
		}
	}

	sort.Slice(W, func(i, j int) bool {
		return W[i].first > W[j].first
	})

	for _, w := range W {
		i, j := w.second, w.third
		// i is reachable from j, check i -> j ?
		res := ask(i, j)
		if res == "Yes" {
			return []int{i, j}
		}
	}

	return []int{-1, -1}
}

type Pair struct {
	first, second int
}

type Triple struct {
	first, second, third int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
