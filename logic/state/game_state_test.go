package state_test

import (
	"gamesys-tictactoe/logic/state"
	"gamesys-tictactoe/logic/symbols"
	"testing"
)

func TestGameStateAllSquaresMarked(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(1, symbols.MARK_X)
	state.MarkSquare(2, symbols.MARK_O)
	state.MarkSquare(3, symbols.MARK_X)
	state.MarkSquare(4, symbols.MARK_O)
	state.MarkSquare(5, symbols.MARK_X)
	state.MarkSquare(6, symbols.MARK_O)
	state.MarkSquare(7, symbols.MARK_O)
	state.MarkSquare(8, symbols.MARK_X)
	state.MarkSquare(9, symbols.MARK_O)
	if !state.AllSquaresMarked() {
		t.Errorf("Not all squares are marked")
	}
	if state.AvailableMoves.Length() > 0 {
		t.Errorf("There are availables moves")
	}
}

func TestGameStateAllSquaresMarkedReturnsFalse(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(5, symbols.MARK_X)
	if state.AllSquaresMarked() {
		t.Errorf("All squares are marked")
	}
}

func TestShouldReturnEmptyStringOnFindWinner(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(1, symbols.MARK_X)
	winner := state.FindWinner()
	if winner != "" {
		t.Errorf("Winner is not empty string but [%s]", winner)
	}
}

func TestShouldReturnXOnFindWinnerDiagonal(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(1, symbols.MARK_X)
	state.MarkSquare(5, symbols.MARK_X)
	state.MarkSquare(9, symbols.MARK_X)
	winner := state.FindWinner()
	if winner != symbols.MARK_X {
		t.Errorf("Winner is not X but [%s]", winner)
	}
}

func TestShouldReturnXOnFindWinnerBackwardDiagonal(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(3, symbols.MARK_X)
	state.MarkSquare(5, symbols.MARK_X)
	state.MarkSquare(7, symbols.MARK_X)
	winner := state.FindWinner()
	if winner != symbols.MARK_X {
		t.Errorf("Winner is not X but [%s]", winner)
	}
}

func TestShouldReturnOOnFindWinnerTopHorizontal(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(1, symbols.MARK_O)
	state.MarkSquare(2, symbols.MARK_O)
	state.MarkSquare(3, symbols.MARK_O)
	winner := state.FindWinner()
	if winner != symbols.MARK_O {
		t.Errorf("Winner is not O but [%s]", winner)
	}
}

func TestShouldReturnOOnFindWinnerCenterHorizontal(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(4, symbols.MARK_O)
	state.MarkSquare(5, symbols.MARK_O)
	state.MarkSquare(6, symbols.MARK_O)
	winner := state.FindWinner()
	if winner != symbols.MARK_O {
		t.Errorf("Winner is not O but [%s]", winner)
	}
}

func TestShouldReturnOOnFindWinnerBottomHorizontal(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(7, symbols.MARK_O)
	state.MarkSquare(8, symbols.MARK_O)
	state.MarkSquare(9, symbols.MARK_O)
	winner := state.FindWinner()
	if winner != symbols.MARK_O {
		t.Errorf("Winner is not O but [%s]", winner)
	}
}

func TestShouldReturnOOnFindWinnerLeftVertical(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(1, symbols.MARK_O)
	state.MarkSquare(4, symbols.MARK_O)
	state.MarkSquare(7, symbols.MARK_O)
	winner := state.FindWinner()
	if winner != symbols.MARK_O {
		t.Errorf("Winner is not O but [%s]", winner)
	}
}

func TestShouldReturnOOnFindWinnerMiddleVertical(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(2, symbols.MARK_O)
	state.MarkSquare(5, symbols.MARK_O)
	state.MarkSquare(8, symbols.MARK_O)
	winner := state.FindWinner()
	if winner != symbols.MARK_O {
		t.Errorf("Winner is not O but [%s]", winner)
	}
}

func TestShouldReturnOOnFindWinnerRightVertical(t *testing.T) {
	state := state.New(3)
	state.MarkSquare(3, symbols.MARK_O)
	state.MarkSquare(6, symbols.MARK_O)
	state.MarkSquare(9, symbols.MARK_O)
	winner := state.FindWinner()
	if winner != symbols.MARK_O {
		t.Errorf("Winner is not O but [%s]", winner)
	}
}
