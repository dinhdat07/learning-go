package set

import (
	"testing"
)

func TestSet_ZeroValue(t *testing.T) {
	var s Set[int]

	if s.Len() != 0 {
		t.Fatalf("expected Len=0, got %d", s.Len())
	}

	if s.Has(1) {
		t.Fatalf("expected Has(1)=false on zero value set")
	}

	if added := s.Add(1); !added {
		t.Fatalf("expected Add(1)=true on zero value set")
	}
	if !s.Has(1) {
		t.Fatalf("expected Has(1)=true after Add")
	}
	if s.Len() != 1 {
		t.Fatalf("expected Len=1, got %d", s.Len())
	}
}

func TestSet_New(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Fatalf("expected non-nil set")
	}
	if s.Len() != 0 {
		t.Fatalf("expected Len=0, got %d", s.Len())
	}
	if s.Has(123) {
		t.Fatalf("expected Has(123)=false on empty set")
	}
}

func TestSet_AddHasRemove(t *testing.T) {
	s := New[string]()

	if !s.Add("a") {
		t.Fatalf("expected Add(a)=true")
	}
	if s.Add("a") {
		t.Fatalf("expected Add(a)=false when already exists")
	}
	if !s.Has("a") {
		t.Fatalf("expected Has(a)=true")
	}
	if s.Has("b") {
		t.Fatalf("expected Has(b)=false")
	}

	if !s.Remove("a") {
		t.Fatalf("expected Remove(a)=true when exists")
	}
	if s.Remove("a") {
		t.Fatalf("expected Remove(a)=false when already removed")
	}
	if s.Has("a") {
		t.Fatalf("expected Has(a)=false after Remove")
	}
}

func TestSet_AddMany(t *testing.T) {
	s := New[int]()

	added := s.AddMany(1, 2, 2, 3, 3, 3)
	if added != 3 {
		t.Fatalf("expected AddMany added=3, got %d", added)
	}
	if s.Len() != 3 {
		t.Fatalf("expected Len=3, got %d", s.Len())
	}
}

func TestSet_Clear(t *testing.T) {
	s := New[int]()
	s.AddMany(1, 2, 3)

	s.Clear()

	if s.Len() != 0 {
		t.Fatalf("expected Len=0 after Clear, got %d", s.Len())
	}
	if s.Has(1) {
		t.Fatalf("expected Has(1)=false after Clear")
	}

	// usable after Clear
	if !s.Add(10) {
		t.Fatalf("expected Add(10)=true after Clear")
	}
	if !s.Has(10) {
		t.Fatalf("expected Has(10)=true after Clear+Add")
	}
}

func TestSet_Values_SnapshotAndMembership(t *testing.T) {
	s := New[int]()
	s.AddMany(1, 2, 3)

	vals := s.Values()
	if len(vals) != 3 {
		t.Fatalf("expected Values len=3, got %d", len(vals))
	}

	seen := make(map[int]bool, 3)
	for _, v := range vals {
		seen[v] = true
	}
	for _, want := range []int{1, 2, 3} {
		if !seen[want] {
			t.Fatalf("expected Values to contain %d, got %v", want, vals)
		}
	}

	// Snapshot property: mutating returned slice must not affect the set
	vals[0] = 999
	if s.Has(999) {
		t.Fatalf("expected snapshot; modifying returned slice should not affect set")
	}
}

func TestSet_ZeroValueHas_NoAllocStateChange(t *testing.T) {
	var s Set[int]
	_ = s.Has(123)

	// We expect Has to be read-only: it must NOT initialize the map.
	// If this fails, your Has() has side-effects.
	if s.m != nil {
		t.Fatalf("expected zero value Has() not to initialize map; got s.m != nil")
	}
}
