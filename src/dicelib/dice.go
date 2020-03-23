package dicelib

import "math/rand"

// Dice : a n-face dice struct
type Dice struct {
	faces   int
	value   int
	randgen *rand.Rand
}

func (dice *Dice) roll() {
	dice.value = dice.randgen.Intn(dice.faces) + 1
}

// NewDice : Create Dice
func NewDice(faces int, randgen *rand.Rand) *Dice {
	if faces == 0 {
		panic("Zero faced dice does not exist in this universe. Expected n > 0. ")
	}
	return &Dice{faces: faces, randgen: randgen}
}
