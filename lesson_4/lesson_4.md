# АТД DynArray

Спецификация:
Атрибуты:
- длина (len)
- вместимость (cap)
- базовый массив (array)

Методы:
- remove(indx) - удаляет элемент по индексу
- add(value, indx) - добавляет элемент по индексу
- get(indx) - возвращает элемент по индексу
- size() - возвращается количество элементов в динамическом массиве
- capacity() - возвращается количество возможных элементов в списке


```go
type Status int

Status (
	REMOVE_NIL = iota // команда Remove() еще не вызывалась 
	REMOVE_OK // последняя команда Remove() отработала хорошо
	REMOVE_OUT_OF_RANGE // последняя команда Remove() выполнилась с ошибкой 
						//out of range
)

Status (
	GET_NIL = iota // операция Get() еще не вызывалась
	GET_OK // последняя операции Get() отработала корректно
	GET_EMPTY_ARRAY // последняя операция Get() закончилась с ошибкой
					// доступ к пустому массиву - empty array
	GET_OUT_OF_RANGE // последняя операция Get() закончилась с ошибкой
					// out of range
)

Status (
	ADD_NIL = iota // команда Add() еще не вызывалась 
	ADD_OK // последняя команда Add() выполнилась успешно
	ADD_OUT_OF_RANGE // последняя команда Add() выполнилась с ошибкой 
					//out of range
)

const deallocateLimit = 0.25
const zeroLen = 0
const capacityLimit = 16
const reduceValue = 1.5
const raiseValue = 2

type DynArray[T any] struct {
	len int
	cap int
	array []T
	
	addStatus Status
	removeStatus Status
	getStatus Status
	
	// конструктор:
	DynArray[T](cap int) (*DynArray[T])
	
	// команды:
	
	// предусловие: вместимость списка не меньше indx
	// постусловие: элемент под индексом - indx удален
	// при необходимости редуцирует занимаемую память
	Remove(indx int)
	
	// предусловие: вместимость списка не меньше indx
	// постусловие: в динамический массив добавлен новый элемент
	// при необходимости идет реаллокация
	Add[T](value T, indx int)
	
	// запросы:
	
	// предусловие: список не пуст
	// постусловие: вернется элемент под индексом indx
	Get[T](indx int) (T) {
		size := Size()
		if indx > size {
			
		}
		return array[indx]
	}
	
	// постулосвие: вернется актуальный размер массива
	Size() (int) {
		return len
	}
	
	// постусловие: вернется актуальная вместимость массива
	Capacity() (int) {
		return cap
	}
	
	// дополнительные запросы:
	GetRemoveStatus() (Status) // успешно; длина списка меньше indx
	GetAddStatus() (Status) // успешно; длина списка меньше indx
	GetGetStatus() (Status) // успешно; список пустой
	
	// дополнительные приватные методы:
	
	// запросы:
	
	// постусловие: возвращается новый пустой массив меньшего размера
	// в который будут переносится значения старого массива
	deallocateArray[T]() ([]T)
	
	// постусловие: возвращается новый пустой массив большего размера
	// в который будут переносится значения старого массива
	allocateArray() ([]T)
	
	// постусловие: возвращается новый, заполненный старыми значениями, массив
	copyRangeFromTo(from, to []T) ([]T)
	
}
```

Реализация:

```go
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

    len   int

    cap   int

    array []T

  

    addStatus    Status

    removeStatus Status

    getStatus    Status

}

  

func DynArrayConstructor[T any](cap int) *DynArray[T] {

    if cap < 16 {

        cap = 16

    }

  

    len := cap

    return &DynArray[T]{

        len:   zeroLen,

        cap:   cap,

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
```