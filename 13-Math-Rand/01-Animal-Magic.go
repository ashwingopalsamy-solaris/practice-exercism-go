package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + rand.Intn(19)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	x := []string{
		"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog",
	}
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	return x
}
