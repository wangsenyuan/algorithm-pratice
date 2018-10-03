package p916

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, A, B []string, expect []string) {
	res := wordSubsets(A, B)
	sort.Strings(expect)
	sort.Strings(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"e", "o"}
	expect := []string{"facebook", "google", "leetcode"}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"l", "e"}
	expect := []string{"apple", "google", "leetcode"}
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"e", "oo"}
	expect := []string{"facebook", "google"}
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"lo", "eo"}
	expect := []string{"google", "leetcode"}
	runSample(t, A, B, expect)
}
func TestSample5(t *testing.T) {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"ec", "oc", "ceo"}
	expect := []string{"facebook", "leetcode"}
	runSample(t, A, B, expect)
}
