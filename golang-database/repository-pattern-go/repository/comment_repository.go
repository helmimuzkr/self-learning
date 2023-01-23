package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"repository-pattern-go/model"
	"strconv"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment model.Comment) (int64, error)
	UpdateComment(ctx context.Context, comment model.Comment, id int32) (int64, error)
	DeleteComment(ctx context.Context, id int32) (int64, error)
	FindAllComment(ctx context.Context) ([]model.Comment, error)
	FindByIdComment(ctx context.Context, id int32 ) (model.Comment, error)
}

type commentRepositoryImpl struct {
	DB *sql.DB
}

// constructor
func NewCommentRepository(db *sql.DB) CommentRepository{
	return &commentRepositoryImpl{DB: db}
}

// INSERT
func (repo *commentRepositoryImpl) InsertComment(ctx context.Context, comment model.Comment) (int64, error) {
	// execution
	query := "INSERT INTO comments(name, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, comment.Name, comment.Comment)
	if err != nil {
		fmt.Println("Error inserting row: " + err.Error())
		errorMsg := errors.New("DATABASE ERROR - " + err.Error())
		return 0, errorMsg
	}

	// take the number rows affected
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return RowsAffected, err
	}
	
	return RowsAffected, nil
}

// UPDATE
func (repo *commentRepositoryImpl) UpdateComment(ctx context.Context, comment model.Comment, id int32) (int64, error) {
	// querying update 
	query := "UPDATE comments SET name = ?, comment = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, query, comment.Name, comment.Comment, id)
	if err != nil {
		fmt.Println("Error updating row: ", err.Error())
		errMsg := errors.New("DATABASE ERROR - " + err.Error())
		return 0, errMsg
	}

	// get the number rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Rows Affected Error: ", err)
	}

	return rowsAffected, nil
}

// DELETE
func (repo *commentRepositoryImpl) DeleteComment(ctx context.Context, id int32) (int64, error) {
	// exec delete 
	query := "DELETE FROM comments WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Println("Error deleting row: ", err)
		errMsg := errors.New("DATABASE ERROR - " + err.Error())
		return 0, errMsg
	}

	RowAffected, err := result.RowsAffected()
	if err != nil {
		return RowAffected, err
	}

	return RowAffected, nil
}

// Find All
func (repo *commentRepositoryImpl) FindAllComment(ctx context.Context) ([]model.Comment, error) {
	// This slices is created to accommodate result data
	var comments []model.Comment
	// querying
	query := "SELECT id, name, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Error find row: ", err.Error())
		errMsg := errors.New("DATABASE ERROR - " + err.Error())
		return comments, errMsg
	}
	defer rows.Close()

	// looping for collecting data
	for rows.Next() {
		comment := model.Comment{} // This struct is created to accommodate data from rows.Next()
		rows.Scan(&comment.Id, &comment.Name, &comment.Comment) 
		comments = append(comments, comment) // then append to slice comments
	}

	return comments, nil
}

// Find By Id
func (repo *commentRepositoryImpl) FindByIdComment(ctx context.Context, id int32) (model.Comment, error) {
	// exec delete 
	query := "SELECT id, name, comment FROM comments WHERE id = ?"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Println("Error select row: ", err)
		errMsg := errors.New("DATABASE ERROR - " + err.Error())
		return model.Comment{}, errMsg
	}
	defer rows.Close()

	// This struct is created to accommodate data from rows.Next()
	comment := model.Comment{}
	if rows.Next() {
		// if exist
		rows.Scan(&comment.Id, &comment.Name, &comment.Comment)
		return comment, nil
	}else {
		// if doesnt exist
		return comment, errors.New("Id : " + strconv.Itoa(int(id)) + "Not Found")
	}
}
