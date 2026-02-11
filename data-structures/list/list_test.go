package list

import (
	"reflect"
	"testing"
)

func collect[T any](l *List[T]) []T {
	var out []T
	l.ForEach(func(v T) bool {
		out = append(out, v)
		return true
	})
	return out
}

func TestList_ZeroValue(t *testing.T) {
	var l List[int]

	if l.Len() != 0 {
		t.Fatalf("expected Len=0, got %d", l.Len())
	}
	if l.Front() != nil {
		t.Fatalf("expected Front=nil on empty list")
	}
	if l.Back() != nil {
		t.Fatalf("expected Back=nil on empty list")
	}

	// must be usable
	l.PushBack(1)
	if l.Len() != 1 {
		t.Fatalf("expected Len=1 after PushBack")
	}
}

func TestList_PushFrontBack_Order(t *testing.T) {
	l := New[int]()

	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)

	got := collect(l)
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestList_Remove(t *testing.T) {
	l := New[int]()

	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	e3 := l.PushBack(3)

	// remove middle
	v, ok := l.Remove(e2)
	if !ok || v != 2 {
		t.Fatalf("expected Remove(2) ok")
	}

	got := collect(l)
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	// remove front
	l.Remove(e1)
	got = collect(l)
	want = []int{3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	// remove back
	l.Remove(e3)
	if l.Len() != 0 {
		t.Fatalf("expected empty list")
	}

	// remove invalid
	if _, ok := l.Remove(e3); ok {
		t.Fatalf("expected removing detached element to fail")
	}
}

func TestList_InsertBeforeAfter(t *testing.T) {
	l := New[int]()

	e2 := l.PushBack(2)
	l.InsertBefore(1, e2)
	l.InsertAfter(3, e2)

	got := collect(l)
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestList_MoveToFrontBack(t *testing.T) {
	l := New[int]()

	e1 := l.PushBack(1)
	l.PushBack(2)
	e3 := l.PushBack(3)

	l.MoveToFront(e3)

	got := collect(l)
	want := []int{3, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	l.MoveToBack(e3)

	got = collect(l)
	want = []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	// no-op move
	l.MoveToFront(e1)
	got = collect(l)
	want = []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestList_Clear(t *testing.T) {
	l := New[int]()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Clear()

	if l.Len() != 0 {
		t.Fatalf("expected Len=0 after Clear")
	}
	if l.Front() != nil {
		t.Fatalf("expected Front=nil after Clear")
	}

	// must still be usable
	l.PushBack(10)
	if l.Len() != 1 {
		t.Fatalf("expected Len=1 after reuse")
	}
}
