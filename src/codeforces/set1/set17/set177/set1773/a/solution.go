package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		ok, p, q := solve(a)
		if !ok {
			buf.WriteString("Impossible\n")
			continue
		}
		buf.WriteString("Possible\n")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", p[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", q[i]))
		}
		buf.WriteByte('\n')
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

	p := make([]int, n)
	q := make([]int, n)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		a[i]--
		q[i] = i
		pos[a[i]] = i
	}
	cnt := 100
	for cnt > 0 {
		rand.Shuffle(n, func(i, j int) {
			q[i], q[j] = q[j], q[i]
		})
		ok := true
		// a[p[q[i]] = i
		for i := n - 1; i >= 0; i-- {
			if q[i] == i {
				if i == 0 {
					ok = false
					break
				}
				k := rand.Intn(i)
				q[i], q[k] = q[k], q[i]
			}
			j := q[i]
			k := pos[i]
			// a[k] = i
			// p[j] = k
			p[j] = k
			if j == k {
				ok = false
				break
			}
		}

		if ok {
			for i := 0; i < n; i++ {
				p[i]++
				q[i]++
			}
			return true, p, q
		}
		cnt--
	}

	return false, nil, nil
}
