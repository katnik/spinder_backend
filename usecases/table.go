package usecases

import (
	"github.com/katnik/spinder_backend/entities"
)

type TableInteractor struct {
	Tables []*entities.Table
}

func rotate(data []int, k int) []int {

	return append(data[k:], data[0:k]...)

}

func (ti TableInteractor) CreateTables(number int) []*entities.Table {

	for i := 0; i < number; i++ {

		t := new(entities.Table)
		t.ID = i
		t.Sitting = ti.Sitting[i]
		t.Standing = standing[i]
		t.Current = ti.Questions[i]
		tables = append(tables, t)
	}
	return tables
}
