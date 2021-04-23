package state

import (
	"gamesys-tictactoe/helpers"
	"gamesys-tictactoe/logic/symbols"
	"gamesys-tictactoe/utils"
)

const HORIZONTAL_PLUS_VERTICAL = 2
const DIAGONALS = 2

type GameState struct {
	Size           int
	Squares        []string
	AvailableMoves *utils.IntegerLinkedList
	Started        bool
}

func New(size int) *GameState {
	s := new(GameState)
	s.Size = size
	s.Squares = helpers.GenerateArray(size*size, symbols.BLANK)
	s.AvailableMoves = utils.NewIntegerLinkedListRange(1, size*size)
	return s
}

func (s *GameState) MarkSquare(no int, symbol string) bool {
	s.Started = true
	index := no - 1
	if s.Squares[index] == symbols.BLANK {
		s.Squares[index] = symbol
		s.AvailableMoves.RemoveFirst(no)
		return true
	}
	return false
}

func (s GameState) AllSquaresMarked() bool {
	return s.AvailableMoves.Length() == 0
}

func (s GameState) FindWinner() string {
	finderFuncs := make([]func() string, 0, s.Size*HORIZONTAL_PLUS_VERTICAL+DIAGONALS)
	finderFuncs = append(finderFuncs, s.findWinnerOnDiagonal)
	finderFuncs = append(finderFuncs, s.findWinnerOnBackwardDiagonal)
	for i := 0; i < s.Size; i++ {
		finderFuncs = append(finderFuncs, s.findWinnerOnHorizontal(i))
		finderFuncs = append(finderFuncs, s.findWinnerOnVertical(i))
	}
	return s.findWinnerOnPatterns(finderFuncs)
}

func (s GameState) findWinnerOnPatterns(finderFuncs []func() string) string {
	var winner string
	for _, finderFunc := range finderFuncs {
		winner = finderFunc()
		if winner != "" {
			return winner
		}
	}
	return ""
}

func (s GameState) findWinnerOnVertical(column int) func() string {
	return func() string {
		return s.findWinnerOnPattern(utils.NewVerticalIndexIterator(s.Size, s.Squares, column))
	}
}

func (s GameState) findWinnerOnHorizontal(row int) func() string {
	return func() string {
		return s.findWinnerOnPattern(utils.NewHorizontalIndexIterator(s.Size, s.Squares, row))
	}
}

func (s GameState) findWinnerOnBackwardDiagonal() string {
	return s.findWinnerOnPattern(utils.NewBackwardDiagonalIndexIterator(s.Size, s.Squares))
}

func (s GameState) findWinnerOnDiagonal() string {
	return s.findWinnerOnPattern(utils.NewDiagonalIndexIterator(s.Size, s.Squares))
}

func (s GameState) findWinnerOnPattern(iter utils.Iterator) string {
	tally := s.tallyMarkedSquares(iter)
	for symbol, value := range tally {
		if value == 3 {
			return symbol
		}
	}
	return ""
}

func (s GameState) tallyMarkedSquares(iter utils.Iterator) map[string]int {
	tally := map[string]int{}
	iter.Iterate(func(index int, value string) {
		count, found := tally[value]
		if !found {
			count = 0
		}
		count++
		if value != symbols.BLANK {
			tally[value] = count
		}
	})
	return tally
}

func (s GameState) Clone() *GameState {
	state := New(s.Size)
	state.Squares = make([]string, len(s.Squares))
	copy(state.Squares, s.Squares)
	state.AvailableMoves = s.AvailableMoves.Clone()
	state.Started = s.Started
	return state
}
