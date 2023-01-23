package repository

import (
	"context"
	"fmt"
	"repository-pattern-go/config"
	"repository-pattern-go/model"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertUser(t *testing.T) {
	userRepo := NewUserRepository(config.GetConnection())
	ctx := context.Background()
	user := model.User{
		Name: "Helmi",
		Username: "mzkr",
		Password: "helmi",
	}
	rowAffected, err := userRepo.CreateUser(ctx, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d row berhasil ditambahkan!\n", rowAffected)
	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	userRepo := NewUserRepository(config.GetConnection())
	ctx := context.Background()
	user := model.User{
		ID: 2,
		Name: "Muhammad",
		Username: "mhmd",
		Password: "muhammad",
	}
	rowsAffected, err := userRepo.UpdateUser(ctx, &user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d row berhasil diupdate!\n", rowsAffected)
	fmt.Println(user)
}
