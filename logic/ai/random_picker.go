package ai

import (
	"gamesys-tictactoe/utils"
	"math/rand"
)

type DefaultRandomPicker struct {
}

func NewDefaultRandomPicker() *DefaultRandomPicker {
	return new(DefaultRandomPicker)
}

func (p DefaultRandomPicker) Pick(list utils.IntegerLinkedList) int {
	randomIndex := rand.Intn(rand.Intn(list.Length()))
	value, _ := list.Get(randomIndex)
	return value
}
