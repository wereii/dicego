package dicelib

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewDiceSet(t *testing.T) {
	dicesAmount := 6
	diceFaces := 6

	expectedType := reflect.TypeOf(&(DiceSet{}))
	got := NewDiceSet(diceFaces, dicesAmount, defaultRandSeed)
	gotType := reflect.TypeOf(got)

	// Check return type
	if gotType != expectedType {
		t.Errorf("Expected type \"%v\", got \"%v\".", gotType, expectedType)
	}

	// Check internal DiceList length
	if n := len(*got.DicesList); n != dicesAmount {
		t.Errorf("Expected DiceList of length %d, was %d instead.", dicesAmount, n)
	}

	// Check amount of faces within the DiceList dices
	for i, dice := range *got.DicesList {
		if dice.faces != diceFaces {
			t.Errorf("Expected Dice with %d faces, got dice with %d, index %d.", diceFaces, dice.faces, i)
		}
	}
}

func TestGetValues(t *testing.T) {
	const dicesAmount int = 6

	dSet := NewDiceSet(6, dicesAmount, defaultRandSeed)

	expectType := reflect.TypeOf([]int{})
	got := dSet.GetValues()
	gotType := reflect.TypeOf(got)
	if gotType != expectType {
		t.Errorf("Return type mismatch, expected %v but got %v.", expectType, gotType)
	}

	if n := len(got); n != dicesAmount {
		t.Errorf("Expected return list of length %d, got %d.", dicesAmount, n)
	}
}

func TestThrow(t *testing.T) {
	const dicesAmount int = 6

	expRand := rand.New(rand.NewSource(defaultRandSeed))

	// Make expected "random" list with constant seed (see main_test.go)
	expected := [dicesAmount]int{}
	d1 := &Dice{faces: 6, randgen: expRand}
	for i := 0; i < dicesAmount; i++ {
		d1.roll()
		expected[i] = d1.value
	}

	dSet := NewDiceSet(6, 6, defaultRandSeed)
	dSet.Throw()

	didError := false

	got := dSet.GetValues()
	for i, val := range got {
		if exp := expected[i]; exp != val {
			t.Errorf("Unexpected value, expected %d, got %d, on index %d.", exp, val, i)
			didError = true
		}
	}

	if didError {
		t.Errorf("Got %v, expected %v", got, expected)
	}

}
