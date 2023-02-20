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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int) int {
	// if mex = 0
	// 不能出现0, max(A[i] - k, 0) > 0
	// min(A[i]) - 1
	// 如果 mex = x > 0
	// 数组中存在 0, 1, 2, ... x - 1
	//  max(A[i] - k, 0) 仍然需要保证 0, 1, 2, .. x - 1存在
	// k = 1, 那么所有的数字-1, 0, 0, 1, 2, ... x - 2, ?
	// 为了让mex = x, 必须存在 x - 1, 但是这个不可能，因为假设存在x-1, 那么原来就需要存在x
	// k = 2, 0, 0, 0, 1, ... x - 3, x - 2, x - 1, ? 同理也不行
	// 如果 k = x, 会怎么样呢？前半部分全部变为0，
	// 0， 0， 0， .... x + 1 - k = 1, 只要后半部也存在，就是ok的
	// 原来的值应该至少是 0, 1, 2, 3... x- 1, x + 1, x + 2, ... x + x - 1, x + x + j (j > 0)
	// 如果 k = x + 1,  0..... 0, 1, ... x - 2, x + j - 1
	sort.Ints(A)
	if A[0] > 0 {
		return A[0] - 1
	}

	A = unique(A)
	n := len(A)

	var dfs func(i int, x, k int, ans int) int

	dfs = func(i int, x, k int, ans int) int {
		if i == n {
			return ans
		}
		// 0, ... x - 1 exists before i
		j := i
		for j < n && A[j]-k == j-i+1 && j-i < x-1 {
			j++
		}
		if j-i < x-1 || A[j-1]-k == x {
			return ans
		}
		return dfs(j, x, k+x, ans+1)
	}

	var mex int

	for i := 0; i < n; i++ {
		if A[i] < mex {
			continue
		}
		if A[i] == mex {
			mex++
			continue
		}
		// A[i] > mex
		if mex == 1 {
			break
		}
		// 分段 [0...x-1], [a, a + 1,  ... a + m], [b...b + n]
		// 每一段都可以贡献k =  a - 1
		// 只要这段的长度 >= mex - 1
		var cnt int
		for j := i; j < n; {
			k := j + 1
			for k < n && A[k] == A[k-1]+1 {
				k++
			}
			if k-j >= mex-1 {
				cnt++
			}
			j = k
		}
		return cnt
	}
	if mex == 1 {
		// 始终可以获得0， 可以有无穷个
		return -1
	}
	// mex = n
	// 只要减去一个数字， 就无法得到mex = n
	return 0
}

func unique(arr []int) []int {
	n := 0
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}
