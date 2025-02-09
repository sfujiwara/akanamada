// Package bandit is implementations of various multi-armed bandit algorithms.
package bandit

import (
	"math/rand"
)

type EpsilonGreedyState struct {
	Trials []int
	Means  []float64
}

type EpsilonGreedy struct {
	State   *EpsilonGreedyState
	Epsilon float64
	NumArms int
	rand    *rand.Rand
}

// NewEpsilonGreedy creates EpsilonGreedy instance.
func NewEpsilonGreedy(narms int, epsilon float64, seed int64) *EpsilonGreedy {
	state := &EpsilonGreedyState{
		Trials: make([]int, narms),
		Means:  make([]float64, narms),
	}

	return &EpsilonGreedy{
		State:   state,
		NumArms: narms,
		Epsilon: epsilon,
		rand:    rand.New(rand.NewSource(seed)),
	}
}

// Update bandit model's state with selected arms and rewards.
func (b EpsilonGreedy) Update(arms []int, rewards []float64) {
	for i := 0; i < len(arms); i++ {
		sum := float64(b.State.Trials[arms[i]])*b.State.Means[arms[i]] + rewards[i]
		mean := sum / float64(b.State.Trials[arms[i]]+1)

		b.State.Trials[arms[i]] += 1
		b.State.Means[arms[i]] = mean
	}
}

// SelectArms returns arms to be selected.
func (b EpsilonGreedy) SelectArms(num int) []int {
	best := 0
	for i := 0; i < b.NumArms; i++ {
		if b.State.Means[best] < b.State.Means[i] {
			best = i
		}
	}

	arms := make([]int, num)
	for i := 0; i < num; i++ {
		if b.rand.Float64() <= b.Epsilon {
			// Select arm at random.
			arms[i] = b.rand.Intn(b.NumArms)
		} else {
			// Select current best arm.
			arms[i] = best
		}
	}

	return arms
}
