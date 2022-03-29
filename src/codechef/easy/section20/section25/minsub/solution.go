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
		n := readNum(reader)
		A := readString(reader)
		S, T := solve(A[:n])
		buf.WriteString(S)
		buf.WriteByte('\n')
		buf.WriteString(T)
		buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(A string) (string, string) {
	cnt := make([]int, 2)

	n := len(A)

	for i := 0; i < n; i++ {
		cnt[int(A[i]-'0')]++
	}
	var flipped bool
	var S, T string
	if cnt[0] < cnt[1] {
		flipped = true
		S, T = process(flip(A), []int{cnt[1], cnt[0]})
	} else {
		S, T = process(A, cnt)
	}

	if flipped {
		S = flip(S)
		T = flip(T)
	}

	return S, T
}

func flip(s string) string {
	buf := []byte(s)
	for i := 0; i < len(s); i++ {
		if buf[i] == '0' {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}
	return string(buf)
}

func process(A string, cnt []int) (string, string) {
	// cnt[0] >= cnt[1]
	n := len(A)
	S := make([]byte, n)
	T := make([]byte, n)
	// seperate into cnt[1] + 1 parts
	if cnt[0] == cnt[1] {
		for i := 0; i < n; i++ {
			S[i] = '1'
			if i%2 == 1 {
				S[i] = '0'
			}
			T[i] = '0'
			if i*2 >= n {
				T[i] = '1'
			}
		}
	} else {
		x, r := cnt[0]/(cnt[1]+1), cnt[0]%(cnt[1]+1)
		for i := n - 1; i >= 0; {
			// every x or (x + 1) 0, then 1
			y := x
			if r > 0 {
				y++
				r--
			}
			for y > 0 {
				S[i] = '0'
				y--
				i--
			}
			if i >= 0 {
				S[i] = '1'
				i--
			}
		}
		for i := 0; i < cnt[0]; i++ {
			T[i] = '0'
		}
		for i := cnt[0]; i < n; i++ {
			T[i] = '1'
		}
	}

	return string(S), string(T)
}
