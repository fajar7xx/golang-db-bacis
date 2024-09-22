package repository

import (
	"context"
	et "fajar7xx/pzn-golang-db/entity"
)

type CommentRepository interface {
	// tambahkan parameter tx jika ingin transaksional
	Insert(ctx context.Context, comment et.Comment) (et.Comment, error)
	FindById(ctx context.Context, id int32) (et.Comment, error)
	FindAll(ctx context.Context) ([]et.Comment, error)
}
