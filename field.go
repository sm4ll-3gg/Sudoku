package main

import (
	"errors"
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
}

func (f *Field) FindSolution() error {
	f.makePrediction()

	for !f.completeController() {
		err := f.trySetValues()
		if err != nil {
			return err
		}
	}

	// f.findUniquePredictions()
	if !f.controller() {
		return errors.New(f.String())
	}

	return nil
}

func (f *Field) setCellValue(c *Cell, i, j int, value uint8) {
	c.SetValue(value)

	f.forEachMatters(i, j, func(c *Cell, _, _ int) bool {
		if !c.Empty() {
			return true
		}

		c.EraseFromPrediction(value)
		return true
	})
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

func (f *Field) trySetValues() (err error) {
	f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		prediction := f.predictor(i, j)
		switch len(prediction) {
		case 0:
			err = errors.New("Empty prediction")
			return false
		case 1:
			f.setCellValue(c, i, j, prediction.First())
		default:
			c.SetPrediction(prediction)
		}

		return true
	})

	return err
}

func (f *Field) resolveDoubles() (err error) {
	f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		err = f.minimalist(c, i, j)
		return err == nil
	})

	return err
}

func (f *Field) findUniquePredictions() {
	f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		f.researcher(c, i, j)
		return true
	})
}
