package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		B := readNNums(reader, 2*n)
		A := solve(n, B)
		if len(A) == 0 {
			buf.WriteString("-1")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", A[i]))
			}
		}
		buf.WriteString(fmt.Sprintln(""))
	}
	fmt.Print(buf.String())
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

func solve(n int, B []int) []int {
	cnt := make(map[int]int)
	for _, num := range B {
		cnt[num]++
	}
	arr := make([][]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = make([]int, 0, n)
	}
	for k, v := range cnt {
		if v >= 5 {
			return nil
		}
		arr[v] = append(arr[v], k)
	}

	for i := 0; i < 5; i++ {
		sort.Ints(arr[i])
		reverse(arr[i])
	}

	A := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		if 2*i+1 <= n {
			// place at middle in prefix
			// l = i
			// r = i
			// 2 * i + 1 <= n
			tmp++
		}
		if 2*i+2 <= n {
			// place at left-median in prefix
			// l = i + 1
			// r = i + 1
			// l + r <= n
			tmp++
		}
		// place at middle at suffix
		// r = (n - 1 - i)
		// l = r
		if 2*(n-1-i)+1 <= n {
			tmp++
		}
		// place at left-median at suffix
		// l = r - 1
		if n-1-i > 0 && n-1-i+1+n-2-i <= n {
			tmp++
		}
		if len(arr[tmp]) == 0 {
			return nil
		}
		l := len(arr[tmp])
		A[i] = arr[tmp][l-1]
		arr[tmp] = arr[tmp][:l-1]
	}

	if !sort.IntsAreSorted(A) {
		return nil
	}

	return A
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
