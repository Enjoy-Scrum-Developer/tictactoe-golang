package utils_test

import (
	"testing"
	"gamesys-tictactoe/utils"
)

func TestShouldBeAbleToAddOneElementAndGetElementFromIntegerLinkedList(t *testing.T) {
	list := utils.NewIntegerLinkedList()
	list.Add(1)
	element, _ := list.Get(0)
	if element != 1 {
		t.Errorf("Either Add or Get method didn't work")
	}
}

func TestShouldBeAbleToAddMultipleElementsAndGetAnyElementFromIntegerLinkedList(t *testing.T) {
	list := utils.NewIntegerLinkedList()
	list.Add(1)
	list.Add(5)
	list.Add(3)
	element, _ := list.Get(2)
	if element != 3 {
		t.Errorf("Either Add or Get method didn't work")
	}
	element, _ = list.Get(0)
	if element != 1 {
		t.Errorf("Either Add or Get method didn't work")
	}
	element, _ = list.Get(1)
	if element != 5 {
		t.Errorf("Either Add or Get method didn't work")
	}
}

func TestShouldBeAbleToRemoveFirstElementFromIntegerLinkedListAndShiftIndices(t *testing.T) {
	list := utils.NewIntegerLinkedList()
	list.Add(1)
	list.Add(5)
	list.Add(8)
	list.Add(10)
	list.Add(2)
	list.RemoveFirst(1)
	element, _ := list.Get(0)
	if element != 5 {
		t.Errorf("Either Remove or Get method didn't work. Get returned %d", element)
	}
	length := list.Length()
	if length != 4 {
		t.Errorf("Either Remove or Length method didn't work. Length returned %d", length)
	}
}

func TestShouldBeAbleToRemoveSecondElementFromIntegerLinkedListAndShiftIndices(t *testing.T) {
	list := utils.NewIntegerLinkedList()
	list.Add(1)
	list.Add(5)
	list.Add(8)
	list.Add(10)
	list.Add(2)
	list.RemoveFirst(5)
	element, _ := list.Get(1)
	if element != 8 {
		t.Errorf("Either Remove or Get method didn't work. Get returned %d", element)
	}
	length := list.Length()
	if length != 4 {
		t.Errorf("Either Remove or Length method didn't work. Length returned %d", length)
	}
}

func TestShouldBeAbleToRemoveLastElementFromIntegerLinkedListAndReturnErrorWhenTryingToGetIt(t *testing.T) {
	list := utils.NewIntegerLinkedList()
	list.Add(1)
	list.Add(5)
	list.Add(8)
	list.Add(10)
	list.Add(2)
	list.RemoveFirst(2)
	_, err := list.Get(4)
	if err == nil {
		t.Errorf("Method call to Get did not return an error")
	}
	length := list.Length()
	if length != 4 {
		t.Errorf("Either Remove or Length method didn't work. Length returned %d", length)
	}
}

func TestShouldBeAbleToCreateIntegerRangeOverLinkedList(t *testing.T) {
	list := utils.NewIntegerLinkedListRange(1, 20)
	value, _ := list.Get(19)
	if value != 20 {
		t.Errorf("NewIntegerLinkedListRange or Get method did not work properly. Returned (%d)", value)
	}
}
