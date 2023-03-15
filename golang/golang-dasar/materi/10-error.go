package materi

import (
	"errors"
	"fmt"
)

type Data struct {
	Users []User
}

func (d *Data) CheckUser(name string) (string, error) {
	for _, value := range d.Users {
		if value.Name == name {
			return value.Name, nil
		}
	}

	return "error", errors.New("error")
}

func errorBasic() {
	user1 := User{1, "helmi", "muh@gmail.com", false}
	user2 := User{2, "muhammad", "muh@gmail.com", false}


	var data Data
	data.Users = append(data.Users, user1, user2)

	result, err := data.CheckUser("helmi")

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

}