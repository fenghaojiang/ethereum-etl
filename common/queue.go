package common

import (
	"container/list"
	"sync"
)

type Queue[T any] struct {
	mu    sync.RWMutex
	queue *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		mu:    sync.RWMutex{},
		queue: list.New(),
	}
}

func (q *Queue[T]) PopFront() T {
	q.mu.Lock()
	front := q.queue.Remove(q.queue.Front())
	q.mu.Unlock()
	return front.(T)
}

func (q *Queue[T]) PushBack(elem T) {
	q.mu.Lock()
	q.queue.PushBack(elem)
	q.mu.Unlock()
}

func (q *Queue[T]) Front() T {
	var empty T
	q.mu.RLock()
	elem := q.queue.Front()
	q.mu.RUnlock()
	if elem == nil {
		return empty
	}
	return elem.Value.(T)
}
