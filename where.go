package Mrgo

import "fmt"

type whereMap struct {
	whereCells []*whereCell
	whereStr   string
	whereVal   []interface{}
}

type whereCell struct {
	bridge    string // AND or OR
	column    string
	condition string // =,!=,<,>,<>,eq,neq...
	value     interface{}
}

func (w whereMap) where(column, condition string, value interface{}) {
	w.whereCells = append(w.whereCells, &whereCell{"and", column, condition, value})
}

func (w whereMap) build() {
	for index, cell := range w.whereCells {
		fmt.Printf("%v-%v", index, cell)
	}
}
