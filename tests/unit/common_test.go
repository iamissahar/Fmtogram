package unit

import (
	"testing"
)

type UnitTest interface {
	startTest(part string, i int, t *testing.T)
}

func check(container []UnitTest, t *testing.T) {
	for i, obj := range container {
		obj.startTest("The First Part:\n%s", i, t)
	}
}

func doubleCheck(container []UnitTest, t *testing.T) {
	if len(container) < 2 {
		t.Fatal("a slice is too short")
	}

	for i, j := 0, 1; j < len(container); i, j = i+1, j+1 {
		container[i].startTest("The Second Part:\n%s", i, t)
		container[j].startTest("The Second Part:\n%s", j, t)
	}
}
