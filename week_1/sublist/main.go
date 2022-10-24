package sublist

import "reflect"

// Relation is the comparison between lists
type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

// Sublist checks difference of two lists and
// returns equal, sublist, superlist or unequal according
// to their relation to each other.
func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return RelationEqual
	} else if contains(l1, l2) {
		return RelationSuperlist
	} else if contains(l2, l1) {
		return RelationSublist
	}
	return RelationUnequal
}

func isEqual(l1, l2 []int) bool {
	// Тут має бути рішення
	// написавши код - необхідно запустити тести
	// Ці коментарі можна видаляти
	// !ВАЖЛИВО - не забудьте виправити return
	if len(l1) != len(l2) {
		return false
	}

	for index, value := range l1 {
		if value != l2[index] {
			return false
		}
	}

	return true
}

func contains(l1, l2 []int) bool {
	// Тут має бути рішення
	// написавши код - необхідно запустити тести
	// Ці коментарі можна видаляти
	// !ВАЖЛИВО - не забудьте виправити return
	if len(l1) < len(l2) {
		return false
	}

	for i := 0; i < (len(l1) - len(l2) + 1); i++ {
		if reflect.DeepEqual(l1[i:(len(l2)+i)], l2) {
			return true
		}
	}

	return false
}
