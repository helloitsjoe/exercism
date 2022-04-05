package strain

type Ints []int
type Lists [][]int
type Strings []string

// Generics in go 1.18!
// func keep[T comparable](filter func(T) bool, input []T) T {
// 	if input == nil {
// 		return nil
// 	}
// 	toKeep := make([]T, 0, len(input))
// 	for _, item := range input {
// 		if filter(item) {
// 			toKeep = append(toKeep, item)
// 		}
// 	}
// 	return toKeep
// }
//
// func (i Ints) Keep(filter func(int) bool) Ints {
// 	return keep[int](filter, i)
// }

func (i Ints) Discard(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}
	toDiscard := make(Ints, 0, len(i))
	for _, num := range i {
		if !filter(num) {
			toDiscard = append(toDiscard, num)
		}
	}
	return toDiscard
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	toKeep := make(Lists, 0, len(l))
	for _, list := range l {
		if filter(list) {
			toKeep = append(toKeep, list)
		}
	}
	return toKeep
}

func (s Strings) Keep(filter func(string) bool) Strings {
	toKeep := make(Strings, 0, len(s))
	for _, list := range s {
		if filter(list) {
			toKeep = append(toKeep, list)
		}
	}
	return toKeep
}
