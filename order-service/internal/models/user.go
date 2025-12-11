package models

type User struct {
	ID    uint64  `db:"id"`
	Email string `db:"email"`
}

const UserTableName = "users"

var UserField = struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Username:  "username",
	Email:     "email",
	Password:  "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}