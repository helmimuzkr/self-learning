package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"repository-pattern-go/model"
	"strconv"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	UpdateUser(ctx context.Context, user *model.User) (int64, error)
	DeleteUser(ctx context.Context, id int32) (int64, error)
	FindAllUser(ctx context.Context) ([]*model.User, error)
	FindByIdUser(ctx context.Context, id int32) (*model.User, error)
}

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

// Create
func (repo *userRepositoryImpl) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	query := "INSERT INTO users(name, username, password) VALUES(?, ?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: failed executing db query %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}
	return rowsAffected, nil
}

// Update
func (repo *userRepositoryImpl) UpdateUser(ctx context.Context, user *model.User) (int64, error) {
	query := "UPDATE users SET name = ?, username = ?, password = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, query, user.Name, user.Username, user.Password, user.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateUser: failed executing db query %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return rowsAffected, errors.New("Id : " + strconv.Itoa(int(user.ID)) + "Not Found")
	}
	return rowsAffected, nil
}

// Delete
func (repo *userRepositoryImpl) DeleteUser(ctx context.Context, id int32) (int64, error) {
	query := "DELETE FROM users WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("DeleteUser: failed executing db query %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return rowsAffected, errors.New("Id : " + strconv.Itoa(int(id)) + "Not Found")
	}
	return rowsAffected, nil
}

// Find All
func (repo *userRepositoryImpl) FindAllUser(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	query := "SELECT name, username FROM users"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return users, fmt.Errorf("FindAllUser: failed executing db query %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		rows.Scan(&user.Name, &user.Username)
		users = append(users, &user)
	}
	return users, nil
}

// Find By Id
func (repo *userRepositoryImpl) FindByIdUser(ctx context.Context, id int32) (*model.User, error) {
	var user *model.User
	query := "SELECT name, username FROM users WHERE id = ?"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	if err != nil {
		return user, fmt.Errorf("FindByIdUser: failed executing db %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user.Name, &user.Username)
		return user, nil
	}else {
		return user, errors.New("Id : " + strconv.Itoa(int(id)) + "Not Found")
	}
}
