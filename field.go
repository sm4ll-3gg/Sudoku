package main

import "errors"

type Field struct {
	field [9][9]Cell
}

func (f *Field) Init(data [9][9]uint8) {
	for i, row := range data {
		for j, value := range row {
			f.field[i][j] = NewCell(value)
		}
	}
}

func (f *Field) FindSolution() error {
	f.makePrediction()
	f.trySetValues()

	if !f.controller() {
		return errors.New(f.String())
	}

	return nil
}

func (f *Field) makePrediction() {
	f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		c.SetPrediction(f.predictor(i, j))
		return true
	})
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
	f.forEachInRow(i, func(c *Cell, i, j int) bool {
		f.updatePrediction(c, i, j, value)
		return true
	})

	f.forEachInColumn(j, func(c *Cell, i, j int) bool {
		f.updatePrediction(c, i, j, value)
		return true
	})

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
