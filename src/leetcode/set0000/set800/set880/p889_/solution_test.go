package p889_

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, pre []int, post []int) {
	tree := constructFromPrePost(pre, post)
	a := preOrderTraversal(tree)
	if !reflect.DeepEqual(a, pre) {
		t.Errorf("Sample %v %v, got answer %v, not correct", pre, post, a)
	}
	b := postOrderTraversal(tree)
	if !reflect.DeepEqual(b, post) {
		t.Errorf("Sample %v %v, got answer %v, not correct", pre, post, b)
	}
}

func TestSample1(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	runSample(t, pre, post)
}
