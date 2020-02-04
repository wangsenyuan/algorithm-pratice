package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		N, K, M := line[0], line[1], line[2]
		trips := make([][]int, M)
		for i := 0; i < M; i++ {
			trips[i] = readNNums(scanner, 3)
		}
		res := solve(N, K, M, trips)
		if len(res) == 0 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
			for i := 0; i < N; i++ {
				if i < N-1 {
					fmt.Printf("%d ", res[i])
				} else {
					fmt.Printf("%d\n", res[i])
				}
			}
		}
	}
}

func solve(N, K, M int, trips [][]int) []int {
	arr := make([]int, 2*N+2)

	for i := 0; i < N; i++ {
		arr[i+N] = i
	}

	for i := N - 1; i >= 0; i-- {
		arr[i] = min(arr[i<<1], arr[i<<1|1])
	}

	update := func(pos int, v int) {
		pos += N
		arr[pos] = v

		for pos > 0 {
			arr[pos>>1] = min(arr[pos], arr[pos^1])
			pos >>= 1
		}
	}

	getMinValue := func(l, r int) int {
		l += N
		r += N
		res := N + 1

		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}

			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}

		return res
	}

	res := make([]int, N)

	sort.Sort(Ints(trips))

	for _, trip := range trips {
		x, _, z := trip[0], trip[1], trip[2]
		z--
		if res[z] > 0 && res[z] != x {
			return nil
		}
		res[z] = x
		update(z, N+1)
	}

	cnt := make([]int, K+1)
	last := make([]int, K+1)

	for i, ind := 0, 0; i < N; i++ {
		for ind < M && trips[ind][2] == i+1 {
			x, y, z := trips[ind][0], trips[ind][1], trips[ind][2]
			z--
			if cnt[x] > y-1 {
				return nil
			}
			for last[x] < i && cnt[x] < y-1 {
				nxt := getMinValue(last[x], N)
				if nxt >= i {
					return nil
				}
				res[nxt] = x
				update(nxt, N+1)
				cnt[x]++
				last[x] = nxt
			}
			if cnt[x] != y-1 {
				return nil
			}
			last[x] = z
			ind++
		}
		if res[i] != 0 {
			cnt[res[i]]++
		}
	}

	p := getMinValue(0, N)

	if p >= N {
		return res
	}

	vv := -1
	if p < N {
		for x := 1; x <= K; x++ {
			if cnt[x] == 0 {
				vv = x
				break
			}
			if last[x] < p {
				vv = x
				break
			}
		}
	}

	if vv == -1 {
		return nil
	}

	for i := 0; i < N; i++ {
		if res[i] == 0 {
			res[i] = vv
		}
	}

	return res
}

type Ints [][]int

func (ints Ints) Len() int {
	return len(ints)
}

func (ints Ints) Less(i, j int) bool {
	a := ints[i]
	b := ints[j]

	if a[2] < b[2] {
		return true
	}

	if a[2] == b[2] && a[1] < b[1] {
		return true
	}

	return false
}

func (ints Ints) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
