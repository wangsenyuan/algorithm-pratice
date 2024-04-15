package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	x := readNNums(reader, n/2)

	res := solve(x)

	if len(res) == 0 {
		fmt.Println("No")
	} else {
		var buf bytes.Buffer
		buf.WriteString("Yes\n")
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

type Pair struct {
	first  int
	second int
}

const N = 200100

var V [N][]Pair

func init() {
	for i := 0; i < N; i++ {
		V[i] = make([]Pair, 0, 1)
	}

	for i := 1; i < N; i++ {
		if 2*i+1 >= N {
			break
		}
		for j := i + 1; j*j-i*i < N; j++ {
			V[j*j-i*i] = append(V[j*j-i*i], Pair{i, j})
		}
	}
}

func solve(x []int) []int {
	n := len(x) * 2
	sq := make([]int, n+1)

	for i := 2; i <= n; i += 2 {
		num := x[i/2-1]
		for _, t := range V[num] {
			if sq[i-2] < t.first {
				sq[i-1] = t.first
				sq[i] = t.second
				break
			}
		}
		if sq[i] == 0 {
			return nil
		}
	}

	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = sq[i]*sq[i] - sq[i-1]*sq[i-1]
	}

	return res[1:]
}
