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
		A := readNNums(reader, n)
		ok, res := solve(n, A)

		if !ok {
			buf.WriteString("-1\n")
			continue
		}
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, tr := range res {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", tr[0], tr[1], tr[2]))
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

func solve(n int, A []int) (bool, [][]int) {
	if n == 1 {
		return true, nil
	}
	if n == 2 {
		return A[0]^A[1] > 0, nil
	}

	var sum int64
	same := true
	for i := 0; i < n; i++ {
		sum += int64(A[i])
		if A[i] != A[0] {
			same = false
		}
	}

	if sum == 0 {
		return false, nil
	}
	var res [][]int

	if n == 3 {
		if A[0] == A[2] {
			if A[0]^A[1] > 0 {
				return true, nil
			}
			// all same
			return true, [][]int{{1, 3, 2}}
		}
		if A[1] == 0 {
			if A[0] > 0 {
				res = append(res, []int{1, 2, 3})
			} else {
				res = append(res, []int{2, 3, 1})
			}
			return true, res
		}
		return false, nil
	}

	if same {
		for i := 0; i < n; i += 2 {
			res = append(res, []int{1, 3, i + 1})
		}
		return true, res
	}

	evens := make(map[int]int)
	odds := make(map[int]int)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			evens[A[i]] = i + 1
		} else {
			odds[A[i]] = i + 1
		}
	}

	if len(evens) >= 2 {
		first, second := getAnyTwoValue(evens)
		for i := 1; i <= n; i += 2 {
			res = append(res, []int{first, second, i})
		}
		first, second = 1, 3
		for i := 2; i <= n; i += 2 {
			res = append(res, []int{first, second, i})
		}
	} else if len(odds) >= 2 {
		first, second := getAnyTwoValue(odds)
		for i := 2; i <= n; i += 2 {
			res = append(res, []int{first, second, i})
		}
		first, second = 2, 4
		for i := 1; i <= n; i += 2 {
			res = append(res, []int{first, second, i})
		}
	}
	return true, res
}

func getAnyTwoValue(mem map[int]int) (int, int) {
	first, second := -1, -1
	for _, v := range mem {
		if first < 0 {
			first = v
		} else if second < 0 {
			second = v
			break
		}
	}
	return first, second
}
