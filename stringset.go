package stringset

// Stringset represents a set of strings. It can be used to efficiently
// determine whether a string is part of a previously given string slice.
type Stringset struct {
	containsEmpty bool
	root          *node
}

type node struct {
	successors map[rune]*node
	ends       map[rune]bool
}

func newNode() *node {
	return &node{
		successors: make(map[rune]*node),
		ends:       make(map[rune]bool),
	}
}

// New creates a new Stringset from the given string slice. New elements can be
// added by calling
//
//	set.Add(string)
func New(elements []string) *Stringset {
	set := &Stringset{
		root: newNode(),
	}
	for _, elem := range elements {
		set.Add(elem)
	}
	return set
}

// Add adds a new element to the Stringset, making it available for checking.
func (s *Stringset) Add(elem string) {
	if elem == "" {
		s.containsEmpty = true
	}

	node := s.root
	for i, r := range elem {
		if node.successors[r] == nil {
			n := newNode()
			node.successors[r] = n
		}
		if i == len(elem)-1 {
			node.ends[r] = true
		}
		node = node.successors[r]
	}
}

// Contains efficiently determines whether the given elem is contained in the
// previously given string slice or in the added elements.
func (s *Stringset) Contains(elem string) bool {
	if elem == "" {
		return s.containsEmpty
	}

	node := s.root
	for i, r := range elem {
		next, ok := node.successors[r]
		// log.Printf("ok: %#+v, i: %#+v, len(elem): %#+v, end: %#+v\n", ok, i, len(elem), node.ends[r])
		if !ok && i != len(elem) {
			return false
		}
		if i == len(elem)-1 && !node.ends[r] {
			return false
		}
		node = next
	}
	return true
}
