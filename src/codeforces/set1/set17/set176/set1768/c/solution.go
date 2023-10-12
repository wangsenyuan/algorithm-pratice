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
		n := readNum(reader)
		a := readNNums(reader, n)
		ok, x, y := solve(a)
		if !ok {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", x[i]))
			}
			buf.WriteByte('\n')
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", y[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int) (bool, []int, []int) {
	n := len(a)

	b := make([][]int, n+1)

	for i, num := range a {
		if len(b[num]) == 0 {
			b[num] = make([]int, 0, 1)
		}
		b[num] = append(b[num], i+1)
		if len(b[num]) >= 3 {
			return false, nil, nil
		}
	}

	x := make([]int, n+1)
	y := make([]int, n+1)
	px := make([]int, n+1)
	py := make([]int, n+1)
	for i := 0; i <= n; i++ {
		x[i] = -1
		y[i] = -1
		px[i] = -1
		py[i] = -1
	}

	for num := n; num > 0; num-- {
		for _, i := range b[num] {
			if px[num] == -1 {
				x[i] = num
				px[num] = i
			} else if py[num] == -1 {
				y[i] = num
				py[num] = i
			}
		}
	}

	for num, vp, vq := n, n, n; num > 0; num-- {
		for _, i := range b[num] {
			for px[vp] != -1 {
				vp--
			}
			for py[vq] != -1 {
				vq--
			}
			if x[i] < 0 && vp > 0 {
				x[i] = vp
				px[vp] = i
			}
			if y[i] < 0 && vq > 0 {
				y[i] = vq
				py[vq] = i
			}
		}
	}

	for i := 1; i <= n; i++ {
		if max(x[i], y[i]) != a[i-1] {
			return false, nil, nil
		}
	}
	return true, x[1:], y[1:]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
