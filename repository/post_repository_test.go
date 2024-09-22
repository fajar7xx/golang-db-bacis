package repository

import (
	"context"
	pgd "fajar7xx/pzn-golang-db"
	et "fajar7xx/pzn-golang-db/entity"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPostCreate(t *testing.T) {
	ctx := context.Background()
	postRepository := NewPostRepository(pgd.GetConnection())

	post := et.Post{
		Name: "Fajar",
		Post: "Post terbaru seputar kehidupan prahara lalalal",
	}

	result, err := postRepository.Create(ctx, post)
	if err != nil {
		panic(err)
	}

	fmt.Println("result ", result)
}


func TestPostFindById(t *testing.T){
	ctx := context.Background()
	postRepository := NewPostRepository(pgd.GetConnection())

	post, err := postRepository.FindById(ctx, 1)
	if err != nil{
		panic(err)
	}

	fmt.Println(post)
}

func TestPostFindAll(t *testing.T){
	ctx := context.Background()
	postRepository := NewPostRepository(pgd.GetConnection())
	

	posts, err := postRepository.FindAll(ctx)
	if err != nil{
		panic(err)
	}

	for _, post := range posts{
		fmt.Println(post)
	}
}