package p1286

import "testing"

func TestSample1(t *testing.T) {
	iter := Constructor("abc", 2)
	res := iter.Next()
	if res != "ab" {
		t.Fatalf("Expect ab, but got %s", res)
	}
	res = iter.Next()

	if res != "ac" {
		t.Fatalf("Expect ac, but got %s", res)
	}

	res = iter.Next()

	if res != "bc" {
		t.Fatalf("Expect bc, but got %s", res)
	}
}

func TestSample2(t *testing.T) {
	iter := Constructor("abcd", 2)
	expect := []string{"ab", "ac", "ad", "bc", "bd", "cd"}

	for i := 0; i < len(expect); i++ {
		res := iter.Next()
		if res != expect[i] {
			t.Fatalf("expect %s, but got %s", expect[i], res)
		}
	}
}
