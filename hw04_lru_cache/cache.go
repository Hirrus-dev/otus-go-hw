package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) (bool, error) // Добавить значение в кэш по ключу.
	Get(key Key) (interface{}, bool, error)       // Получить значение из кэша по ключу.
	Clear() error                                 // Очистить кэш.
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type Item struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) (bool, error) {
	_, ok := l.items[key]
	if ok {
		l.items[key].Value.(*Item).value = value
		if err := l.queue.MoveToFront(l.items[key]); err != nil {
			return false, err
		}
		return true, nil
	}
	if l.queue.Len() == l.capacity {
		delete(l.items, l.queue.Back().Value.(*Item).key)
		if err := l.queue.Remove(l.queue.Back()); err != nil {
			return false, err
		}
	}
	l.queue.PushFront(&Item{
		key:   key,
		value: value,
	})
	l.items[key] = l.queue.Front()
	return false, nil
}

func (l *lruCache) Get(key Key) (interface{}, bool, error) {
	_, ok := l.items[key]
	if ok {
		if err := l.queue.MoveToFront(l.items[key]); err != nil {
			return nil, false, err
		}
		return l.items[key].Value.(*Item).value, true, nil
	}
	return nil, false, nil
}

func (l *lruCache) Clear() error {
	for i := l.queue.Back(); i != nil; i = i.Next {
		if err := l.queue.Remove(i); err != nil {
			return err
		}
	}
	for k := range l.items {
		delete(l.items, k)
	}
	return nil
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
