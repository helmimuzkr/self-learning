package repository

import (
	"context"
	"fmt"
	"repository-pattern-go/config"
	"repository-pattern-go/model"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	// create instance for CommentRepository Interface
	commentRepository := NewCommentRepository(config.GetConnection())
	// close database connection after parent function close
	defer config.GetConnection().Close()


	// context
	ctx := context.Background()

	// input
	comment := model.Comment{
		Name: "Muzakir",
		Comment: "Hallo nama saya Muzakir",
	}

	// running the insert method, returning last insert Id and error
	insertResult, err := commentRepository.InsertComment(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(insertResult)
}

func TestCommentUpdate(t *testing.T) {
	// intansiasi interface CommentRepository
	commentRepo := NewCommentRepository(config.GetConnection())
	// close database connection after parent function close
	defer config.GetConnection().Close()

	// context
	ctx := context.Background()

	// input
	comment := model.Comment{
		Name: "Muhammad",
		Comment: "Test Update",
	}

	// running the update method, returning rows affected and error
	updateResult, err := commentRepo.UpdateComment(ctx, comment, 45)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d rows successfully updated\n", updateResult)
}

func TestCommentDelete(t *testing.T) {
	// create instance for CommentRepository Interface
	commentRepo := NewCommentRepository(config.GetConnection())
	// close database connection after parent function close
	defer config.GetConnection().Close()

	// context
	ctx := context.Background()

	// running the delete method, returning rows affected and error
	deleteResult, err := commentRepo.DeleteComment(ctx, 48)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d data berhasil dihapus\n", deleteResult)
}

func TestCommentFindAll(t *testing.T) {
	// create a new CommentRepository interface
	commentRepo := NewCommentRepository(config.GetConnection())
	// close the connection database after this test func finish
	defer config.GetConnection().Close()

	ctx := context.Background()

	// running tthe FindAll method, returning slices model.Comment and error
	findAllResult, err := commentRepo.FindAllComment(ctx)
	if err != nil {
		panic(err)
	}

	for _, result := range findAllResult {
		fmt.Println(result)
	}
}

func TestCommentFindById(t *testing.T) {
	// creating new instance CommentRepository interface
	commentRepo := NewCommentRepository(config.GetConnection())
	defer config.GetConnection().Close()

	// context
	ctx := context.Background()

	// running the FindById method, returning model.Comment and error
	findByIdResult, err := commentRepo.FindByIdComment(ctx, 46)
	if err != nil {
		panic(err)
	} 
	fmt.Println("Id	: ", findByIdResult.Id)
	fmt.Println("Nama	: ", findByIdResult.Name)
	fmt.Println("Comment : ", findByIdResult.Comment)
}