package main

import (
	"testing"
)

func runSample(t *testing.T, s []byte, expect bool) {
	can, res := solve(s)
	if can != expect {
		t.Errorf("sample %v, expect %t, but got %t", string(s), expect, can)
	} else if can {
		valid := make([]bool, len(s))
		for i := 0; i < len(res); i++ {
			valid[res[i]] = true
		}
		for i := 0; i < len(valid); i++ {
			if !valid[i] {
				t.Errorf("sample %v, result %v is a not valid permutation", string(s), res)
			}
		}
		tmp := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			tmp[i] = s[res[i]]
		}
		left, right := 0, len(s)-1
		for left < right {
			if tmp[left] != tmp[right] {
				break
			}
			left++
			right--
		}
		if right-left > 1 {
			t.Errorf("sample %v, answer %v, is not palidromic", string(s), string(tmp))
		}
	}
}

/**
aa
baa
abc
abab
*/

func TestSample1(t *testing.T) {
	s := []byte("aa")
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := []byte("baa")
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := []byte("abc")
	runSample(t, s, false)
}

func TestSample4(t *testing.T) {
	s := []byte("abab")
	runSample(t, s, true)
}

func TestSample5(t *testing.T) {
	s := []byte("a")
	runSample(t, s, true)
}

func TestSample6(t *testing.T) {
	s := []byte("abcc")
	runSample(t, s, false)
}

/**

szzsbbhkhk
sjkxkjxsoo
oppdooopoo
ialilpijpp
efefxecxe
ppxuouxpp
vyuvvvyuh
ggwgzzzeg
ggywwwgwg
**/
func TestSample7(t *testing.T) {
	runSample(t, []byte("ggxggvvxpp"), true)
}

func TestSample8(t *testing.T) {
	runSample(t, []byte("sjkxkjxsoo"), true)
}

func TestSample9(t *testing.T) {
	runSample(t, []byte("oppdooopoo"), false)
}

func TestSample10(t *testing.T) {
	runSample(t, []byte("ialilpijpp"), false)
}

func TestSample11(t *testing.T) {
	runSample(t, []byte("efefxecxe"), true)
}
