package list

// Element is a node in the doubly linked list.
type Element[T any] struct {
	Value      T
	prev, next *Element[T]
	list       *List[T]
}

// List is a doubly linked list.
// Zero value is ready to use.
type List[T any] struct {
	root Element[T] // sentinel node (root <-> A <-> B <-> C <-> root)
	len  int
}

// New returns an initialized empty list.
func New[T any]() *List[T] {
	l := &List[T]{}
	l.lazyInit()
	return l
}

func (l *List[T]) lazyInit() {
	if l.root.next == nil {
		l.root.next = &l.root
		l.root.prev = &l.root
	}
}

func (l *List[T]) Len() int { return l.len }

func (l *List[T]) Front() *Element[T] {
	l.lazyInit()
	if l.len == 0 { // l.root.next == &l.root
		return nil
	}
	return l.root.next
}

func (l *List[T]) Back() *Element[T] {
	l.lazyInit()
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// PushFront inserts a new element with value v at the front of the list
// and returns the inserted element.
func (l *List[T]) PushFront(v T) *Element[T] {
	l.lazyInit()
	next := l.root.next
	e := &Element[T]{Value: v, prev: &l.root, next: next, list: l}
	l.root.next, next.prev = e, e
	l.len++
	return e
}

// PushBack inserts a new element with value v at the back of the list
// and returns the inserted element.
func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()
	prev := l.root.prev
	e := &Element[T]{Value: v, prev: prev, next: &l.root, list: l}
	l.root.prev, prev.next = e, e
	l.len++
	return e
}

// InsertBefore inserts a new element with value v immediately before mark
// and returns the inserted element. It returns nil if mark is not in the list.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	l.lazyInit()
	if mark == nil || mark.list != l || l.len == 0 {
		return nil
	}

	prev, next := mark.prev, mark
	e := &Element[T]{Value: v, prev: prev, next: next, list: l}
	prev.next, next.prev = e, e
	l.len++
	return e
}

// InsertAfter inserts a new element with value v immediately after mark
// and returns the inserted element. It returns nil if mark is not in the list.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	l.lazyInit()
	if mark == nil || mark.list != l || l.len == 0 {
		return nil
	}

	next, prev := mark.next, mark
	e := &Element[T]{Value: v, prev: prev, next: next, list: l}
	prev.next, next.prev = e, e
	l.len++
	return e
}

// Remove removes e from the list and returns its value and true.
// It returns the zero value and false if e is not in the list.
func (l *List[T]) Remove(e *Element[T]) (v T, ok bool) {
	l.lazyInit()
	if e == nil || e.list != l {
		return v, false
	}

	next, prev := e.next, e.prev
	prev.next, next.prev = next, prev
	l.len--

	// remove all references
	val := e.Value
	e.prev, e.next, e.list = nil, nil, nil
	var zero T
	e.Value = zero

	return val, true
}

// MoveToFront moves element e to the front of the list.
// It returns false if e is not in the list.
func (l *List[T]) MoveToFront(e *Element[T]) bool {
	l.lazyInit()
	if e == nil || e.list != l {
		return false
	}
	// detach
	e.prev.next = e.next
	e.next.prev = e.prev

	// push element to front
	next := l.root.next
	l.root.next, next.prev = e, e
	e.next, e.prev = next, &l.root
	return true
}

// MoveToBack moves element e to the back of the list.
// It returns false if e is not in the list.
func (l *List[T]) MoveToBack(e *Element[T]) bool {
	l.lazyInit()
	if e == nil || e.list != l {
		return false
	}

	// detach e
	e.prev.next = e.next
	e.next.prev = e.prev

	// push element to back
	prev := l.root.prev
	l.root.prev, prev.next = e, e
	e.prev, e.next = prev, &l.root
	return true
}

// Clear removes all elements and releases references.
func (l *List[T]) Clear() {

	l.lazyInit()
	for curr := l.root.next; curr != &l.root; {
		next := curr.next
		l.Remove(curr)
		curr = next
	}

	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
}

// ForEach iterates values from front to back; stop if fn returns false.
func (l *List[T]) ForEach(fn func(v T) bool) {
	l.lazyInit()
	for curr := l.root.next; curr != &l.root; curr = curr.next {
		if !fn(curr.Value) {
			return
		}
	}
}
