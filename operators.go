package strset

// Intersection returns a new set with elements that are in s AND in other.
// Math: S ∩ Z.
func (s Set) Intersection(other Set) Set {
	var longer, shorter Set
	if s.Len() > other.Len() {
		longer = s
		shorter = other
	} else {
		longer = other
		shorter = s
	}
	res := Make()
	for elem := range shorter.store {
		if longer.Has(elem) {
			res.Add(elem)
		}
	}
	return res
}

// Union returns a new Set: with elements that are in s OR in other.
// Math: S ∪ Z.
func (s Set) Union(other Set) Set {
	res := s.Copy()
	for elem := range other.store {
		res.Add(elem)
	}
	return res
}

// Difference returns a new Set with the elements of s minus the elements of other.
// Math: S \ Z.
func (s Set) Difference(other Set) Set {
	res := Make()
	for elem := range s.store {
		if !other.Has(elem) {
			res.Add(elem)
		}
	}
	return res
}

// SymmetricDifference returns a new Set with members present
// in either set but not on both. Think boolean XOR.
// Math: S ∆ Z.
func (s Set) SymmetricDifference(other Set) Set {
	all := s.Union(other)
	common := s.Intersection(other)
	return all.Difference(common)
}

// SubsetOf reports whether s is a subset of other
// Math: S ⊆ Z.
func (s Set) SubsetOf(other Set) bool {
	if s.Len() > other.Len() {
		return false
	}
	return s.allIn(other)
}

// SupersetOf reports whether s is a superset of other
// Math: S ⊇ Z.
func (s Set) SupersetOf(other Set) bool {
	return other.SubsetOf(s)
}

