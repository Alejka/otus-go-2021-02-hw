package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	first *ListItem
	last  *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.first
}

func (l list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
	}

	if l.first == nil { // empty list
		l.last = i
	} else { // link with previous first element
		l.first.Prev = i
		i.Next = l.first
	}

	l.first = i
	l.len++

	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{
		Value: v,
	}

	if l.last == nil { // empty list
		l.first = i
	} else { // link with previous last element
		l.last.Next = i
		i.Prev = l.last
	}

	l.last = i
	l.len++

	return l.last
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil { // first element
		l.first = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil { // last element
		l.last = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.first == i {
		return
	}

	i.Prev.Next = i.Next
	if i.Next == nil { // last element
		l.last = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = l.first
	l.first.Prev = i
	l.first = i
}

func NewList() List {
	return new(list)
}
