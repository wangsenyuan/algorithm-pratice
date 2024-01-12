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

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)

		ask := func(rem []int) []int {
			var buf bytes.Buffer
			buf.WriteString(fmt.Sprintf("- %d", len(rem)))
			for i := 0; i < len(rem); i++ {
				buf.WriteString(fmt.Sprintf(" %d", rem[i]))
			}
			buf.WriteByte('\n')
			fmt.Print(buf.String())
			n -= len(rem)
			return readNNums(reader, n)
		}

		res := solve(n, a, ask)

		fmt.Printf("! %d\n", res)
	}

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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, a []int, fn func([]int) []int) int {
	//a := fn(nil)
	freq := getFreq(a)
	var newFreq []int
outer:
	for {
		a = fn(nil)
		newFreq = getFreq(a)
		for j := 1; j <= 9; j++ {
			if newFreq[j] != freq[j] {
				break outer
			}
		}
	}
	var mimic int
	// 最多两次
	var rem []int
	var arr []int
	for i := 0; i < n; i++ {
		x := a[i]
		if freq[x] >= newFreq[x] {
			rem = append(rem, i+1)
		} else {
			mimic = x
			arr = append(arr, i+1)
		}
	}

	if len(arr) == 1 {
		return arr[0]
	}

	// 剩下的就是 newfreq[x] >
	a = fn(rem)
	for i, x := range a {
		if x != mimic {
			return i + 1
		}
	}

	for {
		a = fn(nil)
		for i, x := range a {
			if x != mimic {
				return i + 1
			}
		}
	}

	return -1
}

func getFreq(a []int) []int {
	freq := make([]int, 10)
	for _, x := range a {
		freq[x]++
	}
	return freq
}
