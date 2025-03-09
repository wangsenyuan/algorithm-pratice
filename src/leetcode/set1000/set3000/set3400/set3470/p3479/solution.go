package p3479

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	tr := BuildTree(baskets)

	var ans int

	for _, v := range fruits {
		i := tr.Find(v)
		if i < 0 {
			ans++
		} else {
			tr.Update(i, 0)
		}
	}
	return ans
}

type Tree struct {
	arr []int
	sz  int
}

func BuildTree(arr []int) *Tree {
	n := len(arr)
	res := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			res[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		res[i] = max(res[2*i+1], res[2*i+2])
	}
	build(0, 0, n-1)
	return &Tree{res, n}
}

func (tr *Tree) Find(v int) int {
	if tr.arr[0] < v {
		return -1
	}
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.arr[2*i+1] >= v {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, tr.sz-1)
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.arr[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.arr[i] = max(tr.arr[2*i+1], tr.arr[2*i+2])
	}
	loop(0, 0, tr.sz-1)
}
