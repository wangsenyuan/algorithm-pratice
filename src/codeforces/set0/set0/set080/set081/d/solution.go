package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, res := process(reader)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (a []int, res []int) {
	n, m := readTwoNums(reader)
	a = readNNums(reader, m)
	res = solve(n, a)
	return
}

func solve(n int, a []int) []int {
	m := len(a)
	if m == 1 {
		return nil
	}

	type pair struct {
		first  int
		second int
	}
	arr := make([]pair, m)
	for i := range m {
		arr[i] = pair{a[i], i + 1}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return y.first - x.first
	})
	res := make([]int, n)

	for i, j, cnt := 0, 0, 0; cnt < n; {
		if j == m {
			return nil
		}
		old := cnt
		cur := arr[j]
		j++
		for cnt < n && cur.first > 0 {
			if res[i] != 0 {
				i = (i + 1) % n
			}
			if res[(i-1+n)%n] == cur.second || res[(i+1)%n] == cur.second {
				break
			}
			res[i] = cur.second
			i = (i + 2) % n
			cnt++
			cur.first--
		}
		if cnt == old {
			return nil
		}

	}
	return res
}
