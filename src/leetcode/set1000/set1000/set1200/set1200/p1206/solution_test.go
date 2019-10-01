package p1206

import "testing"

func TestSample1(t *testing.T) {
	list := Constructor()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.Search(0) {
		t.Fatalf("Sample 1 fail to search 0")
	}
	list.Add(4)

	if !list.Search(1) {
		t.Fatalf("Sample 1 fail to search 1")
	}

	if list.Erase(0) {
		t.Fatalf("Sample 1 fail to erase 0")
	}

	if !list.Erase(1) {
		t.Fatalf("Sample 1 fail to erase 1")
	}
	if list.Search(1) {
		t.Fatalf("Sample 1 fail to search 1, again")
	}
}

func TestSample2(t *testing.T) {
	list := Constructor()
	list.Add(0)
	list.Add(5)
	list.Add(2)
	list.Add(1)

	if !list.Search(0) {
		t.Fatalf("Fail at step 0")
	}

	if !list.Erase(5) {
		t.Fatalf("Fail at step 1")
	}

	if !list.Search(2) {
		t.Fatalf("Fail at step 2")
	}

	if list.Erase(3) {
		t.Fatalf("Fail at step 3")
	}
	if !list.Search(2) {
		t.Fatalf("Fail at step 4")
	}
}
