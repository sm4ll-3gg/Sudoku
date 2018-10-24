package main

import (
	"errors"
)

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

func (f *Field) minimalist(c *Cell, i, j int) error {
	prediction := c.Prediction()

	hash := func(i, j int) uint8 {
		return uint8(i*10 + j)
	}

	equal := make(Set)
	f.forEachMatters(i, j, func(c *Cell, i, j int) bool {
		if prediction.Equal(c.Prediction()) {
			equal.Append(hash(i, j))
		}

		return true
	})

	if len(equal) > len(prediction) {
		return errors.New("Prediction length less then cells count")
	} else if len(equal) != len(prediction) {
		return nil
	}

	f.forEachMatters(i, j, func(c *Cell, i, j int) bool {
		if equal.Contains(hash(i, j)) {
			return true
		} else if !prediction.Equal(c.Prediction()) {
			return true
		}

		for key := range prediction {
			c.EraseFromPrediction(key)
		}

		p := c.Prediction()
		if len(p) != 1 {
			return true
		}

		for val := range p {
			f.setCellValue(c, i, j, val)
		}

		return true
	})

	return nil
}

func (f *Field) researcher(c *Cell, i, j int) {
	prediction := c.Prediction()

	counts := make(map[uint8]uint8, len(prediction))
	f.forEachMatters(i, j, func(c *Cell, ci, cj int) bool {
		if !c.Empty() || (ci == i && cj == j) {
			return true
		}

		curr := c.Prediction()
		for key := range prediction {
			if curr.Contains(key) {
				counts[key]++
			}
		}

		return len(counts) < len(prediction)
	})

	for key, count := range counts {
		if count == 1 {
			f.setCellValue(c, i, j, key)
			break
		}
	}
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

		has.Clear()
	}

	for j := range f.field[0] {
		if !f.forEachInColumn(j, add) {
			return false
		}

		has.Clear()
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !f.forEachInSector(i, j, add) {
				return false
			}

			has.Clear()
		}
	}

	return true
}
