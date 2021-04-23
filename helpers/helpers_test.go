package helpers_test

import (
	"testing"
	"reflect"
	"gamesys-tictactoe/helpers"
)

func TestShouldCreateAndFillArrayWithValuesDynamically(t *testing.T) {
	array := helpers.GenerateArray(9, " ")
	expected := []string {" ", " ", " ", " ", " ", " ", " ", " ", " "}
	if !reflect.DeepEqual(array, expected) {
		t.Errorf("Function GenerateArray() did not create the correct array")
	}
}
