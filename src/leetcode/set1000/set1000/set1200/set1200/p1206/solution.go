package p1206

import (
	"math"
	"math/rand"
	"time"
)

type Node struct {
	forward []*Node
	value   int
	count   int
}

type Skiplist struct {
	header *Node
	level  int
}

func Constructor() Skiplist {
	header := new(Node)
	header.value = math.MaxInt32
	header.forward = make([]*Node, 32)
	rand.Seed(time.Now().UnixNano())
	return Skiplist{header, 0}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.header

	for i := this.level; i >= 0; i-- {
		if cur.forward[i] != nil && cur.forward[i].value < target {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]

	return cur != nil && cur.value == target
}

func randomLevel() int {

	level := 0
	for rand.Float64() < 0.5 && level < 30 {
		level++
	}
	return level
}

func (this *Skiplist) Add(num int) {
	update := make([]*Node, 32)

	cur := this.header

	for i := this.level; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].value < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}

	cur = cur.forward[0]

	if cur == nil || cur.value != num {
		rlevel := randomLevel()
		if rlevel > this.level {
			for i := this.level + 1; i <= rlevel; i++ {
				update[i] = this.header
			}
			this.level = rlevel
		}
		node := new(Node)
		node.value = num
		node.count = 1
		node.forward = make([]*Node, rlevel+1)
		for i := 0; i <= rlevel; i++ {
			node.forward[i] = update[i].forward[i]
			update[i].forward[i] = node
		}
	} else if cur != nil {
		cur.count++
	}
}

func (this *Skiplist) Erase(num int) bool {
	update := make([]*Node, 32)
	cur := this.header

	for i := this.level; i >= 0; i-- {
		if cur.forward[i] != nil && cur.forward[i].value < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}

	cur = cur.forward[0]

	if cur == nil || cur.value != num {
		return false
	}

	// cur.value == num
	cur.count--

	if cur.count == 0 {
		// need to remove it
		for i := this.level; i >= 0; i-- {
			if update[i].forward[i] == cur {
				update[i].forward[i] = cur.forward[i]
			}
		}
	}

	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
