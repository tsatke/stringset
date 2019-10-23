package stringset

import (
	"fmt"
	"io"
	"strconv"
)

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

func New(elements []string) *Stringset {
	set := &Stringset{
		root: newNode(),
	}
	for _, elem := range elements {
		set.Add(elem)
	}
	return set
}

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

func (s *Stringset) dump(w io.Writer) {
	printNode(w, s.root, 0)
}

func printNode(w io.Writer, n *node, indent int) {
	for k, v := range n.successors {
		fmt.Fprintf(w, "%"+strconv.Itoa(indent)+"s", string(k))
		if n.ends[k] {
			fmt.Fprint(w, " (end)")
		}
		fmt.Fprint(w, "\n")
		printNode(w, v, indent+1)
	}
}
