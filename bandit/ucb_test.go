package bandit

import (
	"fmt"
	"math/rand"
)

func ExampleUpperConfidenceBound() {
	// Random value generator.
	r := rand.New(rand.NewSource(42))

	// Create epsilon greedy algorithm instance.
	model := NewUpperConfidenceBound(5, 0.25, 42)

	for i := 0; i < 1000; i++ {
		// Select next arms to be tried.
		arms := model.SelectArms(100)

		// Compute rewards of arms arbitrarily.
		rewards := make([]float64, len(arms))
		for j := 0; j < len(arms); j++ {
			rewards[j] = float64(arms[j]) + r.NormFloat64()
		}

		// Update state with selected arms and rewards.
		model.Update(arms, rewards)
	}

	// Show results.
	for i := 0; i < model.NumArms; i++ {
		fmt.Printf("Arm: %d Trial: %d Mean: %.3f\n", i, model.State.Trials[i], model.State.Means[i])
	}

	// Output:
	// Arm: 0 Trial: 100 Mean: 0.033
	// Arm: 1 Trial: 100 Mean: 0.989
	// Arm: 2 Trial: 100 Mean: 2.120
	// Arm: 3 Trial: 100 Mean: 3.109
	// Arm: 4 Trial: 99600 Mean: 4.003

}
