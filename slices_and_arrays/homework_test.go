package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue[T int | int8 | int16 | int32 | int64] struct {
	values []T

	front int
	end   int
	size  int
}

func NewCircularQueue[T int | int8 | int16 | int32 | int64](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
		end:    -1,
		front:  -1,
	}
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	if q.front == -1 {
		q.front++
	}
	q.end++
	if q.end == cap(q.values) {
		q.end = q.end - cap(q.values)
	}
	q.values[q.end] = value
	q.size++

	return true
}

// Pop - удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	if q.end == q.front {
		q.end = -1
		q.front = -1
		return true
	}
	q.front++
	q.size--

	return true
}

// Front - получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue[T]) Front() T {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue[T]) Back() T {
	if q.Empty() {
		return -1
	}
	return q.values[q.end]
}

func (q *CircularQueue[T]) Empty() bool {
	return q.size == 0
}

func (q *CircularQueue[T]) Full() bool {
	return q.size == cap(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
