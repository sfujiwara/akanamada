# Bandit

GO implementations of BANdit algorithms.

## Basic Usage

```go
package main

import (
	"fmt"
	"github.com/sfujiwara/akanamada/bandit"
	"math/rand/v2"
)

func ComputeRewards(arms []int) []float64 {
	n := len(arms)
	rewards := make([]float64, n)
	for i := 0; i < n; i++ {
		rewards[i] = float64(arms[i]) + rand.NormFloat64()
	}

	return rewards
}

func main() {
	model := bandit.NewEpsilonGreedy(5, 0.1)

	for i := 0; i < 100; i++ {
		arms := model.NextArms(10)
		rewards := ComputeRewards(arms)
		model.Update(arms, rewards)

		fmt.Println("Trials", model.State.Trials, "Means", model.State.Means)
	}
}
```
