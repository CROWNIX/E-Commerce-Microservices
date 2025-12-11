package models

import (
	"time"
)

type Audit struct {
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

var AuditField = struct {
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

func AuditColumns() []string {
	return []string{
		AuditField.CreatedAt,
		AuditField.UpdatedAt,
		AuditField.DeletedAt,
	}
}