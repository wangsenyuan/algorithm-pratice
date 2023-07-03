package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	events := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var tp int
		pos := readInt(s, 0, &tp)
		if tp == 1 || tp == 2 {
			var u, v, a, b int
			pos = readInt(s, pos+1, &u)
			pos = readInt(s, pos+1, &v)
			pos = readInt(s, pos+1, &a)
			pos = readInt(s, pos+1, &b)
			events[i] = []int{tp, u, v, a, b}
		} else {
			var u int
			readInt(s, pos+1, &u)
			events[i] = []int{tp, u}
		}
	}
	res := solve(n, events)
	var buf bytes.Buffer

	for _, x := range res {
		if x.ok {
			buf.WriteString(fmt.Sprintf("%d\n", x.val))
		} else {
			buf.WriteString("NA\n")
		}
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type Ans struct {
	ok  bool
	val int64
}

func solve(n int, events [][]int) []Ans {
	root := NewNode(1, n)

	var ans []Ans

	for _, cur := range events {
		if cur[0] == 1 {
			u, v, a, b := cur[1], cur[2], cur[3], cur[4]
			price := NewItem(int64(a), int64(a)*int64(-u)+int64(b))
			root.UpdatePrice(*price, u, v)
		} else if cur[0] == 2 {
			u, v, a, b := cur[1], cur[2], cur[3], cur[4]
			fee := NewItem(int64(a), int64(a)*int64(-u)+int64(b))
			root.UpdateFee(*fee, u, v)
		} else {
			u := cur[1]
			var fee int64
			var price int64 = math.MinInt64
			var ok bool
			root.Get(u, &fee, &price, &ok)
			if !ok {
				ans = append(ans, Ans{false, -1})
			} else {
				ans = append(ans, Ans{true, fee + price})
			}
		}
	}

	return ans
}

type Item struct {
	a int64
	b int64
}

func NewItem(a, b int64) *Item {
	return &Item{a, b}
}

func (this *Item) Add(that Item) {
	this.a += that.a
	this.b += that.b
}

func (this Item) Calc(x int64) int64 {
	return this.a*x + this.b
}

func (this *Item) Clear() {
	this.a = 0
	this.b = 0
}

type Node struct {
	l       int
	r       int
	covered bool
	price   *Item
	fee     *Item
	left    *Node
	right   *Node
}

func NewNode(l int, r int) *Node {
	node := new(Node)
	node.l = l
	node.r = r
	node.price = NewItem(0, 0)
	node.fee = NewItem(0, 0)
	return node
}

func (node *Node) updateLazy() {
	if node.l == node.r {
		return
	}
	mid := (node.l + node.r) / 2
	if node.left == nil {
		node.left = NewNode(node.l, mid)
		node.right = NewNode(mid+1, node.r)
	}
	node.left.fee.Add(*node.fee)
	node.right.fee.Add(*node.fee)
	node.fee.Clear()
}

func (node *Node) UpdateFee(fee Item, L int, R int) {
	if R < node.l || node.r < L {
		return
	}
	if L <= node.l && node.r <= R {
		node.fee.Add(fee)
		return
	}
	node.updateLazy()
	node.left.UpdateFee(fee, L, R)
	node.right.UpdateFee(fee, L, R)
}

func (node *Node) UpdatePrice(price Item, L int, R int) {
	if R < node.l || node.r < L {
		return
	}
	node.updateLazy()
	if L <= node.l && node.r <= R {
		if !node.covered {
			node.covered = true
			// 直接去地址应该不会共享吧
			node.price = NewItem(price.a, price.b)
			return
		}
		l1 := node.price
		l2 := NewItem(price.a, price.b)
		l1y1 := l1.Calc(int64(node.l))
		l2y1 := l2.Calc(int64(node.l))
		if l1y1 < l2y1 {
			l1, l2 = l2, l1
		}
		l1y2 := l1.Calc(int64(node.r))
		l2y2 := l2.Calc(int64(node.r))
		if l1y2 >= l2y2 {
			node.price = l1
			return
		}
		mid := (node.l + node.r) / 2
		l1m := l1.Calc(int64(mid))
		l2m := l2.Calc(int64(mid))
		if l1m > l2m {
			node.price = l1
			node.right.UpdatePrice(*l2, L, R)
		} else {
			node.price = l2
			node.left.UpdatePrice(*l1, L, R)
		}
	} else {
		node.left.UpdatePrice(price, L, R)
		node.right.UpdatePrice(price, L, R)
	}
}

func (node *Node) Get(p int, fee *int64, price *int64, ok *bool) {
	if p < node.l || node.r < p {
		return
	}
	*fee += node.fee.Calc(int64(p))
	if node.covered {
		*ok = true
		*price = max(*price, node.price.Calc(int64(p)))
	}
	if node.left != nil {
		node.left.Get(p, fee, price, ok)
	}
	if node.right != nil {
		node.right.Get(p, fee, price, ok)
	}
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
