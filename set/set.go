package set

//IHashable is the interface where all types that
//all types which can be stored in a set derive
//from.
type IHashable interface {
	comparable
	Hash(size uint64) uint64
}

//A set is a collection of unique items, the
//uniqueness is implemented in a way similar to a
//hashtable, map or dictionary.
type Set[set_type IHashable] struct {
	entries [][]set_type
	size    uint64
}

//Creates a new set with the given type
func MakeSet[set_type IHashable](size uint64) Set[set_type] {
	return Set[set_type]{
		entries: make([][]set_type, size),
		size:    size,
	}
}

//Adds an item to the set if it is not there already
func (s *Set[set_type]) Push(item set_type) {
	hash := item.Hash(s.size)
	if hash > s.size {
		panic("hash was larger than the size of the Set")
	}

	for _, set_item := range s.entries[hash] {
		if item == set_item {
			return
		}
	}
	s.entries[hash] = append(s.entries[hash], item)
}

//Attempts to remove an item from the set
//and returns whether the item was in the set or not
func (s *Set[set_type]) Pop(item set_type) bool {
	hash := item.Hash(s.size)
	if hash > s.size {
		panic("hash was larger than the size of the Set")
	}

	for i, set_item := range s.entries[hash] {
		if item == set_item {
			s.entries[hash] = append(s.entries[hash][:i], s.entries[hash][i+1:]...)
			return true
		}
	}

	return false
}

//Checks if a set has an item with no side-effects
func (s *Set[set_type]) Has(item set_type) bool {
	hash := item.Hash(s.size)
	if hash > s.size {
		panic("hash was larger than the size of the Set")
	}

	for _, set_item := range s.entries[hash] {
		if item == set_item {
			return true
		}
	}

	return false
}
