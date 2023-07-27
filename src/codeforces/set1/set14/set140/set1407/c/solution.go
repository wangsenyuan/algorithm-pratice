package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(i, j int) int {
		fmt.Printf("? %d %d\n", i, j)
		return readNum(reader)
	}

	n := readNum(reader)

	res := solve(n, ask)
	var buf bytes.Buffer
	buf.WriteString("!")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}
func solve(n int, ask func(int, int) int) []int {
	// ask(i, j) = p[i] % p[j]
	na := make([]int, n)
	P := make([]int, n+1)
	for i := 0; i < n; i++ {
		na[i] = i + 1
	}
	sum := (n + 1) * n / 2

	for len(na) > 1 {
		var keep []int
		for i := 0; i+1 < len(na); i += 2 {
			a, b := na[i], na[i+1]
			x := ask(a, b)
			y := ask(b, a)
			if x > y {
				P[a] = x
				sum -= x
				keep = append(keep, b)
			} else {
				P[b] = y
				sum -= y
				keep = append(keep, a)
			}
		}
		if len(na)&1 == 1 {
			keep = append(keep, na[len(na)-1])
		}
		na = keep
	}

	P[na[0]] = sum
	return P[1:]
}
