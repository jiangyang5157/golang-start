package number_test

import (
	"testing"
	"github.com/jiangyang5157/go/number"
)

func Test_RandomArray(t *testing.T) {
	arr := number.RandomArray(10)
	if arr == nil {
		t.Error("RandomArray(10) returns nil ")
	}
	if len(arr) != 10 {
		t.Error("Lenth of RandomArray(10) is wrong")
	}

	arr = number.RandomArray(0)
	if arr == nil {
		t.Error("RandomArray(0) returns nil ")
	}
	if len(arr) != 0 {
		t.Error("Lenth of RandomArray(0) is wrong")
	}

	arr = number.RandomArray(1)
	if arr[0] != 0 {
		t.Error("RandomArray(1)[0] is wrong ")
	}
}