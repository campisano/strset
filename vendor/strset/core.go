// Package strset provides a Set type for string elements.
package strset

/* Implementation note: The only methods that change a Set
   after it is created are in updaters.go. If you need an
   immutable Set, delete that and updaters_test.go.
*/

import (
	"bytes"
	"sort"
	"strings"
	"fmt"
)

// Set of SetType type.
type Set struct {
	store map[SetType]struct{}
}

// Make creates and returns a new Set.
func Make(elems ...SetType) Set {
	s := Set{}
	s.store = make(map[SetType]struct{})
	for _, elem := range elems {
		s.store[elem] = struct{}{}
	}
	return s
}

// Len reports the number of elements in the set.
func (s Set) Len() int {
	return len(s.store)
}

// Contains reports whether set contains the element.
// Math: S ∋ e.
func (s Set) Contains(elem SetType) bool {
	_, found := s.store[elem]
	return found
}

// ContainsAll reports whether s contains all the given elements.
func (s Set) ContainsAll(elems ...SetType) bool {
	for _, elem := range elems {
		if _, found := s.store[elem]; !found {
			return false
		}
	}
	return true
}

// ToSlice returns a new slice with the elements of s.
// The order of the elements is undefined.
func (s Set) ToSlice() []SetType {
	elems := make([]SetType, len(s.store))
	i := 0
	for k := range s.store {
		elems[i] = k
		i++
	}
	return elems
}

// SetType returns a SetType representation of s with
// elements in lexicographic order.
func (s Set) String() string {
	elems := s.ToSlice()
	selems := []string{}
	for _, e := range elems {
		selems = append(selems, fmt.Sprint(e))
	}
	sort.Strings(selems)
	var buf bytes.Buffer
	buf.WriteString("Set{")
	buf.WriteString(strings.Join(selems, " "))
	buf.WriteByte('}')
	return buf.String()
}

// allIn reports whether all elements of s exist in other.
func (s Set) allIn(other Set) bool {
	for elem := range s.store {
		if _, found := other.store[elem]; !found {
			return false
		}
	}
	return true
}

// Equal reports whether set is equal to other.
func (s Set) Equal(other Set) bool {
	return len(s.store) == len(other.store) && s.allIn(other)
}

// Copy returns a new Set: a copy of s.
func (s Set) Copy() Set {
	res := Make()
	for elem := range s.store {
		res.store[elem] = struct{}{}
	}
	return res
}

// Channel returns a channel to a goroutine
// yielding elements one by one.
func (s Set) Channel() <-chan SetType {
	ch := make(chan SetType)
	go func() {
		for elem := range s.store {
			ch <- elem
		}
		close(ch)
	}()
	return ch
}
