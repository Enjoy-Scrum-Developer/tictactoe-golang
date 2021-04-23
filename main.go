package main

import (
	"fmt"
	"gamesys-tictactoe/logic/ai"
	"gamesys-tictactoe/logic/state"
	"gamesys-tictactoe/logic/symbols"
	"gamesys-tictactoe/presentation"
	"strconv"
	"strings"
)

const AI_LEVEL = 0
const BOARD_SIZE = 3

func main() {
	state := state.New(BOARD_SIZE)
	presentation.ShowBoard(*state)
	var winner string
	for {
		askPlayerToMove(*state)
		winner = state.FindWinner()
		if winner != "" || state.AllSquaresMarked() {
			break
		}
		askComputerToMove(*state)
		winner = state.FindWinner()
		if winner != "" || state.AllSquaresMarked() {
			break
		}
	}
}

func askPlayerToMove(state state.GameState) {
	text := presentation.Prompt(fmt.Sprintf("Enter a number (from 1 to %d): ", state.Size*state.Size))
	no, _ := strconv.Atoi(strings.TrimSpace(text))
	state.MarkSquare(no, symbols.MARK_X)
	presentation.ShowBoard(state)
}

func askComputerToMove(state state.GameState) {
	ai := ai.New(state, ai.NewDefaultRandomPicker())
	no := ai.NextMove(0)
	state.MarkSquare(no, symbols.MARK_O)
	presentation.ShowBoard(state)
}
