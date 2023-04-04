package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		ok, res, splits := solve(A)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, op := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
			}
			buf.WriteString(fmt.Sprintf("%d\n", len(splits)))
			for _, d := range splits {
				buf.WriteString(fmt.Sprintf("%d ", d))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(B []int) (bool, [][]int, []int) {
	// B[i] = B[i+k]
	// k = 最后的length  n / 2
	// 另外对于 A[i], A[i+?]相等
	// 因为每次都是增加两个相同的数字，所以一开始的数字的freq不会减少（且成对增加）
	// 所以要求一开始 A[i]的 freq % 2 = 0
	n := len(B)

	cnt := make(map[int]int)

	for _, num := range B {
		cnt[num]++
	}

	for _, v := range cnt {
		if v&1 == 1 {
			return false, nil, nil
		}
	}
	// always have solution
	var ops [][]int
	var splits []int

	var mdf int

	reverse_pref := func(ln int) {
		for i := 0; i < ln; i++ {
			ops = append(ops, []int{mdf + i + ln, B[i]})
		}
		if ln != 0 {
			splits = append(splits, ln*2)
		}
		mdf += ln * 2
		reverse(B[:ln])
	}

	move_to_front := func(id int) {
		reverse_pref(id)
		reverse_pref(id + 1)
	}

	br := make([]int, n)
	copy(br, B)

	sort.Ints(br)
	reverse(br)

	for i := 0; i < n; i++ {
		j := i
		for B[j] != br[i] {
			j++
		}
		// always have j, to sort B and using the insert operations
		move_to_front(j)
	}

	for i := 0; i < n; i++ {
		lst := i
		for lst < n && B[lst] == B[i] {
			lst++
		}
		splits = append(splits, lst-i)
		i = lst - 1
	}

	return true, ops, splits
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
