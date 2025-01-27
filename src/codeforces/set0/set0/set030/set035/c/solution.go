package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)

	n, m := readTwoNums(reader)
	k := readNum(reader)
	fires := readNNums(reader, 2*k)
	res := solve(n, m, fires)
	fmt.Fprintln(w, res[0], res[1])
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(n int, m int, fires []int) []int {
	que := make([]int, n*m)
	var head, tail int
	marked := make([]bool, n*m)
	for i := 0; i < len(fires); i += 2 {
		r, c := fires[i], fires[i+1]
		r--
		c--
		que[head] = r*m + c
		head++
		marked[r*m+c] = true
	}

	for tail < head {
		r := que[tail] / m
		c := que[tail] % m
		tail++
		for i := 0; i < 4; i++ {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && !marked[x*m+y] {
				marked[x*m+y] = true
				que[head] = x*m + y
				head++
			}
		}
	}
	return []int{que[tail-1]/m + 1, que[tail-1]%m + 1}
}
