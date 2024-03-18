package services

import (
	"errors"

	"github.com/adit1o/train3/entity"
	"github.com/adit1o/train3/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}


func (service CategoryService) FindById(id string) (*entity.CategoryEntity, error) {
	category := service.Repository.FindById(id)

	if category != nil {
		return nil, errors.New("category not found")
	} 

	return category, nil
}