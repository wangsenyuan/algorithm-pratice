package p2296

import "testing"

func TestSample1(t *testing.T) {
	editor := Constructor()
	editor.AddText("leetcode")
	cnt := editor.DeleteText(4)

	if cnt != 4 {
		t.Fatalf("expect 4, but got %d", cnt)
	}
	editor.AddText("practice")
	res := editor.CursorRight(3)
	if res != "etpractice" {
		t.Fatalf("Sample expect etpractice, but got %s", res)
	}

	res = editor.CursorLeft(8)

	if res != "leet" {
		t.Fatalf("Sample expect leet, but got %s", res)
	}

	cnt = editor.DeleteText(10)

	if cnt != 4 {
		t.Fatalf("Sample expect 4, but got %d", cnt)
	}

	res = editor.CursorLeft(2)

	if res != "" {
		t.Fatalf("Sample expect <empty> string, but got %s", res)
	}

	res = editor.CursorRight(6)

	if res != "practi" {
		t.Fatalf("Sample expect practi, but got %s", res)
	}
}
