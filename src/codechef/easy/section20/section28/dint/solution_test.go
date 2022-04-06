package main

import "testing"

func TestBitSet(t *testing.T) {

	set := NewBitSet(300)

	set.Set(100)

	res := set.Count()

	if res != 1 {
		t.Fatalf("expect 1, but got %d", res)
	}

	set.LeftShift(100)

	ok := set.IsSet(100)

	if ok {
		t.Fatalf("100 not set, but got true")
	}

	ok = set.IsSet(200)

	if !ok {
		t.Fatalf("200 should set, but got false")
	}

	set.RightShift(150)

	ok = set.IsSet(50)

	if !ok {
		t.Fatal("50 should set, but got false")
	}
	set.Set(64)

	res = set.Count()

	if res != 2 {
		t.Fatalf("expect 2 set digits, but got %d", res)
	}
}

func TestBitSet1(t *testing.T) {

	set := NewBitSet(1000)

	for i := 0; i < 1000; i++ {
		set.Set(i)
	}

	cnt := set.Count()

	if cnt != 1000 {
		t.Fatalf("Sample expect 1000, but got %d", cnt)
	}

	set.RightShift(64)

	cnt = set.Count()

	if cnt != 1000-64 {
		t.Fatalf("Sample expect 936, but got %d", cnt)
	}

	ok := set.IsSet(999)

	if ok {
		t.Fatal("999 should not be set, but got yes")
	}
}

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 8, 5}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 4}
	expect := 1
	runSample(t, A, expect)
}
