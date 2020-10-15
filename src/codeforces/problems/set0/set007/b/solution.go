package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t, m := readTwoNums(reader)

	manager := NewManager(m)

	var buf bytes.Buffer

	for t > 0 {
		t--
		s, _ := reader.ReadBytes('\n')

		if s[0] == 'a' {
			var n int
			readInt(s, len("alloc "), &n)
			id := manager.Alloc(n)

			if id < 0 {
				buf.WriteString("NULL\n")
			} else {
				buf.WriteString(fmt.Sprintf("%d\n", id))
			}
		} else if s[0] == 'e' {
			var x int
			readInt(s, len("erase "), &x)
			ok := manager.Erase(x)
			if !ok {
				buf.WriteString("ILLEGAL_ERASE_ARGUMENT\n")
			}
		} else {
			manager.Defragment()
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

type Node struct {
	next *Node
	sz   int
	pos  int
	id   int
}

func NewNode(pos int, sz int) *Node {
	node := new(Node)
	node.pos = pos
	node.sz = sz
	return node
}

type Manager struct {
	free *Node
	used *Node
	mem  map[int]*Node
	p    int
	tot  int
}

func NewManager(m int) Manager {
	free := new(Node)
	free.next = NewNode(0, m)
	used := new(Node)
	mem := make(map[int]*Node)
	return Manager{free, used, mem, 0, m}
}

func (this *Manager) Alloc(n int) int {
	head := this.free

	for head.next != nil && head.next.sz < n {
		head = head.next
	}

	if head.next == nil {
		return -1
	}

	tmp := head.next

	pos := tmp.pos

	var tmp2 *Node

	if tmp.sz > n {
		tmp.sz -= n
		tmp.pos += n
		tmp2 = NewNode(pos, n)
	} else {
		head.next = tmp.next
		tmp.next = nil
		tmp2 = tmp
	}

	head = this.used

	for head.next != nil && head.next.pos < pos {
		head = head.next
	}

	tmp2.next = head.next
	head.next = tmp2

	this.p++
	this.mem[this.p] = tmp2
	tmp2.id = this.p

	return this.p
}

func (this *Manager) Erase(x int) bool {
	if node, found := this.mem[x]; !found {
		return false
	} else {
		delete(this.mem, x)

		head := this.used
		//remove it from used
		for head.next != nil && head.next != node {
			head = head.next
		}

		if head.next == nil {
			return false
		}

		head.next = node.next
		node.next = nil
		// put back as free
		head = this.free

		for head.next != nil && head.next.pos < node.pos {
			head = head.next
		}
		// head.pos < pos && ( head.next == nil || head.next.pos > pos)

		node.id = -1
		node.next = head.next
		head.next = node

		if head == this.free {
			head = head.next
		}

		for head.next != nil && head.pos+head.sz == head.next.pos {
			head.sz += head.next.sz
			head.next = head.next.next
		}

		return true
	}
}

func (this *Manager) Defragment() {
	head := this.used
	var sz int

	for head.next != nil {
		cur := head.next
		cur.pos = sz
		sz += cur.sz
		head = head.next
	}

	if this.tot > sz {
		this.free.next = NewNode(sz, this.tot-sz)
	} else {
		this.free.next = nil
	}
}
