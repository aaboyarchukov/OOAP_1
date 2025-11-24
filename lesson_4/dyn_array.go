package main

import (
	"math"
)

type Status int

const (
	REMOVE_NIL = iota //
	REMOVE_OK
	REMOVE_OUT_OF_RANGE
)

const (
	GET_NIL = iota //
	GET_OK
	GET_EMPTY_ARRAY
	GET_OUT_OF_RANGE
)

const (
	ADD_NIL = iota //
	ADD_OK
	ADD_OUT_OF_RANGE
)

const deallocateLimit = 0.25
const zeroLen = 0
const capacityLimit = 16
const reduceValue = 1.5
const raiseValue = 2

type DynArray[T any] struct {
	len   int
	cap   int
	array []T

	addStatus    Status
	removeStatus Status
	getStatus    Status
}

func DynArrayConstructor[T any](cap int) *DynArray[T] {
	if cap < 16 {
		cap = 16
	}

	len := cap
	return &DynArray[T]{
		len:   zeroLen,
		cap:   cap,
		array: make([]T, len, cap),
	}
}

func (da *DynArray[T]) Remove(indx int) {
	if indx >= da.len {
		da.removeStatus = REMOVE_OUT_OF_RANGE
		return
	}

	// remove
	for i := indx; i < da.len-1; i++ {
		da.array[i] = da.array[i+1]
	}
	da.array = da.array[0 : len(da.array)-1]

	da.len--

	if float64(da.len)/float64(da.cap) < deallocateLimit {
		newArray := da.deallocateArray()
		newArray = da.copyRangeFromTo(da.array, newArray)
		da.array = newArray
		da.len = len(newArray)
		da.cap = cap(newArray)
	}

	da.removeStatus = REMOVE_OK
}

func (da *DynArray[T]) Add(value T, indx int) {

	if indx > da.len {
		da.addStatus = ADD_OUT_OF_RANGE
		return
	}

	if da.len == da.cap {
		newArray := da.allocateArray()
		newArray = da.copyRangeFromTo(da.array, newArray)
		da.array = newArray
		da.len = len(newArray)
		da.cap = cap(newArray)
	}

	// append dummy element
	da.array = append(da.array, value)

	// add
	for i := da.len - 1; i >= indx; i-- {
		da.array[i+1] = da.array[i]
	}

	da.array[indx] = value
	da.len++
	da.addStatus = ADD_OK
}

func (da *DynArray[T]) Get(indx int) T {
	if indx > da.len {
		da.getStatus = GET_OUT_OF_RANGE
	}

	if da.len == 0 {
		da.getStatus = GET_EMPTY_ARRAY
	}

	return da.array[indx]
}

func (da *DynArray[T]) Size() int {
	return da.len
}

func (da *DynArray[T]) Capacity() int {
	return da.cap
}

func (da *DynArray[T]) GetRemoveStatus() Status {
	return da.removeStatus
}
func (da *DynArray[T]) GetAddStatus() Status {
	return da.addStatus
}
func (da *DynArray[T]) GetGetStatus() Status {
	return da.getStatus
}

func (da *DynArray[T]) deallocateArray() []T {
	newCap := int64(math.Max(float64(da.cap)/reduceValue, capacityLimit))
	newArray := make([]T, zeroLen, newCap)

	return newArray
}

func (da *DynArray[T]) allocateArray() []T {
	newCap := da.cap * raiseValue
	newArray := make([]T, zeroLen, newCap)

	return newArray
}

func (da *DynArray[T]) copyRangeFromTo(from, to []T) []T {
	to = append(to, from...)

	return to
}
