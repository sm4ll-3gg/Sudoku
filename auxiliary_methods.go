package main

import (
	"strconv"
	"strings"
)

type ForCellFunc func(c *Cell, i, j int) bool

func (f *Field) forEach(foo ForCellFunc) (ok bool) {
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

type ForEachFunc func(i, j int, foo ForCellFunc) bool

func (f *Field) forEachInRow(i, j int, foo ForCellFunc) (ok bool) {
	for j := range f.field[i] {
		ok = foo(&f.field[i][j], i, j)
		if !ok {
			return false
		}
	}

	return true
}

func (f *Field) forEachInColumn(i, j int, foo ForCellFunc) (ok bool) {
	for i := range f.field {
		ok = foo(&f.field[i][j], i, j)
		if !ok {
			return false
		}
	}

	return true
}

func (f *Field) forEachInSector(i, j int, foo ForCellFunc) (ok bool) {
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

func (f *Field) forEachMatters(i, j int, foo ForCellFunc) (ok bool) {
	ok = f.forEachInRow(i, j, foo)
	if !ok {
		return false
	}

	ok = f.forEachInColumn(i, j, foo)
	if !ok {
		return false
	}

	return f.forEachInSector(i, j, foo)
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

func (f *Field) completeController() bool {
	return f.forEach(func(c *Cell, _, _ int) bool {
		if c.Empty() {
			return false
		}

		return true
	})
}
