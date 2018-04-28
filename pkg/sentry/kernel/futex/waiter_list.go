package futex

// List is an intrusive list. Entries can be added to or removed from the list
// in O(1) time and with no additional memory allocations.
//
// The zero value for List is an empty list ready to use.
//
// To iterate over a list (where l is a List):
//      for e := l.Front(); e != nil; e = e.Next() {
// 		// do something with e.
//      }
type waiterList struct {
	head *Waiter
	tail *Waiter
}

// Reset resets list l to the empty state.
func (l *waiterList) Reset() {
	l.head = nil
	l.tail = nil
}

// Empty returns true iff the list is empty.
func (l *waiterList) Empty() bool {
	return l.head == nil
}

// Front returns the first element of list l or nil.
func (l *waiterList) Front() *Waiter {
	return l.head
}

// Back returns the last element of list l or nil.
func (l *waiterList) Back() *Waiter {
	return l.tail
}

// PushFront inserts the element e at the front of list l.
func (l *waiterList) PushFront(e *Waiter) {
	e.SetNext(l.head)
	e.SetPrev(nil)

	if l.head != nil {
		l.head.SetPrev(e)
	} else {
		l.tail = e
	}

	l.head = e
}

// PushBack inserts the element e at the back of list l.
func (l *waiterList) PushBack(e *Waiter) {
	e.SetNext(nil)
	e.SetPrev(l.tail)

	if l.tail != nil {
		l.tail.SetNext(e)
	} else {
		l.head = e
	}

	l.tail = e
}

// PushBackList inserts list m at the end of list l, emptying m.
func (l *waiterList) PushBackList(m *waiterList) {
	if l.head == nil {
		l.head = m.head
		l.tail = m.tail
	} else if m.head != nil {
		l.tail.SetNext(m.head)
		m.head.SetPrev(l.tail)

		l.tail = m.tail
	}

	m.head = nil
	m.tail = nil
}

// InsertAfter inserts e after b.
func (l *waiterList) InsertAfter(b, e *Waiter) {
	a := b.Next()
	e.SetNext(a)
	e.SetPrev(b)
	b.SetNext(e)

	if a != nil {
		a.SetPrev(e)
	} else {
		l.tail = e
	}
}

// InsertBefore inserts e before a.
func (l *waiterList) InsertBefore(a, e *Waiter) {
	b := a.Prev()
	e.SetNext(a)
	e.SetPrev(b)
	a.SetPrev(e)

	if b != nil {
		b.SetNext(e)
	} else {
		l.head = e
	}
}

// Remove removes e from l.
func (l *waiterList) Remove(e *Waiter) {
	prev := e.Prev()
	next := e.Next()

	if prev != nil {
		prev.SetNext(next)
	} else {
		l.head = next
	}

	if next != nil {
		next.SetPrev(prev)
	} else {
		l.tail = prev
	}
}

// Entry is a default implementation of Linker. Users can add anonymous fields
// of this type to their structs to make them automatically implement the
// methods needed by List.
type waiterEntry struct {
	next *Waiter
	prev *Waiter
}

// Next returns the entry that follows e in the list.
func (e *waiterEntry) Next() *Waiter {
	return e.next
}

// Prev returns the entry that precedes e in the list.
func (e *waiterEntry) Prev() *Waiter {
	return e.prev
}

// SetNext assigns 'entry' as the entry that follows e in the list.
func (e *waiterEntry) SetNext(entry *Waiter) {
	e.next = entry
}

// SetPrev assigns 'entry' as the entry that precedes e in the list.
func (e *waiterEntry) SetPrev(entry *Waiter) {
	e.prev = entry
}
