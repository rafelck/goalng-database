package repository

import (
	"context"
	"golang-database/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, category entity.Category) (entity.Category, error)
	FindById(ctx context.Context, id int32) (entity.Category, error)
	FindAll(ctx context.Context) ([]entity.Category, error)
}
