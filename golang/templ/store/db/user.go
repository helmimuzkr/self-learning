package db

import (
	"fmt"

	"github.com/helmimuzkr/nyoba-templ/types"
	"github.com/helmimuzkr/nyoba-templ/utill"
)

func SaveUser(user *types.RegisterUserRequest, id string) error {
	tx, err := utill.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users(id, name, age) VALUES($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("error prepare statement during SaveUser - error: %s", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id, user.Name, user.Age); err != nil {
		tx.Rollback()
		return fmt.Errorf("error exec statement during SaveUser - error: %s", err)
	}

	tx.Commit()

	return nil
}
