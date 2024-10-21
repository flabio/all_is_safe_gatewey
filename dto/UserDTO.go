package dto

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDTO struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
