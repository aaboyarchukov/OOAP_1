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
	GET_NIL = iota
	GET_OK
	GET_ERR
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

type NativeDictionary struct {
	size      int
	cap       int
	step      int
	slots     []string
	fillSlots []bool
	values    []any

	putStatus      Status
	getStatus      Status
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

func New() *NativeDictionary {
	return &NativeDictionary{
		cap:            initCap,
		size:           defaultSize,
		step:           stepsForSeekedSlot,
		slots:          make([]string, defaultSize, defaultCap),
		fillSlots:      make([]bool, defaultSize, defaultCap),
		putStatus:      PUT_NIL,
		deleteStatus:   DELETE_OK,
		hashStatus:     HASH_NIL,
		seekSlotStatus: SEEK_SLOT_NIL,
	}
}

func (nd *NativeDictionary) Hash(key string) int {
	if nd.size == 0 {
		nd.hashStatus = HASH_ERR
		return invalidIndex
	}

	var resultIndx int
	var sum byte
	for i, item := range key {
		sum += byte(item) * byte(i)
	}
	resultIndx = int(sum) % nd.size

	nd.hashStatus = HASH_OK

	return resultIndx
}

func (nd *NativeDictionary) Delete(key string) {
	if nd.size == 0 {
		nd.deleteStatus = DELETE_ERR

		return
	}

	indx := nd.SeekSlot(key)
	nd.fillSlots[indx] = false
	nd.values[indx] = nil
	nd.slots[indx] = ""

	nd.deleteStatus = DELETE_OK
}

func (nd *NativeDictionary) SeekSlot(key string) int {
	if nd.size == 0 {
		nd.seekSlotStatus = SEEK_SLOT_ERR
		return invalidIndex
	}

	hash := nd.Hash(key)

	if !nd.fillSlots[hash] {
		return hash
	}

	if nd.cap < nd.size {

		resultIndx, indx := hash, hash
		for nd.slots[resultIndx] == key && nd.fillSlots[resultIndx] {
			indx += nd.step
			resultIndx = indx % nd.size
		}
		nd.seekSlotStatus = SEEK_SLOT_OK

		return resultIndx
	}
	nd.seekSlotStatus = SEEK_SLOT_ERR

	return invalidIndex
}

func (nd *NativeDictionary) Exist(key string) bool {
	if nd.SeekSlot(key) != invalidIndex {
		return true
	}

	return false
}

func (nd *NativeDictionary) Put(key string, value any) int {
	if nd.size == 0 {
		nd.putStatus = PUT_ERR
		return invalidIndex
	}

	if nd.needEvacuate() {
		nd.evacuation()
	}

	hash := nd.Hash(key)
	if nd.values[hash] == value {
		nd.putStatus = PUT_OK
		return hash
	}

	indx := nd.SeekSlot(key)
	if indx != invalidIndex {
		nd.values[indx] = value
		nd.slots[indx] = key
		nd.fillSlots[indx] = true
		nd.cap++
	}
	nd.putStatus = PUT_OK

	return indx
}

func (nd *NativeDictionary) Get(key string) any {
	var result any
	if nd.size == 0 {
		nd.getStatus = GET_ERR
		return result
	}

	return nd.values[nd.SeekSlot(key)]
}

func (nd *NativeDictionary) Size() int {
	return nd.size
}

func (nd *NativeDictionary) GetPutStatus() Status {
	return nd.putStatus
}

func (nd *NativeDictionary) GetGetStatus() Status {
	return nd.getStatus
}

func (nd *NativeDictionary) GetDeleteStatus() Status {
	return nd.deleteStatus
}

func (nd *NativeDictionary) GetHashStatus() Status {
	return nd.hashStatus
}

func (nd *NativeDictionary) needEvacuate() bool {
	return nd.cap == cap(nd.slots)
}

func (nd *NativeDictionary) evacuation() {
	oldSlots, oldValues := nd.slots, nd.values

	newLen, newCap := len(oldSlots)*allocateCoeff, cap(oldSlots)*allocateCoeff
	newSlots := make([]string, newLen, newCap)
	newFillSlots := make([]bool, newLen, newCap)
	newValues := make([]any, newLen, newCap)

	nd.slots, nd.fillSlots, nd.values = newSlots, newFillSlots, newValues
	nd.cap = initCap
	nd.size = newLen

	for indx, slot := range oldSlots {
		nd.Put(slot, oldValues[indx])
	}
}
