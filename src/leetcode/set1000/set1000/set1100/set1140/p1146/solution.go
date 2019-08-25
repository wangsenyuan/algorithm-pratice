package p1146

import "sort"

type SnapshotArray struct {
	arr [][]int
	snp [][]int
	cur int
}

func Constructor(length int) SnapshotArray {
	arr := make([][]int, length)
	snp := make([][]int, length)

	for i := 0; i < length; i++ {
		arr[i] = make([]int, 0, 10)
		snp[i] = make([]int, 0, 10)
	}
	return SnapshotArray{arr, snp, 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	a := this.arr[index]
	b := this.snp[index]
	if len(b) == 0 || b[len(b)-1] < this.cur {
		a = append(a, val)
		b = append(b, this.cur)
	} else {
		a[len(a)-1] = val
		b[len(a)-1] = this.cur
	}
	this.arr[index] = a
	this.snp[index] = b
}

func (this *SnapshotArray) Snap() int {
	this.cur++
	return this.cur - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	a := this.arr[index]
	b := this.snp[index]

	i := sort.Search(len(b), func(i int) bool {
		return b[i] > snap_id
	})
	if i > 0 {
		return a[i-1]
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
