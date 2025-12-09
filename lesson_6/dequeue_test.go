package main

import (
	"testing"
)

func TestAddTail(t *testing.T) {
	q := NewDequeue[int]()

	q.AddTail(10)
	q.AddTail(20)
	q.AddTail(30)

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	if q.GetTail() != 30 {
		t.Fatalf("expected tail 30, got %d", q.GetTail())
	}

	if q.GetHead() != 10 {
		t.Fatalf("expected head 10, got %d", q.GetHead())
	}
}

func TestAddHead(t *testing.T) {
	q := NewDequeue[int]()

	q.AddHead(10)
	q.AddHead(20)
	q.AddHead(30)

	if q.Size() != 3 {
		t.Fatalf("expected size 3, got %d", q.Size())
	}

	if q.GetHead() != 30 {
		t.Fatalf("expected head 30, got %d", q.GetHead())
	}

	if q.GetTail() != 10 {
		t.Fatalf("expected tail 10, got %d", q.GetTail())
	}
}

func TestRemoveHead_OK(t *testing.T) {
	q := NewDequeue[int]()
	q.AddTail(10)
	q.AddTail(20)
	q.AddTail(30)

	q.RemoveHead()

	if q.Size() != 2 {
		t.Fatalf("expected size 2, got %d", q.Size())
	}

	if q.GetRemoveHeadStatus() != REMOVE_HEAD_OK {
		t.Fatalf("expected status REMOVE_HEAD_OK")
	}

	if q.GetHead() != 20 {
		t.Fatalf("expected new head 20, got %d", q.GetHead())
	}
}

func TestRemoveHead_Empty(t *testing.T) {
	q := NewDequeue[int]()

	q.RemoveHead()

	if q.GetRemoveHeadStatus() != REMOVE_HEAD_ERR {
		t.Fatalf("expected REMOVE_HEAD_ERR on empty dequeue")
	}
}

func TestRemoveTail_OK(t *testing.T) {
	q := NewDequeue[int]()
	q.AddTail(10)
	q.AddTail(20)
	q.AddTail(30)

	q.RemoveTail()

	if q.Size() != 2 {
		t.Fatalf("expected size 2, got %d", q.Size())
	}

	if q.GetRemoveTailStatus() != REMOVE_TAIL_OK {
		t.Fatalf("expected status REMOVE_TAIL_OK")
	}

	if q.GetTail() != 20 {
		t.Fatalf("expected new tail 20, got %d", q.GetTail())
	}
}

func TestRemoveTail_Empty(t *testing.T) {
	q := NewDequeue[int]()

	q.RemoveTail()

	if q.GetRemoveTailStatus() != REMOVE_TAIL_ERR {
		t.Fatalf("expected REMOVE_TAIL_ERR on empty dequeue")
	}
}

func TestGetHead_OK(t *testing.T) {
	q := NewDequeue[int]()
	q.AddTail(5)

	head := q.GetHead()
	if head != 5 {
		t.Fatalf("expected head 5, got %d", head)
	}

	if q.GetGetHeadStatus() != GET_HEAD_NIL {
		t.Fatalf("expected GET_HEAD_NIL (status set only on error)")
	}
}

func TestGetHead_Empty(t *testing.T) {
	q := NewDequeue[int]()

	val := q.GetHead()

	if val != 0 {
		t.Fatalf("expected zero value, got %d", val)
	}

	if q.GetGetHeadStatus() != GET_HEAD_ERR {
		t.Fatalf("expected GET_HEAD_ERR on empty")
	}
}

func TestGetTail_OK(t *testing.T) {
	q := NewDequeue[int]()
	q.AddTail(100)

	tail := q.GetTail()
	if tail != 100 {
		t.Fatalf("expected tail 100, got %d", tail)
	}

	if q.GetGetTailStatus() != GET_TAIL_NIL {
		t.Fatalf("expected GET_TAIL_NIL (status set only on error)")
	}
}

func TestGetTail_Empty(t *testing.T) {
	q := NewDequeue[int]()

	val := q.GetTail()

	if val != 0 {
		t.Fatalf("expected zero value, got %d", val)
	}

	if q.GetGetTailStatus() != GET_TAIL_ERR {
		t.Fatalf("expected GET_TAIL_ERR on empty")
	}
}

func TestAddRemoveCombination(t *testing.T) {
	q := NewDequeue[int]()

	q.AddTail(1)   // [1]
	q.AddTail(2)   // [1,2]
	q.AddHead(0)   // [0,1,2]
	q.RemoveTail() // [0,1]
	q.RemoveHead() // [1]
	q.AddHead(5)   // [5,1]

	if q.Size() != 2 {
		t.Fatalf("expected size 2, got %d", q.Size())
	}

	if q.GetHead() != 5 {
		t.Fatalf("expected head 5, got %d", q.GetHead())
	}

	if q.GetTail() != 1 {
		t.Fatalf("expected tail 1, got %d", q.GetTail())
	}
}
