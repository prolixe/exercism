package change

import (
	"fmt"
	"sort"
)

const testVersion = 1

func Change(coins []int, target int) ([]int, error) {

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	changeList := make([][]int, 0)
	for i := 0; i < len(coins); i++ {
		t := target
		change := make([]int, 0)
		for _, coin := range coins[i:] {

			if coin > t {
				continue
			}
			for t >= coin {
				t -= coin
				change = append(change, coin)
				if t == 0 {
					changeList = append(changeList, change)
					break
				}
			}
		}
	}
	if len(changeList) == 0 {
		return nil, fmt.Errorf("could not find a valid combination of coins for %d", target)
	}

	// Get the smallest change list.
	smallestList := changeList[0]
	for _, cl := range changeList {
		if len(cl) < len(smallestList) {
			smallestList = cl
		}
	}

	sort.IntSlice(smallestList).Sort()
	return smallestList, nil

}
