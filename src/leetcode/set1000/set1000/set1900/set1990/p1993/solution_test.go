package p1993

import "testing"

func TestSample1(t *testing.T) {
	tree := Constructor([]int{-1, 4, 0, 2, 5, 0, 4, 8, 2, 4})
	res := tree.Upgrade(9, 22)
	require(t, res, false, "upgrade at step 0 should get false")
	res = tree.Unlock(6, 21)
	require(t, res, false, "unlock at step 1 should get false")
	res = tree.Unlock(8, 40)
	require(t, res, false, "unlock at step 2 should get false")
	res = tree.Upgrade(3, 24)
	require(t, res, false, "upgrade at step 3 should get false")
	res = tree.Unlock(3, 24)
	require(t, res, false, "unlock at step 4 should get false")
	res = tree.Upgrade(6, 17)
	require(t, res, false, "unlock at step 5 should get false")
	res = tree.Upgrade(4, 42)
	require(t, res, false, "upgrade at step 6 should get false")
	res = tree.Lock(6, 41)
	require(t, res, true, "lock at step 7 (6, 41) should get true")
	res = tree.Unlock(5, 40)
	require(t, res, false, "unlock at step 8 should get false")
	res = tree.Lock(6, 14)
	require(t, res, false, "lock at step 9 should get false")
	res = tree.Upgrade(9, 6)
	require(t, res, false, "upgrade at step 10 should get false")
}

func require(t *testing.T, res, expect bool, msg string) {
	if res != expect {
		t.Fatal(msg)
	}
}
