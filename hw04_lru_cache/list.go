package hw04lrucache

import (
	"errors"
)

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem) error          // удалить элемент
	MoveToFront(i *ListItem) error     // переместить элемент в начало
}

type ListItem struct {
	Value interface{}
	Next  *ListItem // следующий элемент
	Prev  *ListItem // предыдущий элемент
}

type list struct {
	len   int
	first *ListItem
	last  *ListItem
}

func (l *list) Len() int { // f размер кеша
	return l.len
}

func (l *list) Front() *ListItem { // f первый элемент кеша
	return l.first
}

func (l *list) Back() *ListItem { // f последний элемент кеша
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if l.first == nil {
		item.Next = nil
		l.last = item
	} else {
		item.Next = l.first
		item.Next.Prev = item
	}
	item.Prev = nil
	l.len++
	l.first = item
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	item.Next = nil
	if l.last == nil {
		item.Prev = nil
		l.first = item
	} else {
		item.Prev = l.last
		item.Prev.Next = item
	}
	l.len++
	l.last = item
	return item
}

func (l *list) Remove(i *ListItem) error {
	if err := isItemNil(i); err != nil {
		return err
	}
	switch {
	case i.Prev == nil && i.Next == nil:
		l.first = nil
		l.last = nil
	case i.Prev != nil && i.Next == nil:
		i.Prev.Next = nil
		l.last = i.Prev
	case i.Prev == nil && i.Next != nil:
		i.Next.Prev = nil
		l.first = i.Next
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
	return nil
}

func isItemNil(i *ListItem) error {
	if i == nil {
		return errors.New("nil item")
	}
	return nil
}

func (l *list) MoveToFront(i *ListItem) error {
	if err := l.Remove(i); err != nil {
		return err
	}
	l.PushFront(i.Value)
	return nil
}

func NewList() List {
	return new(list)
}
