package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		res := solve(line[0], line[1], line[2], line[3])
		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
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

func solve(r, g, b, w int) bool {
	cnt := r&1 + g&1 + b&1 + w&1
	if cnt <= 1 {
		return true
	}
	if r > 0 && g > 0 && b > 0 {
		r--
		g--
		b--
		w++
		cnt = r&1 + g&1 + b&1 + w&1
		return cnt <= 1
	}
	return false
}

func solve1(r, g, b, w int) bool {
	cnt := r + g + b + w
	if cnt&1 == 1 {
		// make sure only one of the colors is odd, others even
		// move x to w, if it is even, then no change to w's parity, else odd, it will change w's parity
		p := r&1 + g&1 + b&1
		if p == 0 {
			// all even, w is odd
			return true
		}
		if p == 3 {
			// all odd, w is even, choose some odd x to make w odd
			return min3(r, g, b) > 0
		}
		// p == 1 or p == 2
		if p == 1 {
			// one odd,  have to use some even x, else will create another two odd colors
			return w&1 == 0
		}
		// p == 2
		// two odd, then have to use some odd x, and w need to be odd too
		return w&1 == 1 && min3(r, g, b) > 0
	}
	// make sure all colors even
	r &= 1
	g &= 1
	b &= 1

	if r+g+b > 0 && r+g+b < 3 {
		return false
	}
	if r+g+b == 0 {
		return w&1 == 0
	}
	return w&1 == 1
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
