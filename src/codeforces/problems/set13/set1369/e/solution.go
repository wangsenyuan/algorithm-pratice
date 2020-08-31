package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	W := readNNums(reader, n)
	F := make([][]int, m)
	for i := 0; i < m; i++ {
		F[i] = readNNums(reader, 2)
	}
	res := solve(W, F)
	if len(res) == 0 {
		fmt.Println("DEAD")
		return
	}
	fmt.Println("ALIVE")
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func solve(W []int, friends [][]int) []int {
	n := len(W)
	m := len(friends)

	freq := make([]int, n)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, 3)
	}
	for i, f := range friends {
		f[0]--
		f[1]--
		x, y := f[0], f[1]
		freq[x]++
		freq[y]++
		adj[x] = append(adj[x], i)
		adj[y] = append(adj[y], i)
	}
	que := make([]int, n)
	vis := make([]bool, n)
	var end int
	for i := 0; i < n; i++ {
		if freq[i] <= W[i] {
			vis[i] = true
			que[end] = i
			end++
		}
	}

	ok := make([]bool, m)
	res := make([]int, 0, m)
	var front int

	for front < end {
		x := que[front]
		front++
		for _, i := range adj[x] {
			if !ok[i] {
				ok[i] = true
				res = append(res, i+1)
				f := friends[i]
				y := f[0] ^ f[1] ^ x
				freq[y]--
				if freq[y] <= W[y] && !vis[y] {
					que[end] = y
					end++
					vis[y] = true
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			return nil
		}
	}
	reverse(res)
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
