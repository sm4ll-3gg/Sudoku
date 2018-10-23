package main

type Cell struct {
	value uint8

	prediction Set
}

func NewCell(val uint8) Cell {
	return Cell{
		value: val,
	}
}

func (c Cell) Empty() bool {
	return c.value == 0
}

func (c Cell) Value() uint8 {
	return c.value
}

func (c *Cell) SetValue(val uint8) {
	c.value = val
	c.SetPrediction(nil)
}

func (c Cell) Prediction() Set {
	return c.prediction
}

func (c *Cell) SetPrediction(p Set) {
	c.prediction = p
}

func (c *Cell) EraseFromPrediction(val uint8) {
	delete(c.prediction, val)
}
