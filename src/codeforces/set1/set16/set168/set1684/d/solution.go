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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(k, A)
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

func solve(k int, traps []int) int64 {
	n := len(traps)
	// n <= 2 * 1e5
	// k <= n
	// 首先如果跳了x个trap，最好是跳过那些比较大的值
	//  而且越靠后越有利
	ts := make([]Trap, n)
	for i := 0; i < n; i++ {
		ts[i] = Trap{traps[i], i}
	}

	// 在只能删除一个情况下，对于x,它带来的收益是 damage[x] - (size - x)
	// size对于所有的都是constant， 所有应该选择 damage[x] + x 最大的那些
	sort.Slice(ts, func(i, j int) bool {
		a := ts[i].damage + ts[i].id
		b := ts[j].damage + ts[j].id
		return a >= b
	})
	var res int64

	for i := 0; i < n; i++ {
		res += int64(traps[i])
	}

	for x := 0; x < k; x++ {
		// skip trap i
		cur := ts[x]
		i := cur.id
		// 在它后面的一共有 n - i - 1
		// 其中已经跳过的有 get(n) - get(i)
		// get(i) 是在i前面发生的跳过，这些已经被加入到res中， 这样子，似乎是可以消除掉的？
		res += int64(n - i - 1 - x - cur.damage)
	}

	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Trap struct {
	damage int
	id     int
}

//
//func (this Trap) Less(that Trap) bool {
//	a := this.damage
//	b := that.damage
//	if this.id < that.id {
//		b++
//	} else {
//		a++
//	}
//
//	return a > b
//}
