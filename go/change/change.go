package change

import (
	"fmt"
	"sort"
)

const testVersion = 1

type CoinsList []int

func (c *CoinsList) Sum() int {
	total := 0
	if (*c) == nil {
		return 0
	}
	for _, i := range *c {
		total += i
	}
	return total
}

func Change(coins []int, target int) ([]int, error) {

	// Deal with special cases
	if target == 0 {
		return []int{}, nil
	}

	if target < coins[0] {
		return nil, fmt.Errorf("change %d smallest than smallest of coins", target)
	}

	// Make a list of all minimal change combo up to target.
	smallestChangeCombos := make([]CoinsList, target+1)
	for t := range smallestChangeCombos {
		changeCombinations := make([]CoinsList, 0)
		for _, c := range coins {
			if c > t {
				continue
			}
			// Check the smallest change combo to see if we can get our t value
			// with our coin and a combo for the difference t - c
			subList := smallestChangeCombos[t-c]
			// Need to validate that the sum of the sublist is really t, since
			// some change combo might be impossible.
			if subList.Sum()+c == t {
				changeCombinations = append(changeCombinations, append(subList, c))
			}
		}
		// Get the smallest combo out of the possible combinations of change
		if len(changeCombinations) > 0 {
			smallest := changeCombinations[0]
			for _, pC := range changeCombinations {
				if len(pC) < len(smallest) {
					smallest = pC
				}
			}
			smallestChangeCombos[t] = smallest
		}
	}
	sort.Sort(sort.IntSlice(smallestChangeCombos[target]))
	if smallestChangeCombos[target] == nil {
		return nil, fmt.Errorf("no coins can make %d", target)
	}
	return smallestChangeCombos[target], nil

}
