package ai

import (
	"gamesys-tictactoe/logic/symbols"
	"gamesys-tictactoe/utils"
)

const NO_SCORE = -100

func (ai Ai) buildSquareScores() []int {
	squareScores := make([]int, ai.State.Size*ai.State.Size)

	diagonalIterator := utils.NewDiagonalIndexIterator(ai.State.Size, ai.State.Squares)
	squareScores = ai.computeSquareScores(diagonalIterator, squareScores)

	backwardDiagonalIterator := utils.NewBackwardDiagonalIndexIterator(ai.State.Size, ai.State.Squares)
	squareScores = ai.computeSquareScores(backwardDiagonalIterator, squareScores)

	for i := 0; i < ai.State.Size; i++ {
		horizontalIterator := utils.NewHorizontalIndexIterator(ai.State.Size, ai.State.Squares, i)
		squareScores = ai.computeSquareScores(horizontalIterator, squareScores)
		verticalIterator := utils.NewVerticalIndexIterator(ai.State.Size, ai.State.Squares, i)
		squareScores = ai.computeSquareScores(verticalIterator, squareScores)
	}

	for i := 0; i < len(squareScores); i++ {
		if squareScores[i] < 0 {
			squareScores[i] = 0
		}
	}

	return squareScores
}

func (ai Ai) computeSquareScores(iter utils.Iterator, squareScores []int) []int {
	var playerMarks int = 0
	iter.Iterate(func(index int, value string) {
		if value != symbols.BLANK {
			squareScores[index] = NO_SCORE
		}
		if value == ai.PlayerSymbol {
			playerMarks += 1
		}
	})
	if playerMarks == 0 {
		iter.Iterate(func(index int, value string) {
			squareScores[index] += 1
		})
	}
	return squareScores
}
