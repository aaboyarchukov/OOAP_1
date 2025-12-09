package main

type Status int

const (
	REMOVE_HEAD_NIL = iota
	REMOVE_HEAD_OK
	REMOVE_HEAD_ERR
)

const (
	REMOVE_TAIL_NIL = iota
	REMOVE_TAIL_OK
	REMOVE_TAIL_ERR
)

const (
	GET_HEAD_NIL = iota
	GET_HEAD_OK
	GET_HEAD_ERR
)

const (
	GET_TAIL_NIL = iota
	GET_TAIL_OK
	GET_TAIL_ERR
)

type Dequeue[T any] struct {
	dequeue []T
	size    int

	removeHeadStatus Status
	getHeadStatus    Status

	removeTailStatus Status
	getTailStatus    Status
}

type ParentQueue[T any] interface {
	Size() int
	AddTail(value T)
	RemoveHead()
	GetHead() T
}

type DequeueContract[T any] interface {
	ParentQueue[T]
	AddFront(value T)
	RemoveTail()
	GetTail()
}

func NewDequeue[T any]() Dequeue[T] {
	return Dequeue[T]{}
}

func (q *Dequeue[T]) Size() int {
	return q.size
}

func (q *Dequeue[T]) RemoveHead() {
	var secondElementIndx = 1
	if q.size == 0 {
		q.removeHeadStatus = REMOVE_HEAD_ERR
		return
	}

	if q.size == 1 {
		q.dequeue = make([]T, 0)
	} else {
		q.dequeue = q.dequeue[secondElementIndx:]

	}

	q.size--
	q.removeHeadStatus = REMOVE_HEAD_OK
}

func (q *Dequeue[T]) RemoveTail() {
	var lastElementIndx = q.size - 1

	if q.size == 0 {
		q.removeTailStatus = REMOVE_TAIL_ERR
		return
	}

	if q.size == 1 {
		q.dequeue = make([]T, 0)
	} else {
		q.dequeue = q.dequeue[:lastElementIndx]

	}

	q.size--
	q.removeTailStatus = REMOVE_TAIL_OK
}

func (q *Dequeue[T]) AddTail(itm T) {
	q.dequeue = append(q.dequeue, itm)
	q.size++
}

func (q *Dequeue[T]) AddHead(itm T) {
	q.dequeue = append([]T{itm}, q.dequeue...)
	q.size++
}

func (q *Dequeue[T]) GetHead() T {
	var firstElementIndx = 0
	var zeroValue T

	if q.size == 0 {
		q.getHeadStatus = GET_HEAD_ERR
		return zeroValue
	}

	return q.dequeue[firstElementIndx]
}

func (q *Dequeue[T]) GetTail() T {
	var lastElementIndx = q.size - 1
	var zeroValue T

	if q.size == 0 {
		q.getTailStatus = GET_TAIL_ERR
		return zeroValue
	}

	return q.dequeue[lastElementIndx]
}

func (q *Dequeue[T]) GetRemoveHeadStatus() Status {
	return q.removeHeadStatus
}

func (q *Dequeue[T]) GetGetHeadStatus() Status {
	return q.getHeadStatus
}

func (q *Dequeue[T]) GetRemoveTailStatus() Status {
	return q.removeTailStatus
}

func (q *Dequeue[T]) GetGetTailStatus() Status {
	return q.getTailStatus
}
