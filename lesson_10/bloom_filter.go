package main

type BloomFilter struct {
	values    int64
	filterLen int8
}

func New(filterLen int8) *BloomFilter {
	// max amount of digits in int64 number -> 19
	const maxFilterLen = 19
	if filterLen > maxFilterLen {
		filterLen = maxFilterLen
	}

	return &BloomFilter{
		filterLen: filterLen,
	}
}

func (bf *BloomFilter) Add(value string) {
	bf.values |= 1 << bf.Hash1(value)
	bf.values |= 1 << bf.Hash2(value)
}

func (bf *BloomFilter) IsValue(value string) bool {
	var mask int64
	mask |= 1 << bf.Hash1(value)
	mask |= 1 << bf.Hash2(value)

	if mask == bf.values&mask {
		return true
	}

	return false
}

const CONST_17 int64 = 17

func (bf *BloomFilter) Hash1(s string) int64 {
	var sum int64 = 0
	for _, char := range s {
		code := int64(char)
		sum += code * CONST_17
	}
	sum %= int64(bf.filterLen)

	return sum
}

const CONST_223 int64 = 223

func (bf *BloomFilter) Hash2(s string) int64 {
	var sum int64 = 0
	for _, char := range s {
		code := int64(char)
		sum += code * CONST_223
	}
	sum %= int64(bf.filterLen)

	return sum
}
