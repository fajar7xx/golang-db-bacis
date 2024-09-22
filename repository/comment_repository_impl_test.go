package repository

import (
	"context"
	pgd "fajar7xx/pzn-golang-db"
	et "fajar7xx/pzn-golang-db/entity"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(pgd.GetConnection())

	ctx := context.Background()
	comment := et.Comment{
		Email:   "fajarrepository@mail.lokal",
		Comment: "komentar ini mengenai repository pattern",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("result ", result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(pgd.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 34)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(pgd.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments{
		fmt.Println(comment)
	}
}
