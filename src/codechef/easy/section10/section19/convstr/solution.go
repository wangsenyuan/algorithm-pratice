package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A, _ := reader.ReadString('\n')
		B, _ := reader.ReadString('\n')
		ok, res := solve(n, A, B)
		if !ok {
			fmt.Println("-1")
		} else {
			fmt.Println(len(res))
			for i := 0; i < len(res); i++ {
				fmt.Printf("%d ", len(res[i]))
				for j := 0; j < len(res[i]); j++ {
					fmt.Printf("%d ", res[i][j])
				}
				fmt.Println()
			}
		}
	}
}

func solve(n int, A string, B string) (bool, [][]int) {
	for i := 0; i < n; i++ {
		if A[i] < B[i] {
			return false, nil
		}
	}

	cnt := make([]int, 26)

	for i := 0; i < n; i++ {
		cnt[int(A[i]-'a')]++
	}

	for i := 0; i < n; i++ {
		if cnt[int(B[i]-'a')] == 0 {
			return false, nil
		}
	}

	res := make([][]int, 0, 26)

	for c := 25; c >= 0; c-- {
		if cnt[c] == 0 {
			continue
		}

		tmp := make([]int, 0, cnt[c])

		for i := 0; i < n; i++ {
			if int(B[i]-'a') == c {
				if A[i] != B[i] {
					tmp = append(tmp, i)
				}
			}
		}
		if len(tmp) == 0 {
			continue
		}
		for i := 0; i < n; i++ {
			if int(A[i]-'a') == c {
				tmp = append(tmp, i)
				break
			}
		}

		res = append(res, tmp)
	}

	return true, res
}
