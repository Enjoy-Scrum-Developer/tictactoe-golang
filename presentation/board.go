package presentation

import (
	"fmt"
	"gamesys-tictactoe/logic/state"
	"strings"
)

func ShowBoard(state state.GameState) {
	border := joinAndEnclose(generateAndFillArray(state.Size, "---"), "-")
	fmt.Println(border)
	for i := 0; i < state.Size; i++ {
		row := joinAndEnclose(state.Squares[i*state.Size:(i+1)*state.Size], " | ")
		fmt.Println(strings.TrimSpace(row))
		fmt.Println(border)
	}
	fmt.Println()
}

func joinAndEnclose(array []string, separator string) string {
	return separator + strings.Join(array, separator) + separator
}

func generateAndFillArray(size int, value string) []string {
	array := make([]string, size)
	for i := 0; i < size; i++ {
		array[i] = value
	}
	return array
}
