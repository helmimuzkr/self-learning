package types

type User struct {
	Name string `json:"name" form:"name"`
	Age  int32  `json:"age" form:"age"`
}

type RegisterUserRequest struct {
	User
}
