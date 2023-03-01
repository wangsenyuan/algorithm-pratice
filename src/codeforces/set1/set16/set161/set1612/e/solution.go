package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	students := make([][]int, n)
	for i := 0; i < n; i++ {
		students[i] = readNNums(reader, 2)
	}

	res := solve(students)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

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

func solve(students [][]int) []int {
	// 假设选择了m条消息
	// mi <= 2 * 1e5
	// 被更多学生选择的消息可以优先选择（被选中的机会也大)
	// 假设按照选择次数排序后，
	// 假设一共pin了x个消息
	// 那么考虑任何一个学生，他能看到需要消息的期望值是多少呢？
	// 如果不包含他关心的消息, = 0
	// 如果包含， 如果x <= ki, 则为1， 否则为 C(x-1, k-1) / C(x, k) = k / x

	// 这里有个问题，如何找到这个x呢？首先这个函数不是线性的，所以是不能二分查找的（能变通吗？）
	// k <= 20可以用到了。
	// 分成两种情况x <= 20的情况， 这种直接暴力计算即可
	// x > 20, 首先不存在1的情况

	ids := make(map[int][]int)
	for i, st := range students {
		m := st[0]
		if len(ids[m]) == 0 {
			ids[m] = make([]int, 0, 1)
		}
		ids[m] = append(ids[m], i)
	}

	var arr []int

	for k := range ids {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	take := func(it []int, x int) float64 {
		var res float64
		for _, i := range it {
			st := students[i]
			k := st[1]
			if k >= x {
				res++
			} else {
				res += float64(k) / float64(x)
			}
		}
		return res
	}

	pq := make(FloatHeap, 0, len(arr))

	bruteForce := func(x int) float64 {
		// 选出x个消息，使得expect value最大
		// 应该选择，那些k值大于等于x的
		var sum float64
		for i := 0; i < x; i++ {
			val := take(ids[arr[i]], x)
			pr := Pair{val, i}
			heap.Push(&pq, pr)
			sum += val
		}

		for i := x; i < len(arr); i++ {
			val := take(ids[arr[i]], x)
			pr := Pair{val, i}
			heap.Push(&pq, pr)
			sum += val
			tmp := heap.Pop(&pq).(Pair)
			sum -= tmp.first
		}
		return sum
	}

	var expect float64
	var m int
	for x := 1; x <= len(arr) && x <= 20; x++ {
		tmp := bruteForce(x)
		if tmp > expect {
			expect = tmp
			m = x
		}
		for pq.Len() > 0 {
			heap.Pop(&pq)
		}
	}

	var res []int

	bruteForce(m)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(Pair)
		res = append(res, arr[cur.second])
	}

	return res
}

type Pair struct {
	first  float64
	second int
}

// An IntHeap is a min-heap of ints.
type FloatHeap []Pair

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i].first < h[j].first }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FloatHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Pair))
}

func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
