package repository

import "github.com/adit1o/train3/entity"

type CategoryRepository interface {
	FindById(id string) *entity.CategoryEntity
}