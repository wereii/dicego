package main

import (
	"dicego/src/dicelib"
	"fmt"
	"time"
)

func main() {
	seed := time.Now().UnixNano()

	fmt.Println("Round seed is:", seed)

	dSet := dicelib.NewDiceSet(6, 6, seed)
	dSet.Throw()

	fmt.Printf("Threw: %v\n", dSet.GetValues())
}
