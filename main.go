package main

import (
	"dicego/src/dicelib"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Println("Round seed is:", seed)

	dices := [6]*dicelib.Dice{}

	fmt.Printf("%#v", dices)

}
