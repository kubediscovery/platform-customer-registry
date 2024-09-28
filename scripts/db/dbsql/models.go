// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbsql

import (
	"database/sql"
)

type PlatformLabDestroy struct {
	ID            string         `json:"id"`
	ProjectName   string         `json:"project_name"`
	Repository    string         `json:"repository"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Available     bool           `json:"available"`
	ErrorMessage  sql.NullString `json:"error_message"`
	TargetDestroy string         `json:"target_destroy"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at"`
}

type PlatformRegistry struct {
	ID          string `json:"id"`
	ProjectName string `json:"project_name"`
	Repository  string `json:"repository"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
	DeletedAt   string `json:"deleted_at"`
}
