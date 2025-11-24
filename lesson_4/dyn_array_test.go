package main

import "testing"

func TestAdd(t *testing.T) {
	mapStatus := map[Status]string{
		ADD_NIL:          "ADD_NIL",
		ADD_OK:           "ADD_OK",
		ADD_OUT_OF_RANGE: "ADD_OUT_OF_RANGE",
	}

	tests := []struct {
		name              string
		dynArray          *DynArray[any]
		indx              []int
		value             []any
		expectedAddStatus Status
		expectedLen       int
		expectedCap       int
	}{
		{"Test: normal", DynArrayConstructor[any](6), []int{0}, []any{1}, ADD_OK, 1, 16},
		{"Test: empty array", DynArrayConstructor[any](0), []int{0}, []any{1}, ADD_OK, 1, 16},
		{"Test: full array", &DynArray[any]{
			len:   16,
			cap:   16,
			array: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}, []int{0}, []any{100}, ADD_OK, 17, 32},
		{"Test: add several elements", &DynArray[any]{
			len:   16,
			cap:   16,
			array: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}, []int{0, 5, 1, 18}, []any{100, 200, 300, -5}, ADD_OK, 20, 32},
	}

	for _, test := range tests {
		for i := range test.indx {
			test.dynArray.Add(test.value[i], i)
			addStatus := test.dynArray.GetAddStatus()

			if addStatus != test.expectedAddStatus {
				t.Errorf("invalid addStatus! expected: %s, result: %s", mapStatus[test.expectedAddStatus], mapStatus[addStatus])
			}

			if test.value[i] != test.dynArray.array[i] {
				t.Errorf("invalid value at indx %d! expected: %d, result: %d", i, test.value[i], test.dynArray.array[i])
			}

		}

		len := test.dynArray.Size()
		cap := test.dynArray.Capacity()

		if len != test.expectedLen {
			t.Errorf("invalid len! expected: %d, result: %d", test.expectedLen, len)
		}

		if cap != test.expectedCap {
			t.Errorf("invalid cap! expected: %d, result: %d", test.expectedCap, cap)
		}
	}
}
func TestRemove(t *testing.T) {
	mapStatus := map[Status]string{
		REMOVE_NIL:          "REMOVE_NIL",
		REMOVE_OK:           "REMOVE_OK",
		REMOVE_OUT_OF_RANGE: "REMOVE_OUT_OF_RANGE",
	}

	tests := []struct {
		name                 string
		dynArray             *DynArray[any]
		indx                 []int
		expectedRemoveStatus Status
		expectedLen          int
		expectedCap          int
	}{
		{"Test: normal", &DynArray[any]{
			len:   16,
			cap:   16,
			array: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}, []int{0}, ADD_OK, 15, 16},
	}

	for _, test := range tests {
		for indx := range test.indx {
			test.dynArray.Remove(indx)
			removeStatus := test.dynArray.GetRemoveStatus()

			if removeStatus != test.expectedRemoveStatus {
				t.Errorf("invalid addStatus! expected: %s, result: %s", mapStatus[test.expectedRemoveStatus], mapStatus[removeStatus])
			}

		}

		len := test.dynArray.Size()
		cap := test.dynArray.Capacity()

		if len != test.expectedLen {
			t.Errorf("invalid len! expected: %d, result: %d", test.expectedLen, len)
		}

		if cap != test.expectedCap {
			t.Errorf("invalid cap! expected: %d, result: %d", test.expectedCap, cap)
		}
	}
}
