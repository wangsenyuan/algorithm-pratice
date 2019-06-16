package p979

import "testing"

func runSample(t *testing.T, tree string, expect int) {
	root := ParseAsTree(tree)
	res := distributeCoins(root)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", tree, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "[3,0,0]", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "[0,3,0]", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "[1,0,2]", 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "[1,0,0,null,3]", 4)
}
