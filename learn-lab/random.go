package learnlab

import (
	"fmt"
	"math/rand"
)

func randomNumbers() {

	// Seed : starting point for generating a sequence of random numbers
	val := rand.New(rand.NewSource(42))
	
	// returns integer in the interval -> [0, n) with seed 42
	fmt.Println(val.Intn(101))

	// float between [0.0, 1.0)
	fmt.Println(val.Float64())
}