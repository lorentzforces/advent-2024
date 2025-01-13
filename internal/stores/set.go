package stores

import "iter"

// A super-basic implementation of a set collection.
type Set[T comparable] struct{
	mappedVals map[T]struct{}
}

func EmptySet[T comparable]() Set[T] {
	return Set[T]{
		mappedVals: make(map[T]struct{}, 0),
	}
}

func (self Set[T]) Len() int {
	return len(self.mappedVals)
}

// Adds an item to the Set. Returns true if the Set already contained the given item.
func (self Set[T]) Put(item T) bool {
	if _, hasItem := self.mappedVals[item]; hasItem {
		return true
	}

	self.mappedVals[item] = struct{}{}
	return false
}

func (self Set[T]) Contains(item T) bool {
	_, hasItem := self.mappedVals[item]
	return hasItem
}

// Returns an iterator over the contents of the Set.
func (self Set[T]) Vals() iter.Seq[T] {
	return func(loopFunc func(T) bool) {
		for v := range self.mappedVals {
			if !loopFunc(v) { return }
		}
	}
}
