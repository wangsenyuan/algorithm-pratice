package main

import (
	"bufio"
	"container/heap"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007
const X = 40000

func solve(n int, A []int) int {
	cnt := make(map[int][]int)

	update := func(x, c int) {
		if v, found := cnt[x]; !found {
			cnt[x] = make([]int, 0, 10)
			cnt[x] = append(cnt[x], c)
		} else {
			cnt[x] = append(v, c)
		}
	}

	for _, num := range A {
		x := 2
		for x*x <= num {
			if num%x == 0 {
				var c int
				for num%x == 0 {
					c++
					num /= x
				}
				update(x, c)
			}
			x++
		}

		if num > 1 {
			update(num, 1)
		}
	}

	var res int64 = 1

	for k, v := range cnt {
		x := findMinSum(v)
		if x >= X {
			if len(v)%2 == 1 {
				res *= int64(k)
				res %= MOD
			}
			continue
		}
		res *= int64(pow(k, x))
		res %= MOD
	}

	return int(res)
}

func pow(a, b int) int {
	A := int64(a)
	r := int64(1)

	for b > 0 {
		if b&1 == 1 {
			r *= A
			r %= MOD
		}
		A *= A
		A %= MOD
		b >>= 1
	}

	return int(r)
}

func findMinSum(arr []int) int {
	h := IntHeap(arr)

	heap.Init(&h)

	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)
		a--
		b--
		if a > 0 {
			heap.Push(&h, a)
		}
		if b > 0 {
			heap.Push(&h, b)
		}
	}

	if h.Len() > 0 {
		return heap.Pop(&h).(int)
	}

	return 0
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
