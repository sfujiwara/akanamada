package optimize

import "fmt"

func terminate(left []bool, right []bool) bool {
	for _, l := range left {
		if !l {
			return false
		}
	}

	for _, r := range right {
		if !r {
			return false
		}
	}

	return true
}

func Hungarian(weightMatrix [][]int, maximize bool) [][]bool {
	nrow := len(weightMatrix)
	ncol := len(weightMatrix)

	// Initialize vertex labeling: l.h = max{w(l, r): r in R}, r.h = 0.
	lh := make([]int, nrow)
	rh := make([]int, ncol)

	for row := 0; row < nrow; row++ {
		for col := 0; col < ncol; col++ {
			if lh[row] < weightMatrix[row][col] {
				lh[row] = weightMatrix[row][col]
			}
		}
	}

	// Initial matching.
	matching := make([][]bool, nrow)
	for row := 0; row < ncol; row++ {
		matching[row] = make([]bool, ncol)
	}

	// True if the node is matched (reduce sum of matching).
	rmatch := make([]bool, ncol)
	lmatch := make([]bool, nrow)

	// Initial matching by greedy bipartite matching.
	for row := 0; row < nrow; row++ {
		for col := 0; col < ncol; col++ {
			if weightMatrix[row][col] == lh[row]+rh[col] && !rmatch[col] {
				matching[row][col] = true
				rmatch[col] = true
				lmatch[row] = true
				break
			}
		}
	}

	// Main loop of hungarian algorithm.
	for i := 0; i < 8; i++ {
		fmt.Println("Start BFS.")
		// Queue for BFS.
		var queue [][]int

		// Breadth first forest.
		fl := make([]bool, nrow)
		fr := make([]bool, ncol)

		fmt.Println("l.h:", lh)
		fmt.Println("r.h:", rh)
		fmt.Println("Matching:")
		for i := 0; i < len(matching); i++ {
			fmt.Println(matching[i])
		}
		fmt.Println("L-Matching:", lmatch)
		fmt.Println("R-Matching:", rmatch)

		// Termination criteria.
		if terminate(lmatch, rmatch) {
			fmt.Println("Found the best matching!")
			break
		}

		for row := 0; row < nrow; row++ {
			if !lmatch[row] {
				queue = append(queue, []int{row})
				fl[row] = true
			}
		}

		// Loop for BFS to find augmenting path.
		for j := 0; j < 100; j++ {
			fmt.Println("Queue:", queue)
			fmt.Println("F_L:", fl)
			fmt.Println("F_R:", fr)

			if len(queue) == 0 {
				fmt.Println("Queue is empty")
				fmt.Println(fl, fr)
				// TODO: Update vertex labeling.
				delta := 10000
				for row := 0; row < nrow; row++ {
					if !fl[row] {
						continue
					}
					for col := 0; col < ncol; col++ {
						if fr[col] {
							continue
						}
						cand := lh[row] + rh[col] - weightMatrix[row][col]
						if cand < delta {
							delta = cand
						}
					}
				}
				fmt.Println("Update vertex labeling with delta:", delta)
				for row := 0; row < nrow; row++ {
					if fl[row] {
						lh[row] -= delta
					}
				}
				for col := 0; col < ncol; col++ {
					if fr[col] {
						rh[col] += delta
					}
				}
				break
			}

			// Dequeue.
			currentPath := queue[0]
			queue = queue[1:]

			if len(currentPath)%2 == 1 {
				// Case: current node is in L.
				row := currentPath[len(currentPath)-1]

				for col := 0; col < ncol; col++ {
					if fr[col] {
						continue
					}
					connected := weightMatrix[row][col] == lh[row]+rh[col] && !matching[row][col]
					if connected {
						nextPath := make([]int, len(currentPath))
						copy(nextPath, currentPath)
						nextPath = append(nextPath, col)
						queue = append(queue, nextPath)
						fr[col] = true
					}
				}
			} else {
				// Case: current node is in R.
				col := currentPath[len(currentPath)-1]
				isAugmentingPath := true

				for row := 0; row < nrow; row++ {
					if fl[row] {
						continue
					}
					connected := weightMatrix[row][col] == lh[row]+rh[col] && matching[row][col]

					if connected {
						nextPath := make([]int, len(currentPath))
						copy(nextPath, currentPath)
						nextPath = append(nextPath, row)
						queue = append(queue, nextPath)
						fl[row] = true
						isAugmentingPath = false
					}
				}

				if isAugmentingPath {
					fmt.Println(currentPath, "is augmenting path! Update matching.")

					// Update matching.
					for k := 0; k < len(currentPath)-1; k++ {
						if k%2 == 0 {
							// Edge L to R.
							matching[currentPath[k]][currentPath[k+1]] = true
						} else {
							// Edge R to L.
							matching[currentPath[k+1]][currentPath[k]] = false
						}
					}

					lmatch[currentPath[0]] = true
					rmatch[currentPath[len(currentPath)-1]] = true

					break
				}
			}
		}
	}

	return matching
}
