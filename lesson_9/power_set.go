package main

type Status = int

const (
	SET_NIL = iota
	SET_OK
	SET_ERR
)

const (
	REMOVE_NIL = iota
	REMOVE_OK
	REMOVE_ERR
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

type PowerSet struct {
	size      int
	cap       int
	step      int
	values    []any
	fillSlots []bool

	setStatus      Status
	removeStatus   Status
	hashStatus     Status
	seekSlotStatus Status
}

const initSize = 0
const invalidIndex = -1
const allocateCoeff = 2

func New() *PowerSet {
	const stepsForSeekedSlot = 3
	const defaultCap = 100
	const defaultSize = 100

	return &PowerSet{
		size:         initSize,
		cap:          defaultCap,
		step:         stepsForSeekedSlot,
		values:       make([]any, defaultSize, defaultCap),
		fillSlots:    make([]bool, defaultSize, defaultCap),
		setStatus:    SET_NIL,
		removeStatus: REMOVE_NIL,
	}
}

func (ps *PowerSet) Hash(value any) int {
	if ps.cap == 0 {
		ps.hashStatus = HASH_ERR
		return invalidIndex
	}

	var resultIndx int

	switch valueType := value.(type) {
	case string:
		var sum byte
		for i, item := range valueType {
			sum += byte(item) * byte(i)
		}
		resultIndx = int(sum) % ps.cap

	case int:
		resultIndx = valueType % (ps.cap - 1)
	}

	ps.hashStatus = HASH_OK
	return resultIndx
}

func (ps *PowerSet) SeekSlot(value any) int {
	if ps.cap == 0 {
		ps.seekSlotStatus = SEEK_SLOT_ERR
		return invalidIndex
	}

	hash := ps.Hash(value)

	if !ps.fillSlots[hash] {
		return hash
	}

	if ps.values[hash] == value {
		return invalidIndex
	}

	if ps.size < ps.cap {
		resultIndx, indx := hash, hash
		for ps.fillSlots[resultIndx] {
			indx += ps.step
			resultIndx = indx % ps.cap
		}
		ps.seekSlotStatus = SEEK_SLOT_OK

		return resultIndx
	}
	ps.seekSlotStatus = SEEK_SLOT_ERR

	return invalidIndex
}

func (ps *PowerSet) needEvacuate() bool {
	return ps.cap == cap(ps.values)
}

func (ps *PowerSet) evacuation() {
	oldSlots := ps.values

	newLen, newCap := len(oldSlots)*allocateCoeff, cap(oldSlots)*allocateCoeff
	newSlots := make([]any, newLen, newCap)
	newFillSlots := make([]bool, newLen, newCap)

	ps.values, ps.fillSlots = newSlots, newFillSlots
	ps.cap = newCap
	ps.size = initSize

	for _, slot := range oldSlots {
		ps.Set(slot)
	}
}

func (ps *PowerSet) Get(value any) bool {
	return ps.SeekSlot(value) == invalidIndex
}

func (ps *PowerSet) Set(value any) {
	if !ps.Get(value) {
		indx := ps.SeekSlot(value)
		ps.values[indx] = value
		ps.fillSlots[indx] = true
		ps.setStatus = SET_OK
		ps.size++

		return
	}

	ps.setStatus = SET_ERR
}

func (ps *PowerSet) Remove(value any) {
	if !ps.Get(value) {
		ps.removeStatus = REMOVE_ERR
		return
	}

	resultIndx := ps.Hash(value)
	for resultIndx < ps.cap && ps.values[resultIndx] != value {
		resultIndx++
	}

	if resultIndx >= ps.cap {
		ps.removeStatus = REMOVE_ERR
		return
	}

	ps.values[resultIndx] = nil
	ps.fillSlots[resultIndx] = false
	ps.size--

	ps.removeStatus = REMOVE_OK
}

func (ps *PowerSet) Union(powerSetFirst *PowerSet, powerSetSecond *PowerSet) *PowerSet {
	newSet := New()

	for value := range powerSetFirst.values {
		newSet.Set(value)
	}

	for value := range powerSetSecond.values {
		newSet.Set(value)
	}

	return newSet
}

func (ps *PowerSet) Intersection(powerSetFirst *PowerSet, powerSetSecond *PowerSet) *PowerSet {
	newSet := New()
	if powerSetFirst.size < powerSetSecond.size {
		powerSetFirst, powerSetSecond = powerSetSecond, powerSetFirst
	}

	for value := range powerSetSecond.values {
		if powerSetFirst.Get(value) {
			newSet.Set(value)
		}
	}

	return newSet
}
func (ps *PowerSet) Difference(powerSetFirst *PowerSet, powerSetSecond *PowerSet) *PowerSet {
	newSet := New()

	if powerSetSecond.size > powerSetFirst.size {
		powerSetFirst, powerSetSecond = powerSetSecond, powerSetFirst
	}

	for value := range powerSetFirst.values {
		if !powerSetSecond.Get(value) {
			newSet.Set(value)
		}
	}

	return newSet
}

func (ps *PowerSet) Issubset(powerSetFirst *PowerSet, powerSetSecond *PowerSet) bool {
	for value := range powerSetFirst.values {
		if !powerSetSecond.Get(value) {
			return false
		}
	}

	return true
}

func (ps *PowerSet) Equals(powerSetFirst *PowerSet, powerSetSecond *PowerSet) bool {
	if powerSetFirst.size != powerSetSecond.size {
		return false
	}

	return powerSetFirst.Issubset(powerSetFirst, powerSetSecond)
}

func (ps *PowerSet) GetSetStatus() Status {
	return ps.setStatus
}

func (ps *PowerSet) GetRemoveStstua() Status {
	return ps.removeStatus
}
