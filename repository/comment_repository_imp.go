package repository

import (
	"context"
	"database/sql"
	"errors"
	et "fajar7xx/pzn-golang-db/entity"
	"strconv"
)

// agar struct ini tidak bisa di akses oleh luar
type commentRepositoryImpl struct {
	// parameter agar bisa di pakai pada semua functin dibawahnya
	DB *sql.DB
}

// returnnya harapannya adalah commenrepository
func NewCommentRepository(db *sql.DB) CommentRepository{
	return &commentRepositoryImpl{
		DB: db,
	}
}

func (r *commentRepositoryImpl) Insert(ctx context.Context, comment et.Comment) (et.Comment, error) {
	query := "INSERT INTO comments(email, comment)VALUES(?,?)"
	result, err := r.DB.ExecContext(ctx, query, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (r *commentRepositoryImpl) FindById(ctx context.Context, id int32) (et.Comment, error) {
	comment := et.Comment{}
	query := "SELECT id, email, comment FROM comments where id = ? LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		// ada
		rows.Scan(
			&comment.Id,
			&comment.Email,
			&comment.Comment,
		)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}

}

func (r *commentRepositoryImpl) FindAll(ctx context.Context) ([]et.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []et.Comment
	for rows.Next() {
		comment := et.Comment{}
		rows.Scan(
			&comment.Id,
			&comment.Email,
			&comment.Comment,
		)
		comments = append(comments, comment)
	}
	return comments, nil
}
