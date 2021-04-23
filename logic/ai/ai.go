package ai

import (
	"gamesys-tictactoe/logic/state"
	"gamesys-tictactoe/logic/symbols"
	"gamesys-tictactoe/utils"
)

const NO_MOVE = 0

type RandomPicker interface {
	Pick(list utils.IntegerLinkedList) int
}

type Ai struct {
	State        state.GameState
	AiSymbol     string
	PlayerSymbol string
	MovePicker   RandomPicker
}

func New(state state.GameState, movePicker RandomPicker) *Ai {
	ai := new(Ai)
	ai.State = state
	ai.MovePicker = movePicker
	ai.PlayerSymbol = symbols.MARK_X
	ai.AiSymbol = symbols.MARK_O
	return ai
}

func (ai Ai) NextMove(aiLevel int) int {
	nextMove := findWinningMove(ai.State, ai.AiSymbol)
	if nextMove != NO_MOVE {
		return nextMove
	}
	if aiLevel == 0 {
		return ai.MovePicker.Pick(*ai.State.AvailableMoves)
	}
	nextMove = findWinningMove(ai.State, ai.PlayerSymbol)
	if nextMove != NO_MOVE {
		return nextMove
	}
	if aiLevel == 1 {
		return ai.MovePicker.Pick(*ai.State.AvailableMoves)
	}
	return ai.findBestMove(ai.State, ai.AiSymbol, ai.PlayerSymbol)
}

func (ai Ai) findBestMove(state state.GameState, aiSymbol string, playerSymbol string) int {
	squareScores := ai.buildSquareScores()
	_, bestMoves := utils.MaxValues(squareScores)
	if bestMoves.Length() > 1 {
		return ai.MovePicker.Pick(bestMoves) + 1
	}
	bestMove, _ := bestMoves.Get(0)
	return bestMove + 1
}

func findWinningMove(state state.GameState, symbol string) int {
	availableMovesCount := state.AvailableMoves.Length()
	for i := 0; i < availableMovesCount; i++ {
		stateCopy := state.Clone()
		no, _ := stateCopy.AvailableMoves.Get(i)
		if stateCopy.MarkSquare(no, symbol) {
			if stateCopy.FindWinner() == symbol {
				return no
			}
		}
	}
	return NO_MOVE
}
