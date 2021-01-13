package users

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Credential struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
