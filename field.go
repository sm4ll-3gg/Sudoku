package main

import (
	"errors"
	"log"
)

type Field struct {
	field [9][9]Cell

	filled int
}

func (f *Field) Init(data [9][9]uint8) {
	for i, row := range data {
		for j, value := range row {
			f.field[i][j] = NewCell(value)

			if value != 0 {
				f.filled++
			}
		}
	}
}

func (f *Field) FindSolution() error {
	curr := f.filled

	for !f.completeController() {
		curr = f.filled

		err := f.makePrediction()
		if err != nil {
			return err
		}

		err = f.trySetValues()
		if err != nil {
			return err
		} else if curr != f.filled {
			continue
		}

		err = f.resolveDoubles()
		if err != nil {
			return err
		} else if curr != f.filled {
			continue
		}

		f.findUniquePredictions()

		if curr == f.filled {
			log.Println("The end")
			break
		}
	}

	if !f.controller() {
		return errors.New(f.String())
	}

	return nil
}

func (f *Field) setCellValue(c *Cell, i, j int, value uint8) {
	c.SetValue(value)
	f.filled++

	f.forEachMatters(i, j, func(c *Cell, ci, cj int) bool {
		if !c.Empty() {
			return true
		}

		p := c.Prediction()
		p.Erase(value)

		if len(p) != 1 {
			return true
		}

		f.setCellValue(c, ci, cj, p.First())

		return true
	})
}

func (f *Field) makePrediction() (err error) {
	f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		p := f.predictor(i, j)
		if len(p) == 0 {
			err = errors.New("Empty prediction")
			return false
		}

		c.SetPrediction(p)
		return true
	})

	return err
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

		err = f.minimalist(c, i, j, f.forEachInRow)
		if err != nil {
			return false
		}

		err = f.minimalist(c, i, j, f.forEachInColumn)
		if err != nil {
			return false
		}

		err = f.minimalist(c, i, j, f.forEachInSector)
		return err == nil
	})

	return err
}

func (f *Field) findUniquePredictions() {
	f.forEach(func(c *Cell, i, j int) bool {
		if !c.Empty() {
			return true
		}

		f.researcher(c, i, j, f.forEachInRow)
		f.researcher(c, i, j, f.forEachInColumn)
		f.researcher(c, i, j, f.forEachInSector)
		return true
	})
}
