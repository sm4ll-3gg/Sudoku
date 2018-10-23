package main

func (f Field) predictor(i, j int) (Set, bool) {
	predict, ok := f.rowObserver(i)
	if !ok {
		return EmptySet, false
	}

	other, ok := f.columnObserver(j)
	if !ok {
		return EmptySet, false
	}

	predict = predict.Or(other)

	other, ok = f.sectorObserver(i, j)
	if !ok {
		return EmptySet, false
	}

	return predict.Or(other).Not(), true
}

func (f Field) rowObserver(i int) (Set, bool) {
	has := make(Set, 9)

	row := f.field[i]
	for j := range row {
		val := row[j].Value()
		switch _, ok := has[val]; {
		case ok:
			return EmptySet, false
		case !row[j].Empty():
			has.Append(val)
		}
	}

	return has, true
}

func (f Field) columnObserver(j int) (Set, bool) {
	has := make(Set, 9)

	for _, row := range f.field {
		val := row[j].Value()
		switch _, ok := has[val]; {
		case ok:
			return EmptySet, false
		case !row[j].Empty():
			has.Append(val)
		}
	}

	return has, true
}

func (f Field) sectorObserver(i, j int) (Set, bool) {
	has := make(Set, 9)

	ok := f.forEachInSector(i, j, func(c *Cell, _, _ int) bool {
		val := c.Value()
		switch _, ok := has[val]; {
		case ok:
			return false
		case !c.Empty():
			has.Append(val)
		}

		return true
	})

	if !ok {
		return EmptySet, false
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
