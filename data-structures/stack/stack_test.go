package stack

import "testing"

func TestStack_ZeroValue(t *testing.T) {
	var st Stack[int]

	if st.Len() != 0 {
		t.Fatalf("expected Len=0, got %d", st.Len())
	}

	if _, ok := st.Peek(); ok {
		t.Fatalf("expected Peek on empty stack to return ok=false")
	}

	if _, ok := st.Pop(); ok {
		t.Fatalf("expected Pop on empty stack to return ok=false")
	}

	st.Push(1)
	st.Push(2)

	if st.Len() != 2 {
		t.Fatalf("expected Len=2, got %d", st.Len())
	}

	v, ok := st.Peek()
	if !ok || v != 2 {
		t.Fatalf("expected Peek=(2,true), got (%v,%v)", v, ok)
	}
}

func TestStack_New(t *testing.T) {
	st := New[int](10)
	if st == nil {
		t.Fatalf("expected non-nil stack")
	}
	if st.Len() != 0 {
		t.Fatalf("expected Len=0, got %d", st.Len())
	}
	if st.Cap() < 10 {
		t.Fatalf("expected Cap>=10, got %d", st.Cap())
	}
}

func TestStack_PushPop_Order(t *testing.T) {
	var st Stack[int]

	for i := 1; i <= 5; i++ {
		st.Push(i)
	}

	for i := 5; i >= 1; i-- {
		v, ok := st.Pop()
		if !ok {
			t.Fatalf("expected ok=true at i=%d", i)
		}
		if v != i {
			t.Fatalf("expected %d, got %d", i, v)
		}
	}

	if st.Len() != 0 {
		t.Fatalf("expected Len=0 after popping all, got %d", st.Len())
	}
}

func TestStack_Peek_DoesNotRemove(t *testing.T) {
	var st Stack[string]
	st.Push("a")
	st.Push("b")

	v, ok := st.Peek()
	if !ok || v != "b" {
		t.Fatalf("expected Peek=(b,true), got (%v,%v)", v, ok)
	}

	if st.Len() != 2 {
		t.Fatalf("expected Len=2 after Peek, got %d", st.Len())
	}

	v2, ok := st.Pop()
	if !ok || v2 != "b" {
		t.Fatalf("expected Pop=(b,true), got (%v,%v)", v2, ok)
	}
}

func TestStack_Clear(t *testing.T) {
	var st Stack[int]
	for i := 0; i < 100; i++ {
		st.Push(i)
	}

	st.Clear()

	if st.Len() != 0 {
		t.Fatalf("expected Len=0 after Clear, got %d", st.Len())
	}

	// Must still be usable after Clear
	st.Push(42)
	v, ok := st.Pop()
	if !ok || v != 42 {
		t.Fatalf("expected Pop=(42,true) after Clear+Push, got (%v,%v)", v, ok)
	}
}

func TestStack_Cap_Growth_Sanity(t *testing.T) {
	st := New[int](10)
	if st.Cap() < 1 {
		t.Fatalf("expected Cap>=1, got %d", st.Cap())
	}

	// push enough to force growth
	for i := 0; i < 1000; i++ {
		st.Push(i)
	}

	if st.Len() != 1000 {
		t.Fatalf("expected Len=1000, got %d", st.Len())
	}
	if st.Cap() < st.Len() {
		t.Fatalf("expected Cap>=Len, got Cap=%d Len=%d", st.Cap(), st.Len())
	}
}
