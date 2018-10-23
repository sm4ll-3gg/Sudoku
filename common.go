package main

var full = Set{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
	6: {},
	7: {},
	8: {},
	9: {},
}

type Set map[uint8]struct{}

var EmptySet = Set(nil)

func (s Set) Append(value uint8) {
	s[value] = struct{}{}
}

func (s Set) Or(other Set) Set {
	res := make(Set)

	for val := range s {
		res.Append(val)
	}

	for val := range other {
		res.Append(val)
	}

	return res
}

func (s Set) Not() Set {
	res := make(Set)

	for val := range full {
		if _, ok := s[val]; !ok {
			res[val] = struct{}{}
		}
	}

	return res
}
