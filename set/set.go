package set

type IHashable interface {
	comparable
	Hash(size uint64) uint64
}

type Set[set_type IHashable] struct {
	entries [][]set_type
	size    uint64
}

func MakeSet[set_type IHashable](size uint64) Set[set_type] {
	return Set[set_type]{
		entries: make([][]set_type, size),
		size:    size,
	}
}

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
