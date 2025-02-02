package bandit

import (
	"math/rand"
)

type EpsilonGreedyState struct {
	Trials []int
	Means  []float64
}

type EpsilonGreedyParam struct {
	Epsilon float64
}

type EpsilonGreedy struct {
	State   *EpsilonGreedyState
	Param   *EpsilonGreedyParam
	NumArms int
}

func NewEpsilonGreedy(narms int, epsilon float64) *EpsilonGreedy {
	param := &EpsilonGreedyParam{
		Epsilon: epsilon,
	}
	state := &EpsilonGreedyState{
		Trials: make([]int, narms),
		Means:  make([]float64, narms),
	}

	return &EpsilonGreedy{
		Param:   param,
		State:   state,
		NumArms: narms,
	}
}

func (b EpsilonGreedy) Update(arms []int, rewards []float64) {
	for i := 0; i < b.NumArms; i++ {
		sum := float64(b.State.Trials[arms[i]])*b.State.Means[arms[i]] + rewards[i]
		mean := sum / float64(b.State.Trials[arms[i]]+1)

		b.State.Trials[arms[i]] += 1
		b.State.Means[arms[i]] = mean
	}
}

func (b EpsilonGreedy) NextArms(num int) []int {
	r := rand.Float64()

	best := 0
	for i := 0; i < b.NumArms; i++ {
		if b.State.Means[best] < b.State.Means[i] {
			best = i
		}
	}

	arms := make([]int, num)
	for i := 0; i < num; i++ {
		if r <= b.Param.Epsilon {
			arms[i] = rand.Intn(b.NumArms)
		} else {
			arms[i] = best
		}
	}

	return arms
}
