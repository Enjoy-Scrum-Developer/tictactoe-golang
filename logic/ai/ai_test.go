package ai_test

import (
	"gamesys-tictactoe/logic/ai"
	"gamesys-tictactoe/logic/state"
	"gamesys-tictactoe/logic/symbols"
	"gamesys-tictactoe/utils"
	"testing"
)

type MockRandomPicker struct {
	Picked int
}

func (p MockRandomPicker) Pick(list utils.IntegerLinkedList) int {
	return p.Picked
}

func newMockRandomPicker(picked int) *MockRandomPicker {
	mockRandomPicker := new(MockRandomPicker)
	mockRandomPicker.Picked = picked
	return mockRandomPicker
}

func TestShouldReturnRandomMoveAtLevel0(t *testing.T) {
	state := state.New(3)
	randomPicker := newMockRandomPicker(7)
	ai := ai.New(*state, randomPicker)
	state.MarkSquare(1, symbols.MARK_X)
	state.MarkSquare(2, symbols.MARK_O)
	state.MarkSquare(3, symbols.MARK_X)
	move := ai.NextMove(0)
	if move != 7 {
		t.Errorf("AI did not execute properly on level 0. Got %d", move)
	}
}

func TestShouldReturnWinningMoveAtLevel0(t *testing.T) {
	state := state.New(3)
	randomPicker := newMockRandomPicker(4)
	ai := ai.New(*state, randomPicker)
	state.MarkSquare(5, symbols.MARK_X)
	state.MarkSquare(2, symbols.MARK_O)
	state.MarkSquare(1, symbols.MARK_X)
	state.MarkSquare(9, symbols.MARK_O)
	state.MarkSquare(3, symbols.MARK_X)
	state.MarkSquare(8, symbols.MARK_O)
	state.MarkSquare(4, symbols.MARK_X)
	move := ai.NextMove(0)
	if move != 7 {
		t.Errorf("AI did not execute properly on level 0. Got %d", move)
	}
}

func TestShouldReturnBestMoveAtLevel2(t *testing.T) {
	state := state.New(3)
	randomPicker := newMockRandomPicker(5)
	ai := ai.New(*state, randomPicker)
	move := ai.NextMove(2)
	if move != 5 {
		t.Errorf("AI did not execute properly on level 2. Got %d", move)
	}
}

func TestShouldReturnBestMoveAtLevel2MidGame(t *testing.T) {
	state := state.New(3)
	randomPicker := newMockRandomPicker(6)
	ai := ai.New(*state, randomPicker)
	state.MarkSquare(5, symbols.MARK_X)
	state.MarkSquare(1, symbols.MARK_O)
	move := ai.NextMove(2)
	if move != 7 {
		t.Errorf("AI did not execute properly on level 2. Got %d", move)
	}
}
