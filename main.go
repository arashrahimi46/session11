package main

import (
	"fmt"
	"net/http"
	"sort"
)

func main() {
	var list1 = []int{1, 2, 3, 5}
	var list2 = []int{2, 3, 4, 5}
	MergeIntSlice(list1, list2)
	testShadowVariable()
	http.ListenAndServe(":5050", nil)
}

func MergeIntSlice(list1, list2 []int) []int {
	mergedSlice := append(list1, list2...)

	sort.Ints(mergedSlice)
	return mergedSlice
}

func testShadowVariable() {
	//example of shadow variable
	testArray := []int{1, 2, 3, 4, 5}
	for i, value := range testArray {
		fmt.Println(value, i)
	}
}
