package utils_test

import (
	"gamesys-tictactoe/utils"
	"reflect"
	"testing"
)

func TestShouldGenerateDiagonalIndexIterator(t *testing.T) {
	squares := []string{
		"X", " ", " ",
		" ", "O", " ",
		" ", " ", "X",
	}
	diagonalIndexIterator := utils.NewDiagonalIndexIterator(3, squares)
	valuesMap := map[int]string{}
	diagonalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		0: "X",
		4: "O",
		8: "X",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Diagonal Index Iterator did not work properly. Returned (%v)", valuesMap)
	}
}

func TestShouldGenerateBackwardDiagonalIndexIterator(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	backwardDiagonalIndexIterator := utils.NewBackwardDiagonalIndexIterator(3, squares)
	valuesMap := map[int]string{}
	backwardDiagonalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		2: "X",
		4: "O",
		6: " ",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Backward Diagonal Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}

func TestShouldGenerateHorizontalIndexIteratorTop(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	horizontalIndexIterator := utils.NewHorizontalIndexIterator(3, squares, 0)
	valuesMap := map[int]string{}
	horizontalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		0: "X",
		1: " ",
		2: "X",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Horizontal Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}

func TestShouldGenerateHorizontalIndexIteratorMiddle(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	horizontalIndexIterator := utils.NewHorizontalIndexIterator(3, squares, 1)
	valuesMap := map[int]string{}
	horizontalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		3: " ",
		4: "O",
		5: " ",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Horizontal Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}

func TestShouldGenerateHorizontalIndexIteratorBottom(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	horizontalIndexIterator := utils.NewHorizontalIndexIterator(3, squares, 2)
	valuesMap := map[int]string{}
	horizontalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		6: " ",
		7: " ",
		8: "X",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Horizontal Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}

func TestShouldGenerateVerticalIndexIteratorLeft(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	verticalIndexIterator := utils.NewVerticalIndexIterator(3, squares, 0)
	valuesMap := map[int]string{}
	verticalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		0: "X",
		3: " ",
		6: " ",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Vertical Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}

func TestShouldGenerateVerticalIndexIteratorCenter(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	verticalIndexIterator := utils.NewVerticalIndexIterator(3, squares, 1)
	valuesMap := map[int]string{}
	verticalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		1: " ",
		4: "O",
		7: " ",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Vertical Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}

func TestShouldGenerateVerticalIndexIteratorRight(t *testing.T) {
	squares := []string{
		"X", " ", "X",
		" ", "O", " ",
		" ", " ", "X",
	}
	verticalIndexIterator := utils.NewVerticalIndexIterator(3, squares, 2)
	valuesMap := map[int]string{}
	verticalIndexIterator.Iterate(func(index int, value string) {
		valuesMap[index] = value
	})
	expected := map[int]string{
		2: "X",
		5: " ",
		8: "X",
	}
	if !reflect.DeepEqual(valuesMap, expected) {
		t.Errorf("Vertical Index Iterator did not work properly. Return (%v)", valuesMap)
	}
}
