package Mrgo

import "fmt"

type joinMap struct {
	joinCells []*joinCell
	joinStr   string
}

type joinCell struct {
	syntax  string // left right inner
	table   string
	columnA string
	columnB string
}

func (j joinMap) join(syntax, table, columnA, columnB string) {
	j.joinCells = append(j.joinCells, &joinCell{syntax, table, columnA, columnB})
}

func (j joinMap) build() {
	for index, cell := range j.joinCells {
		fmt.Printf("%v-%v", index, cell)
	}
}
