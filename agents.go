package main

func (f Field) predictor(i, j int) (Set, bool) {
	predict, ok := f.rowObserver(i)
	if !ok {
		return nil, false
	}

	other, ok := f.columnObserver(j)
	if !ok {
		return nil, false
	}

	predict = predict.Or(other)

	other, ok = f.sectorObserver(i, j)
	if !ok {
		return nil, false
	}

	return predict.Or(other).Not(), true
}

func observer(c *Cell, has Set) (ok bool) {
	val := c.Value()
	switch _, ok := has[val]; {
	case ok:
		return false
	case !c.Empty():
		has.Append(val)
	}

	return true
}

func (f Field) rowObserver(i int) (Set, bool) {
	has := make(Set, 9)

	ok := f.forEachInRow(i, func(c *Cell, _, _ int) bool {
		return observer(c, has)
	})

	if !ok {
		return nil, false
	}

	return has, true
}

func (f Field) columnObserver(j int) (Set, bool) {
	has := make(Set, 9)

	ok := f.forEachInColumn(j, func(c *Cell, _, _ int) bool {
		return observer(c, has)
	})

	if !ok {
		return nil, false
	}

	return has, true
}

func (f Field) sectorObserver(i, j int) (Set, bool) {
	has := make(Set, 9)

	ok := f.forEachInSector(i, j, func(c *Cell, _, _ int) bool {
		return observer(c, has)
	})

	if !ok {
		return nil, false
	}

	return has, true
}

func (f Field) controller() bool {
	for i := range f.field {
		if _, ok := f.rowObserver(i); !ok {
			return false
		}
	}

	for j := range f.field[0] {
		if _, ok := f.columnObserver(j); !ok {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if _, ok := f.sectorObserver(i, j); !ok {
				return false
			}
		}
	}

	return true
}
