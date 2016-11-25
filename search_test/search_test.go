package number_test

import (
	"testing"
	"github.com/jiangyang5157/go/number"
	"fmt"
	"github.com/jiangyang5157/go/sort"
	"github.com/jiangyang5157/go/search"
)

func Test_Sort(t *testing.T) {
	arr := sort.MergeSort(number.RandomArray(10));
	canBefind := arr[3]
	canNotBefind := 222;
	fmt.Printf("Search %v and %v in the sorted arrya %v\n", canBefind, canNotBefind, arr)

	fmt.Printf("Search %v, index %v by LinearSearch\n", canBefind, search.LinearSearch(arr, canBefind))
	fmt.Printf("Search %v, index %v by LinearSearch\n", canNotBefind, search.LinearSearch(arr, canNotBefind))
	fmt.Printf("Search %v, index %v by BinarySearch\n", canBefind, search.BinarySearch(arr, canBefind))
	fmt.Printf("Search %v, index %v by BinarySearch\n", canNotBefind, search.BinarySearch(arr, canNotBefind))
}