package models

type Category struct {
	ID       uint64  `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Image    *string `db:"image" json:"image,omitempty"`
	ParentID *uint64 `db:"parent_id" json:"parent_id"`

	Audit
}

const CategoryTableName = "categories"

var CategoryField = struct {
	ID        string
	Name      string
	Image     string
	ParentID  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	Image:     "image",
	ParentID:  "parent_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}