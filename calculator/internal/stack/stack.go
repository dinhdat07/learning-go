package stack

// Stack is a LIFO data structure.
type Stack[T any] struct {
	s []T
}

// New creates a stack with optional initial capacity.
func New[T any](cap ...int) *Stack[T] {
	switch len(cap) {
	case 0:
		return &Stack[T]{s: nil}
	case 1:
		if cap[0] < 0 {
			panic("stack: negative capacity")
		}
		return &Stack[T]{s: make([]T, 0, cap[0])}
	default:
		panic("stack: New expects at most 1 capacity argument")
	}
}

// Push adds an element to the top of the stack.
func (st *Stack[T]) Push(v T) {
	st.s = append(st.s, v)
}

// Pop removes and returns the top element.
// ok == false if stack is empty.
func (st *Stack[T]) Pop() (v T, ok bool) {
	n := len(st.s)
	if n == 0 {
		return v, false
	}
	v = st.s[n-1]

	// release reference
	var zero T
	st.s[n-1] = zero

	st.s = st.s[:n-1]
	return v, true
}

// Peek returns the top element without removing it.
func (st *Stack[T]) Peek() (v T, ok bool) {
	n := len(st.s)
	if n == 0 {
		return v, false
	}
	return st.s[n-1], true
}

// Len returns the number of elements.
func (st *Stack[T]) Len() int { return len(st.s) }

// Cap returns the current capacity.
func (st *Stack[T]) Cap() int { return cap(st.s) }

// Clear removes all elements and releases references.
func (st *Stack[T]) Clear() {
	st.s = nil
}
