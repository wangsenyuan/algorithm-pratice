package p2071

import "testing"

func runSample(t *testing.T, arr []int, expect []int) {
	head := createList(arr)
	res := reverseEvenLengthGroups(head)
	var i int
	for res != nil {
		if res.Val != expect[i] {
			t.Fatalf("Sample expect %v, but fail at %d with value %d", expect, i, res.Val)
		}
		i++
		res = res.Next
	}
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	node := new(ListNode)
	node.Val = arr[0]
	node.Next = createList(arr[1:])
	return node
}

func TestSample1(t *testing.T) {
	arr := []int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4}
	expect := []int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7}
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4, 1}
	expect := []int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7, 1}
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4, 1, 2}
	expect := []int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7, 2, 1}
	runSample(t, arr, expect)
}
