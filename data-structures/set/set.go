package set

// Set is a hash set backed by map[T]struct{}.
// Zero value is ready to use.
type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](cap ...int) *Set[T] {
	switch len(cap) {
	case 0:
		return &Set[T]{m: nil}
	case 1:
		if cap[0] < 0 {
			panic("set: negative capacity")
		}
		return &Set[T]{m: make(map[T]struct{}, cap[0])}
	default:
		panic("set: New expects at most 1 capacity argument")
	}
}

// Add inserts v. Returns true if v was newly added.
func (s *Set[T]) Add(v T) bool {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	if _, ok := s.m[v]; ok {
		return false
	}
	s.m[v] = struct{}{}
	return true
}

// AddMany inserts many values. Returns number of newly added items.
func (s *Set[T]) AddMany(vs ...T) int {
	added := 0
	for _, v := range vs {
		if s.Add(v) {
			added++
		}
	}
	return added
}

// Remove deletes v. Returns true if v existed.
func (s *Set[T]) Remove(v T) bool {
	if s.m == nil {
		return false
	}
	if _, ok := s.m[v]; !ok {
		return false
	}
	delete(s.m, v)
	return true
}

// Has reports whether v exists in the set.
func (s *Set[T]) Has(v T) bool {
	if s.m == nil {
		return false
	}
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int { return len(s.m) }

// Clear removes all elements and releases references.
func (s *Set[T]) Clear() { s.m = nil }

// Values returns a snapshot slice of elements (order not guaranteed).
func (s *Set[T]) Values() []T {
	if len(s.m) == 0 {
		return nil
	}
	out := make([]T, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}
