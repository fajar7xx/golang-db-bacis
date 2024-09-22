package repository

import (
	"context"
	"database/sql"
	"errors"
	et "fajar7xx/pzn-golang-db/entity"
	"strconv"
)

// buat dlu contractnya
type PostRepository interface {
	Create(ctx context.Context, post et.Post) (et.Post, error)
	FindById(ctx context.Context, id int) (et.Post, error)
	FindAll(ctx context.Context) ([]et.Post, error)
}

// buat struct tapi jangan buat untuk public akses, ini private
type postRepositoryImpl struct {
	// paramater apa aja yang di terima
	DB *sql.DB
}

// ini yang akan di export dengan return struct di atas
// dengan return interface
func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{
		DB: db,
	}
}

func (r *postRepositoryImpl) Create(ctx context.Context, post et.Post) (et.Post, error) {
	// start db transaction
	// Get a Tx for making transaction requests.
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil{
		return post, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	
	// create a new post
	query := "INSERT INTO posts(name, post) VALUES(?,?)"
	result, err := r.DB.ExecContext(ctx, query, post.Name, post.Post)
	if err != nil{
		return post, err
	}

	// get the id of the post just created
	id, err := result.LastInsertId()
	if err != nil{
		return post, err
	}

	// commit
	if err := tx.Commit(); err != nil{
		return post, nil
	}

	post.Id = int(id)
	return post, nil
}

func (r *postRepositoryImpl) FindById(ctx context.Context, id int) (et.Post, error) {
	post := et.Post{}
	query := "SELECT id, name, post FROM posts where id = ? LIMIT 1"
	row, err := r.DB.QueryContext(ctx, query, id)
	if err != nil{
		return post, err
	}
	defer row.Close()

	if row.Next(){
		row.Scan(
			&post.Id,
			&post.Name,
			&post.Post,
		)
		return post, nil
	}else{
		return post, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (r *postRepositoryImpl) FindAll(ctx context.Context) ([]et.Post, error) {
	query := "SELECT id, name, post FROM posts"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var posts []et.Post
	for rows.Next(){
		post := et.Post{}
		rows.Scan(
			&post.Id,
			&post.Name,
			&post.Post,
		)
		posts = append(posts, post)
	}
	return posts, nil
}
