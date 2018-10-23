package main

import (
	"strconv"
	"strings"
)

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

func (f *Field) forEachInRow(i int, foo ForEachFunc) (ok bool) {
	for j := range f.field[i] {
		ok = foo(&f.field[i][j], i, j)
		if !ok {
			return false
		}
	}

	return true
}

func (f *Field) forEachInColumn(j int, foo ForEachFunc) (ok bool) {
	for i := range f.field {
		ok = foo(&f.field[i][j], i, j)
		if !ok {
			return false
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

func (f *Field) forEachMatters(i, j int, foo ForEachFunc) (ok bool) {
	ok = f.forEachInRow(i, foo)
	if !ok {
		return false
	}

	ok = f.forEachInColumn(j, foo)
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
