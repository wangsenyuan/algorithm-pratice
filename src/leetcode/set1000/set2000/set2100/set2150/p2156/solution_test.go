package p2156

import "testing"

func runSample(t *testing.T, s string, power int, modulo int, k int, hashValue int, expect string) {
	res := subStrHash(s, power, modulo, k, hashValue)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "leetcode"
	power := 7
	modulo := 20
	k := 2
	hashValue := 0
	expect := "ee"
	runSample(t, s, power, modulo, k, hashValue, expect)
}

func TestSample2(t *testing.T) {
	s := "fbxzaad"
	power := 31
	modulo := 100
	k := 3
	hashValue := 32
	expect := "fbx"
	runSample(t, s, power, modulo, k, hashValue, expect)
}
