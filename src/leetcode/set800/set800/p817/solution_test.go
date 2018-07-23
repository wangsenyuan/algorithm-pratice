package p817

import "testing"

func runSample(t *testing.T, head *ListNode, G []int, expect int) {
	res := numComponents(head, G)
	if res != expect {
		t.Errorf("sample expect %d, but got %d", expect, res)
	}
}

func constructListNode(arr []int) *ListNode {
	var loop func(i int) *ListNode
	loop = func(i int) *ListNode {
		if i == len(arr) {
			return nil
		}
		return &ListNode{arr[i], loop(i + 1)}
	}

	return loop(0)
}

func TestSample1(t *testing.T) {
	head := constructListNode([]int{0, 1, 2, 3})
	G := []int{0, 1, 3}
	expect := 2
	runSample(t, head, G, expect)
}

func TestSample2(t *testing.T) {
	head := constructListNode([]int{0, 1, 2, 3, 4})
	G := []int{0, 3, 1, 4}
	expect := 2
	runSample(t, head, G, expect)
}

func TestSample3(t *testing.T) {
	head := constructListNode([]int{0, 1, 2})
	G := []int{0, 2}
	expect := 2
	runSample(t, head, G, expect)
}

func TestSample4(t *testing.T) {
	head := constructListNode([]int{0, 1, 2})
	G := []int{0, 1}
	expect := 1
	runSample(t, head, G, expect)
}
