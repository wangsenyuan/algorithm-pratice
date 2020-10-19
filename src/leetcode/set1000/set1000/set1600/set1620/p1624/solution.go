package p1624

const MOD = 1000000007
const N = 100000

//
//type Fancy struct {
//	ops  []*Node
//	nums []int
//}
//
//func Constructor() Fancy {
//	ops := make([]*Node, 2*N)
//	for i := 0; i < len(ops); i++ {
//		ops[i] = NewNode(1, 0)
//	}
//
//	nums := make([]int, 0, N)
//
//	return Fancy{ops, nums}
//}
//
//func (this *Fancy) Append(val int) {
//	this.nums = append(this.nums, val)
//}
//
//func (this *Fancy) AddAll(inc int) {
//	if len(this.nums) == 0 {
//		return
//	}
//	op := NewNode(1, int64(inc))
//
//	for i := N - len(this.nums); i < len(this.ops); i += i & -i {
//		this.ops[i].Concat(op)
//	}
//}
//
//func (this *Fancy) MultAll(m int) {
//	if len(this.nums) == 0 {
//		return
//	}
//	op := NewNode(int64(m), 0)
//
//	for i := N - len(this.nums); i < len(this.ops); i += i & -i {
//		this.ops[i].Concat(op)
//	}
//}
//
//func (this *Fancy) GetIndex(idx int) int {
//	if idx >= len(this.nums) {
//		return -1
//	}
//
//	op := NewNode(1, 0)
//
//	for i := N - (idx + 1); i > 0; i -= i & -i {
//		op.Concat(this.ops[i])
//	}
//
//	res := op.a * int64(this.nums[idx])
//	res %= MOD
//	res += op.b
//	res %= MOD
//	return int(res)
//}
//
//
//type Node struct {
//	a int64
//	b int64
//}
//
//func NewNode(a, b int64) *Node {
//	node := new(Node)
//	node.a = a
//	node.b = b
//	return node
//}
//
//func (f *Node) Concat(g *Node) {
//	//h := NewNode(0, 0)
//	f.a = (f.a * g.a) % MOD
//	f.b = (g.a*f.b + g.b) % MOD
//}

type Fancy struct {
	nums []int64
	a    []int64
	b    []int64
}

func Constructor() Fancy {
	nums := make([]int64, 0, N)
	a := make([]int64, 0, N)
	b := make([]int64, 0, N)
	a = append(a, 1)
	b = append(b, 0)
	return Fancy{nums, a, b}
}

func (this *Fancy) Append(val int) {
	this.nums = append(this.nums, int64(val))
	this.a = append(this.a, this.a[len(this.a)-1])
	this.b = append(this.b, this.b[len(this.b)-1])
}

func (this *Fancy) AddAll(inc int) {
	n := len(this.b)
	this.b[n-1] += int64(inc)
	this.b[n-1] %= MOD
}

func (this *Fancy) MultAll(m int) {
	n := len(this.a)
	this.a[n-1] *= int64(m)
	this.a[n-1] %= MOD
	this.b[n-1] *= int64(m)
	this.b[n-1] %= MOD
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.nums) {
		return -1
	}
	n := len(this.a)
	a0 := inverse(this.a[idx]) * this.a[n-1] % MOD
	b0 := (this.b[n-1] - this.b[idx]*a0%MOD + MOD) % MOD
	ans := a0*this.nums[idx]%MOD + b0
	ans %= MOD
	return int(ans)
}

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}

func inverse(num int64) int64 {
	return pow(num, MOD-2)
}
