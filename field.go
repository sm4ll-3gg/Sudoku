package main

import (
	"strconv"
	"strings"
)

type Field struct {
	field [9][9]Cell
}

func (f *Field) Init(data [9][9]uint8) {
	for i, row := range data {
		for j, value := range row {
			f.field[i][j] = NewCell(value)
		}
	}

	f.makePrediction()
	f.trySetValues()
	f.controller()
}

type ForEachFunc func(c *Cell, i, j int) bool

func (f *Field) forEach(foo ForEachFunc) (ok bool) {
	for i := range f.field {
		for j := range f.field[i] {
			ok = foo(&f.field[i][j], i, j)
			if !ok {
				return false
			}
		}
	}

	return true
}

func (f *Field) forEachInSector(i, j int, foo ForEachFunc) (ok bool) {
	si := (i / 3) * 3
	sj := (j / 3) * 3

	for i := si; i < si+3; i++ {
		for j := sj; j < sj+3; j++ {
			ok = foo(&f.field[i][j], i, j)
			if !ok {
				return false
			}
		}
	}

	return true
}

func (f *Field) makePrediction() {
	ok := f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		p, ok := f.predictor(i, j)
		if !ok {
			return false
		}

		c.SetPrediction(p)
		return true
	})

	if !ok {
		panic("Fuckup\n" + f.String())
	}
}

func (f *Field) trySetValues() {
	f.forEach(func(c *Cell, i, j int) bool {
		prediction := c.Prediction()
		if len(prediction) != 1 {
			return true
		}

		var value uint8
		for v := range prediction {
			value = v
		}

		c.SetValue(value)
		f.updatePredictions(i, j, value)

		return true
	})
}

func (f *Field) updatePredictions(i, j int, value uint8) {
	for j := range f.field[i] {
		f.updatePrediction(&f.field[i][j], i, j, value)
	}

	for i := range f.field {
		f.updatePrediction(&f.field[i][j], i, j, value)
	}

	f.forEachInSector(i, j, func(c *Cell, i, j int) bool {
		f.updatePrediction(c, i, j, value)

		return true
	})
}

func (f *Field) updatePrediction(c *Cell, i, j int, value uint8) {
	c.EraseFromPrediction(value)

	prediction := c.Prediction()
	if len(prediction) != 1 {
		return
	}

	var val uint8
	for v := range prediction {
		val = v
	}

	c.SetValue(val)
	f.updatePredictions(i, j, val)
}

func (f Field) String() string {
	builder := strings.Builder{}

	addDelimer := func() {
		builder.WriteRune('\n')
		builder.WriteString(strings.Repeat("_", 9+3*8))
		builder.WriteRune('\n')
	}

	for i, row := range f.field {
		addDelimer()

		for j := range row {
			val := strconv.Itoa(int(f.field[i][j].Value()))
			builder.WriteString(val)

			if j != len(row)-1 {
				builder.WriteString(" | ")
			}
		}
	}
	addDelimer()

	return builder.String()
}
