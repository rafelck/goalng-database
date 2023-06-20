package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCategoryInsert(t *testing.T) {
	categoryRepository := NewCategoryRepository(golang_database.GetConnection())

	ctx := context.Background()
	category := entity.Category{
		Name: "Kursi Belajar",
	}

	result, err := categoryRepository.Insert(ctx, category)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	categoryRepository := NewCategoryRepository(golang_database.GetConnection())
	ctx := context.Background()
	var id int32
	id = 300
	category, err := categoryRepository.FindById(ctx, id)

	if err != nil {
		panic(err)
	}

	fmt.Println(category)
}

func TestFindAll(t *testing.T) {
	categoryRepository := NewCategoryRepository(golang_database.GetConnection())
	ctx := context.Background()
	categories, err := categoryRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}
}
