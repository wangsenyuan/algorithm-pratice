package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		H := readNNums(reader, n)
		res := solve(n, k, H)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const INF = 1 << 30

func solve(n int, k int, H []int) int {
	que := make([]Pair, n+1)
	var front, end int

	que[end] = Pair{0, -1}
	end++

	var sum int
	var i int

	for i < k {
		// if took ith homework
		for front < end && que[end-1].first <= -H[i] {
			end--
		}
		sum += H[i]
		que[end] = Pair{-H[i], i}
		end++
		i++
	}

	for i < n {
		// if took ith homework
		for front < end && que[front].second+k < i-1 {
			front++
		}
		cur := sum + que[front].first
		sum += H[i]
		for end > front && que[end-1].first <= cur-sum {
			end--
		}
		que[end] = Pair{cur - sum, i}
		end++
		i++
	}

	for front < end && que[front].second+k < n-1 {
		front++
	}

	return - que[front].first
}

type Pair struct {
	first, second int
}
