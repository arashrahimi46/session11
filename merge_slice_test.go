package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	list1    []int
	list2    []int
	expected []int
}

func TestMergeIntSlice(t *testing.T) {
	testcases := []testCase{
		{
			list1:    []int{1, 3, 5, 7},
			list2:    []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		//list 2 empty
		{
			list1:    []int{2, 4, 6, 8},
			list2:    []int{},
			expected: []int{2, 4, 6, 8},
		},
	}

	//test all cases
	for _, testcase := range testcases {
		mergedSlice := MergeIntSlice(testcase.list1, testcase.list2)
		assert.ElementsMatchf(t, mergedSlice, testcase.expected, "Its OK")
	}
}

func TestMergeIntSlice_WithEmptyList(t *testing.T) {
	var list1 = []int{1, 2, 3, 5}
	var list2 = []int{}
	expectedSlice := list1
	mergedSlice := MergeIntSlice(list1, list2)
	assert.ElementsMatchf(t, mergedSlice, expectedSlice, "Its OK")

}
