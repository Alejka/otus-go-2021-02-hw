package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if listItem, ok := l.items[key]; ok {
		listItem.Value.(*cacheItem).value = value
		l.queue.MoveToFront(listItem)
		return true
	}

	if l.queue.Len() == l.capacity { // cache is full, we should remove the oldest element
		l.purgeOldestItem()
	}

	cacheItem := &cacheItem{key: string(key), value: value}
	listItem := l.queue.PushFront(cacheItem)
	l.items[key] = listItem

	return false
}

func (l *lruCache) Get(key Key) (value interface{}, ok bool) {
	if listItem, ok := l.items[key]; ok {
		l.queue.MoveToFront(listItem)
		return listItem.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func (l *lruCache) purgeOldestItem() {
	oldestListItem := l.queue.Back()
	l.queue.Remove(oldestListItem)
	oldestKey := oldestListItem.Value.(*cacheItem).key
	delete(l.items, Key(oldestKey))
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
