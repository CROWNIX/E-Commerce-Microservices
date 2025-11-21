package users

type CreateUserInput struct {
	Username string
	Email    string
	Password string
}

type RegisterOutput struct {
	Total uint64 `db:"total"`
}

type GetUserOutput struct {
	ID       uint64 `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
