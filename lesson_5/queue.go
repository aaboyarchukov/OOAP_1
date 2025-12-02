package main

type Status = int

const (
	REMOVE_NIL = iota //
	REMOVE_OK
	REMOVE_ERR
)

const (
	GET_HEAD_NIL = iota //
	GET_HEAD_OK
	GET_HEAD_ERR
)

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

type Queue[T any] struct {
	head *QNode[T]
	tail *QNode[T]
	size int

	removeStatus  Status
	getHeadStatus Status
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		size:          0,
		removeStatus:  REMOVE_NIL,
		getHeadStatus: GET_HEAD_NIL,
	}
}

func GetQueue[T any](values []T) *Queue[T] {
	var result Queue[T]

	for _, item := range values {
		result.Add(item)
	}

	return &result
}

func (q *Queue[T]) Size() int {
	return q.size
}

// t = O(1)
func (q *Queue[T]) Remove() {
	if q.size == 0 {
		q.removeStatus = REMOVE_ERR
		return
	}

	next := q.head.next
	q.head = next

	if next == nil {
		q.tail = next
	}

	q.size--
	q.removeStatus = REMOVE_OK
}

// t = O(1)
func (q *Queue[T]) Add(itm T) {
	node := &QNode[T]{
		value: itm,
		next:  nil,
	}

	if q.head == nil {
		q.head = node
	} else {
		q.tail.next = node
	}

	q.tail = node
	q.size++
}

func (q *Queue[T]) GetHead() T {
	var zeroValue T

	if q.size == 0 {
		q.getHeadStatus = GET_HEAD_ERR
		return zeroValue
	}

	return q.head.value
}

func (q *Queue[T]) GetRemoveStatus() Status {
	return q.removeStatus
}

func (q *Queue[T]) GetGetHeadStatus() Status {
	return q.getHeadStatus
}

func EqualQueue[T comparable](q1 *Queue[T], q2 *Queue[T]) bool {
	if q1.size != q2.size {
		return false
	}

	if q1 == nil && q1 == q2 {
		return true
	}

	tempNode1, tempNode2 := q1.head, q2.head

	for tempNode1 != nil {
		if tempNode1.value != tempNode2.value {
			return false
		}
		tempNode1 = tempNode1.next
		tempNode2 = tempNode2.next
	}
	return true
}
