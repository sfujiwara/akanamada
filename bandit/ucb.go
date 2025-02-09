package bandit

import (
	"math"
	"math/rand"
)

type UpperConfidenceBoundState struct {
	Trials []int
	Means  []float64
}

type UpperConfidenceBound struct {
	State   *UpperConfidenceBoundState
	NumArms int
	Epsilon float64
	rand    *rand.Rand
}

// NewUpperConfidenceBound creates UpperConfidenceBound instance.
func NewUpperConfidenceBound(narms int, epsilon float64, seed int64) *UpperConfidenceBound {
	state := &UpperConfidenceBoundState{
		Trials: make([]int, narms),
		Means:  make([]float64, narms),
	}

	return &UpperConfidenceBound{
		State:   state,
		NumArms: narms,
		Epsilon: epsilon,
		rand:    rand.New(rand.NewSource(seed)),
	}
}

// Update bandit model's state with selected arms and rewards.
func (b UpperConfidenceBound) Update(arms []int, rewards []float64) {
	for i := 0; i < len(arms); i++ {
		sum := float64(b.State.Trials[arms[i]])*b.State.Means[arms[i]] + rewards[i]
		mean := sum / float64(b.State.Trials[arms[i]]+1)

		b.State.Trials[arms[i]] += 1
		b.State.Means[arms[i]] = mean
	}
}

// SelectArms returns arms to be selected.
func (b UpperConfidenceBound) SelectArms(num int) []int {
	// Compute sum of trials.
	trial := 0
	for i := 0; i < b.NumArms; i++ {
		trial += b.State.Trials[i]
	}

	bestValue := -1e7
	bestIndex := -10
	for i := 0; i < b.NumArms; i++ {
		exploration := b.State.Means[i]
		exploitation := math.Sqrt(2 * math.Log(float64(trial+1)) / float64(b.State.Trials[i]+1))
		v := exploration + exploitation

		if v > bestValue {
			bestValue = v
			bestIndex = i
		}
	}

	arms := make([]int, num)
	for i := 0; i < num; i++ {
		arms[i] = bestIndex
	}

	return arms
}
