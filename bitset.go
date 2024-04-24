package bitset

import (
	"bytes"
	"fmt"
)

// An BitSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type BitSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *BitSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *BitSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Remove removes x from the set
func (s *BitSet) Remove(x int) bool {
	exists := s.Has(x)
	if !exists {
		return false
	}
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit

	nonZeroWordIndex := -1
	// find the first non-zero word
	for i := len(s.words) - 1; i >= 0; i-- {
		if s.words[i] != 0 {
			nonZeroWordIndex = i
			break
		}
	}

	s.words = s.words[:nonZeroWordIndex+1]
	
	return true
}

// Clear remove all elements from the set
func (s *BitSet) Clear() {
	s.words = make([]uint64, 0)
}

// Len return the numbers of the elements
func (s *BitSet) Len() int {
	if len(s.words) == 0 {
		return 0
	}

	sum := 0
	for _, word := range s.words {
		sum += bitCount(word)
	}
	return sum
}

// Copy return a copy of the set
func (s *BitSet) Copy() *BitSet {
	words := make([]uint64, len(s.words))
	copy(words, s.words)
	return &BitSet{words: words}
}

// UnionWith sets s to the union of s and t.
func (s *BitSet) UnionWith(t *BitSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *BitSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func bitCount(n uint64) int {
	count := 0
	for n > 0 {
		count++
		n &= n - 1
	}
	return count
}
