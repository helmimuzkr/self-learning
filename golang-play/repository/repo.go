package repository

import "fmt"

type Repository interface {
	Find() []string
	Delete(id int) string
}

type RepsitoryImpl struct {
}

func NewRepo() Repository {
	return &RepsitoryImpl{}
}

func (r *RepsitoryImpl) Find() []string {
	data := []string{"data1", "data2"}
	return data
}
func (r *RepsitoryImpl) Delete(id int) string {
	data := fmt.Sprintf("Berhasil menghapus id: %d", id)
	return data
}
