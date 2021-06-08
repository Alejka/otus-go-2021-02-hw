package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func filledList() List {
	l := NewList()
	l.PushBack(70)
	l.PushBack(80)
	l.PushBack(60)
	l.PushBack(40)
	l.PushBack(10)
	l.PushBack(30)
	l.PushBack(50) // [70, 80, 60, 40, 10, 30, 50]
	return l
}

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("MoveToFront() middle", func(t *testing.T) {
		l := filledList() // [70, 80, 60, 40, 10, 30, 50]

		middle := l.Back().Prev.Prev.Prev // 40
		l.MoveToFront(middle)             // [40, 70, 80, 60, 10, 30, 50]
		require.Equal(t, 7, l.Len())
		require.Equal(t, 40, l.Front().Value)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{40, 70, 80, 60, 10, 30, 50}, elems)
	})

	t.Run("Remove() first last", func(t *testing.T) {
		l := filledList() // [70, 80, 60, 40, 10, 30, 50]

		l.Remove(l.Front()) // [80, 60, 40, 10, 30, 50]
		l.Remove(l.Back())  // [80, 60, 40, 10, 30]
		require.Equal(t, 5, l.Len())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{80, 60, 40, 10, 30}, elems)
	})

	t.Run("Remove() from list with one item", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.Remove(l.Front())
		require.Equal(t, 0, l.Len())
	})
}
