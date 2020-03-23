package dicelib

import (
	"math/rand"
	"testing"
)

const (
	defaultRandSeed int64 = 1
)

func TestMakeDice(t *testing.T) {
	rSrc := rand.NewSource(defaultRandSeed)
	dRand := rand.New(rSrc)

	expected := &Dice{faces: 1, randgen: dRand}
	got := NewDice(1, dRand)
	if *got != *expected {
		t.Errorf("Expected \"%#v\" got \"%#v\"", *expected, *got)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Creating zero faced dice should panic.")
		}
	}()

	NewDice(0, dRand)
}

func TestRoll(t *testing.T) {
	diceFaces := 6

	eSrc := rand.NewSource(defaultRandSeed)
	expRand := rand.New(eSrc)

	gSrc := rand.NewSource(defaultRandSeed)
	gotRand := rand.New(gSrc)

	exp := &Dice{faces: diceFaces, randgen: expRand}
	exp.roll()

	d1 := NewDice(diceFaces, gotRand)
	d1.roll()

	expected := exp.value
	got := d1.value

	if got != expected {
		t.Errorf("Should roll %d, got %d", expected, d1.value)
	}
}
