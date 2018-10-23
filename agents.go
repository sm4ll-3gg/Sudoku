package main

func (f Field) predictor(i, j int) Set {
	has := make(Set, 9)
	f.forEachMatters(i, j, func(c *Cell, _, _ int) bool {
		if !c.Empty() {
			has.Append(c.Value())
		}

		return true
	})

	return has.Not()
}

func (f Field) controller() bool {
	has := make(Set, 9)
	add := func(c *Cell, _, _ int) bool {
		switch {
		case c.Empty():
			return true
		case has.Contains(c.Value()):
			return false
		default:
			has.Append(c.Value())
		}

		return true
	}

	for i := range f.field {
		if !f.forEachInRow(i, add) {
			return false
		}
	}

	has.Clear()

	for j := range f.field[0] {
		if !f.forEachInColumn(j, add) {
			return false
		}
	}

	has.Clear()

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !f.forEachInSector(i, j, add) {
				return false
			}
		}
	}

	return true
}
