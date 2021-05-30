package p1884

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, servers []int, task []int, expect []int) {
	res := assignTasks(servers, task)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	servers := []int{3, 3, 2}
	tasks := []int{1, 2, 3, 2, 1, 2}
	expect := []int{2, 2, 0, 2, 1, 2}
	runSample(t, servers, tasks, expect)
}

func TestSample2(t *testing.T) {
	servers := []int{5, 1, 4, 3, 2}
	tasks := []int{2, 1, 2, 4, 5, 2, 1}
	expect := []int{1, 4, 1, 4, 1, 3, 2}
	runSample(t, servers, tasks, expect)
}
