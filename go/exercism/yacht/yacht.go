package yacht

import (
	"fmt"
	"sort"
)

func testYacht(dice []int) int {
	initial := dice[0]
	for _, num := range dice {
		if num != initial {
			return 0
		}
	}

	return 50
}

func testOnes(dice []int) int {
	uniq := make(map[int]int)
	max := 0

	for _, num := range dice {
		ex := uniq[num]

		if ex != 0 {
			uniq[num] += 1
		} else {
			uniq[num] = 1
		}

		if uniq[num] > max {
			max = uniq[num]
		}
	}

	return max
}

func Score(dice []int, category string) int {
	sort.Ints(dice)

	switch category {
	case "yacht":
		return testYacht(dice)
	case "ones":
		return testOnes(dice)
	default:
		fmt.Println("Default")
		return 0
	}
}
