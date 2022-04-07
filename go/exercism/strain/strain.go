package strain

type Ints []int
type Lists [][]int
type Strings []string

// Generics in go 1.18!
func keep[T any](filter func(T) bool, input []T) []T {
	if input == nil {
		return nil
	}

	toKeep := make([]T, 0, len(input))
	for _, item := range input {
		if filter(item) {
			toKeep = append(toKeep, item)
		}
	}
	return toKeep
}

func discard[T any](filter func(T) bool, input []T) []T {
	if input == nil {
		return nil
	}
	toDiscard := make([]T, 0, len(input))
	for _, item := range input {
		if !filter(item) {
			toDiscard = append(toDiscard, item)
		}
	}
	return toDiscard
}

func (i Ints) Keep(filter func(int) bool) Ints {
	return keep[int](filter, i)
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return discard[int](filter, i)
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	return keep[[]int](filter, l)
}

func (s Strings) Keep(filter func(string) bool) Strings {
	return keep[string](filter, s)
}
