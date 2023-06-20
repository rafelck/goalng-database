package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type categoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{DB: db}
}

func (repository *categoryRepositoryImpl) Insert(ctx context.Context, category entity.Category) (entity.Category, error) {
	query := "INSERT INTO category(name) VALUES (?)"
	result, err := repository.DB.ExecContext(ctx, query, category.Name)
	if err != nil {
		return category, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return category, err
	}
	category.Id = int32(id)

	return category, nil
}

func (repository *categoryRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Category, error) {
	query := "SELECT id, name FROM category WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	category := entity.Category{}
	if err != nil {
		return category, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
		return category, nil
	} else {
		return category, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *categoryRepositoryImpl) FindAll(ctx context.Context) ([]entity.Category, error) {
	query := "SELECT id, name FROM category"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}

	return categories, nil
}
