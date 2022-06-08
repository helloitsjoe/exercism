package yacht

import (
	"sort"
)

func makeMap(dice []int) map[int]int {
	uniq := make(map[int]int)

	for _, num := range dice {
		// uniq[num] is initialized to 0
		uniq[num] += 1
	}

	return uniq
}

func sum(dice []int) int {
	total := 0
	for _, num := range dice {
		total += num
	}
	return total
}

func testYacht(dice []int) int {
	initial := dice[0]
	for _, num := range dice {
		if num != initial {
			return 0
		}
	}

	return 50
}

func testSingles(value int, dice []int) int {
	total := 0
	for _, num := range dice {
		if num == value {
			total += num
		}
	}
	return total
}

func testFull(dice []int) int {
	uniq := makeMap(dice)
	total := sum(dice)

	for _, v := range uniq {
		if !(v == 2 || v == 3) {
			return 0
		}
	}
	return total
}

func testFourOfAKind(dice []int) int {
	uniq := makeMap(dice)

	// TODO: Can this be cleaner?
	for k, v := range uniq {
		if !(v == 1 || v == 4 || v == 5) {
			return 0
		}
		if v == 4 || v == 5 {
			return k * 4
		}
	}
	return 0
}

func testStraight(size string, dice []int) int {
	prev := dice[0]

	if (size == "little" && dice[0] == 2) || (size == "big" && dice[0] == 1) {
		return 0
	}

	for _, num := range dice[1:] {
		if num != prev+1 {
			return 0
		}
		prev = num
	}

	return 30
}

func Score(dice []int, category string) int {
	sort.Ints(dice)
	// fmt.Println("category", category)
	switch category {
	case "yacht":
		return testYacht(dice)
	case "ones":
		return testSingles(1, dice)
	case "twos":
		return testSingles(2, dice)
	case "threes":
		return testSingles(3, dice)
	case "fours":
		return testSingles(4, dice)
	case "fives":
		return testSingles(5, dice)
	case "sixes":
		return testSingles(6, dice)
	case "full house":
		return testFull(dice)
	case "four of a kind":
		return testFourOfAKind(dice)
	case "little straight":
		return testStraight("little", dice)
	case "big straight":
		return testStraight("big", dice)
	case "choice":
		return sum(dice)
	default:
		return 0
	}
}
