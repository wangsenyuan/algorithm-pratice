package p1192

import (
	"reflect"
	"sort"
	"testing"
)

type AS [][]int

func (this AS) Len() int {
	return len(this)
}

func (this AS) Less(i, j int) bool {
	if this[i][0] < this[j][0] {
		return true
	}
	if this[i][0] == this[j][0] {
		return this[i][1] < this[j][1]
	}
	return false
}

func (this AS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortArrs(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		sort.Ints(arr[i])
	}

	sort.Sort(AS(arr))
}

func runSample(t *testing.T, n int, connections [][]int, expect [][]int) {
	res := criticalConnections(n, connections)

	sortArrs(expect)
	sortArrs(res)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %d %v, expect %v , but got %v", n, connections, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	connections := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}
	expect := [][]int{{1, 3}}
	runSample(t, n, connections, expect)
}
