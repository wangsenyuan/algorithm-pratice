package main

import "testing"

func TestSample1(t *testing.T) {
	manager := NewManager(10)

	id := manager.Alloc(5)

	expect(id == 1, t, "expect id to be 1")

	id = manager.Alloc(3)

	expect(id == 2, t, "expect id to be 2")

	ok := manager.Erase(1)

	expect(ok, t, "should erase 1 successfully")

	id = manager.Alloc(6)

	expect(id == -1, t, "should not be able to alloc 6")

	manager.Defragment()

	id = manager.Alloc(6)

	expect(id == 3, t, "should alloc 6 bytes, with id = 3")
}

func expect(expr bool, t *testing.T, message string) {
	if !expr {
		t.Fatal(message)
	}
}
