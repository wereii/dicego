package dicelib

import "math/rand"

// DiceSet : Set of N M-face dices
//  Basically your hand of N dices.
type DiceSet struct {
	DicesList *[]*Dice
	randgen   *rand.Rand
}

// NewDiceSet : Create DiceSet
func NewDiceSet(diceFaces int, diceAmount int, seed int64) *DiceSet {
	rSrc := rand.NewSource(seed)
	dRand := rand.New(rSrc)

	dices := make([]*Dice, diceAmount)
	dSet := &DiceSet{DicesList: &dices, randgen: dRand}

	for i := 0; i < len(dices); i++ {
		dices[i] = NewDice(diceFaces, dSet.randgen)
	}

	return dSet
}

// Throw : Rolls all underlying dices in the set
func (dset *DiceSet) Throw() {
	for i := 0; i < len(*dset.DicesList); i++ {
		(*dset.DicesList)[i].roll()
	}
}

// GetValues : List of all underyling dice values
func (dset *DiceSet) GetValues() (vals []int) {
	for _, v := range *dset.DicesList {
		vals = append(vals, v.value)
	}
	return
}
