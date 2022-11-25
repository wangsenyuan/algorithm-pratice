package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(a, b int) int {
		fmt.Printf("? %d %d\n", a, b)
		return readNum(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--

		n := readNum(reader)

		res := solve(n, ask)

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

func solve(n int, ask func(int, int) int) int {
	arr := make([]int, 1<<n)

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	m := len(arr)

	for m >= 4 {
		var j int
		for i := 0; i < m; i += 4 {
			a, b, c, d := arr[i], arr[i+1], arr[i+2], arr[i+3]
			x := ask(a, c)
			if x == 1 {
				// a win, b & c won't win
				x = ask(a, d)
				if x == 1 {
					arr[j] = a
				} else {
					// x == 2
					arr[j] = d
				}
			} else if x == 2 {
				// c win, then a & d won't win
				x = ask(b, c)
				if x == 1 {
					arr[j] = b
				} else {
					arr[j] = c
				}
			} else {
				// x == 0,
				x = ask(b, d)
				if x == 1 {
					arr[j] = b
				} else {
					arr[j] = d
				}
			}

			j++
		}
		m = j
	}
	if m == 2 {
		x := ask(arr[0], arr[1])
		if x == 2 {
			arr[0] = arr[1]
		}
	}
	return arr[0]
}
