package p1993

type LockingTree struct {
	n      int
	in     []int
	out    []int
	locks  []int
	parent []int
	g      *Graph
}

func Constructor(parent []int) LockingTree {
	n := len(parent)
	g := NewGraph(n, n)
	for i := 1; i < n; i++ {
		g.AddEdge(parent[i], i)
	}

	in := make([]int, n)
	out := make([]int, n)
	var dfs func(u int, t *int)
	dfs = func(u int, t *int) {
		in[u] = *t
		*t++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			dfs(g.to[i], t)
		}
		*t++
		out[u] = *t
	}
	var time int
	dfs(0, &time)

	locks := make([]int, n)

	return LockingTree{n, in, out, locks, parent, g}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.locks[num] > 0 {
		return false
	}
	this.locks[num] = user

	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locks[num] != user {
		return false
	}

	this.locks[num] = 0

	return true
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	cur := num

	for cur >= 0 {
		if this.locks[cur] > 0 {
			return false
		}
		cur = this.parent[cur]
	}
	var ok bool
	var dfs func(u int)

	dfs = func(u int) {
		if this.locks[u] > 0 {
			this.Unlock(u, this.locks[u])
			ok = true
		}
		for i := this.g.nodes[u]; i > 0; i = this.g.next[i] {
			v := this.g.to[i]
			dfs(v)
		}
	}

	dfs(num)

	if !ok {
		return false
	}

	return this.Lock(num, user)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
