package p872

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, tree1, tree2 string, expect bool) {
	root1 := ParseAsTree(tree1)
	root2 := ParseAsTree(tree2)
	res := leafSimilar(root1, root2)

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", tree1, tree2, expect, res)
	}
}

func runDfs(t *testing.T, tree string, expect []int) {
	res := dfs(ParseAsTree(tree))
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", tree, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runDfs(t, "[3,5,1,6,2,9,8,null,null,7,4]", []int{6, 7, 4, 9, 8})
}
