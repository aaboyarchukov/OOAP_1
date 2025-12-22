# АТД HashTable

Спецификация:

- put(value T) int -- сохранение элемента
- delete(value T) -- удаление элемента по ключу
- exist(value T) bool -- проверка на существование

АТД HashTable:

```go
Status (
	PUT_NIL = iota
	PUT_OK
	PUT_ERR
)
Status (
	DELETE_NIL = iota
	DELETE_OK
	DELETE_ERR
)
Status (
	HASH_NIL = iota
	HASH_OK
	HASH_ERR
)
Status (
	SEEK_SLOT_NIL = iota
	SEEK_SLOT_OK
	SEEK_SLOT_ERR
)

type HashTable struct {
	// конструктор
	HashTable[T any]() HashTable T
	
	// команды
	
	// предусловие: есть своюбодный слот под value
	// постусловие: элемент сохранен в таблице, размер таблицы увеличился на 1
	Put(value T)
	
	// предсуловие: в таблице имеется значение value
	// постусловие: элемент сохранен в таблице, размер таблицы увеличился на 1
	Delete(value T)
	
	// запросы
	
	// постусловие: поиск элемента по таблице завершился
	Exist(value T) bool
	
	// возвращает индекс для элемента
	// предусловие: таблица не пустая
	// постусловие: вычислен индекс для элемента
	Hash(value T) int
	
	// доп. запросы
	GetPutStatus() Status // успешно; элемент уже был в таблице
	GetDeleteStatus() Status // успешно; таблица пуста
	GetHashStatus() Status // успешно; таблица пуста
	
	SeekSlot(value string) int // успешно; таблица пуста
	
	// доп. анонимные запросы
	evacuate() // эвакуирует данные, когда массив переполнился
	needEvacuate() bool // определяет нужна ли эвакуация
	

}
```


Реализация:

```go
package main

import (
	_ "os"
	_ "strconv"
)

type Status = int

const (
	PUT_NIL = iota
	PUT_OK
	PUT_ERR
)
const (
	DELETE_NIL = iota
	DELETE_OK
	DELETE_ERR
)
const (
	HASH_NIL = iota
	HASH_OK
	HASH_ERR
)
const (
	SEEK_SLOT_NIL = iota
	SEEK_SLOT_OK
	SEEK_SLOT_ERR
)

type HashTable struct {
	size      int
	cap       int
	step      int
	slots     []any
	fillSlots []bool

	putStatus      Status
	deleteStatus   Status
	hashStatus     Status
	seekSlotStatus Status
}

const stepsForSeekedSlot = 3
const initCap = 0
const defaultCap = 100
const defaultSize = 100
const invalidIndex = -1
const allocateCoeff = 2

func New() HashTable {
	ht := HashTable{
		cap:            initCap,
		size:           defaultSize,
		step:           stepsForSeekedSlot,
		slots:          nil,
		fillSlots:      nil,
		putStatus:      PUT_NIL,
		deleteStatus:   DELETE_OK,
		hashStatus:     HASH_NIL,
		seekSlotStatus: SEEK_SLOT_NIL,
	}
	ht.slots, ht.fillSlots = make([]any, defaultSize, defaultCap), make([]bool, defaultSize, defaultCap)

	return ht
}

func (ht *HashTable) Hash(value any) int {
	if ht.size == 0 {
		ht.hashStatus = HASH_ERR
		return invalidIndex
	}

	var resultIndx int

	switch valueType := value.(type) {
	case string:
		var sum byte
		for i, item := range valueType {
			sum += byte(item) * byte(i)
		}
		resultIndx = int(sum) % ht.size

	case int:
		resultIndx = valueType % (ht.size - 1)
	}

	ht.hashStatus = HASH_OK
	return resultIndx
}

func (ht *HashTable) Delete(value any) {
	if ht.size == 0 {
		ht.deleteStatus = DELETE_ERR

		return
	}

	indx := ht.SeekSlot(value)
	ht.fillSlots[indx] = false
	ht.slots[indx] = nil

	ht.deleteStatus = DELETE_OK
}

func (ht *HashTable) SeekSlot(value any) int {
	if ht.size == 0 {
		ht.seekSlotStatus = SEEK_SLOT_ERR
		return invalidIndex
	}

	hash := ht.Hash(value)

	if !ht.fillSlots[hash] {
		return hash
	}

	if ht.cap < ht.size {

		resultIndx, indx := hash, hash
		for ht.fillSlots[resultIndx] {
			indx += ht.step
			resultIndx = indx % ht.size
		}
		ht.seekSlotStatus = SEEK_SLOT_OK

		return resultIndx
	}
	ht.seekSlotStatus = SEEK_SLOT_ERR

	return invalidIndex
}

func (ht *HashTable) Exist(value any) bool {
	if ht.SeekSlot(value) != invalidIndex {
		return true
	}

	return false
}

func (ht *HashTable) Put(value any) int {
	if ht.size == 0 {
		ht.putStatus = PUT_ERR
		return invalidIndex
	}

	if ht.needEvacuate() {
		ht.evacuation()
	}

	hash := ht.Hash(value)
	if ht.slots[hash] == value {
		ht.putStatus = PUT_OK
		return hash
	}

	indx := ht.SeekSlot(value)
	if indx != invalidIndex {
		ht.slots[indx] = value
		ht.fillSlots[indx] = true
		ht.cap++
	}
	ht.putStatus = PUT_OK

	return indx
}

func (ht HashTable) GetPutStatus() Status {
	return ht.putStatus
}
func (ht HashTable) GetDeleteStatus() Status {
	return ht.deleteStatus
}
func (ht HashTable) GetHashStatus() Status {
	return ht.hashStatus
}

func (ht *HashTable) needEvacuate() bool {
	return ht.cap == cap(ht.slots)
}

func (ht *HashTable) evacuation() {
	oldSlots := ht.slots

	newLen, newCap := len(oldSlots)*allocateCoeff, cap(oldSlots)*allocateCoeff
	newSlots := make([]any, newLen, newCap)
	newFillSlots := make([]bool, newLen, newCap)

	ht.slots, ht.fillSlots = newSlots, newFillSlots
	ht.cap = initCap
	ht.size = newLen

	for _, slot := range oldSlots {
		ht.Put(slot)
	}
}

```
