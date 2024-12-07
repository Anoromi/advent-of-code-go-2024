package main

import (
	"reflect"
	"testing"
)

func TestFindBefore(t *testing.T) {
	result := unwrapDependencies(map[int][]int{
		0: {4, 3, 2},
		4: {},
		3: {},
		2: {5},
		5: {4},
	})

	if !reflect.DeepEqual(result, map[int]map[int]bool{
		0: {2: true, 3: true, 4: true, 5: true},
		2: {4: true, 5: true},
		3: {},
		4: {},
		5: {4: true},
	}) {
		t.Fatalf(`findBefore, result is %v`, result)
	}
}
