package entity

import (
	"net/http"
	"time"

	"github.com/guregu/null"
	goCommonError "github.com/madevara24/go-common/errors"
)

const (
	// Table Name
	USERS_TABLE_NAME = "users"

	// error code
	ERR_CODE_USER_NOT_FOUND    = "USER_001"
	ERR_CODE_EMAIL_EXISTS      = "USER_002"
	ERR_CODE_EMAIL_REQUIRED    = "USER_003"
	ERR_CODE_PASSWORD_REQUIRED = "USER_004"
)

var (
	ErrUserNotFound     = goCommonError.NewErr(http.StatusNotFound, ERR_CODE_USER_NOT_FOUND, "user not found")
	ErrEmailExists      = goCommonError.NewErr(http.StatusBadRequest, ERR_CODE_EMAIL_EXISTS, "email already exists")
	ErrEmailRequired    = goCommonError.NewErr(http.StatusBadRequest, ERR_CODE_EMAIL_REQUIRED, "email is required")
	ErrPasswordRequired = goCommonError.NewErr(http.StatusBadRequest, ERR_CODE_PASSWORD_REQUIRED, "password is required")
)

type User struct {
	UUID      string      `json:"uuid" db:"uuid" primaryKey:"true"`
	Email     string      `json:"email" db:"email" updatable:"true"`
	Password  string      `json:"password" db:"password" updatable:"true"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at" updatable:"true"`
	DeletedAt null.Time   `json:"deleted_at" db:"deleted_at" updatable:"true"`
	CreatedBy string      `json:"created_by" db:"created_by"`
	UpdatedBy string      `json:"updated_by" db:"updated_by" updatable:"true"`
	DeletedBy null.String `json:"deleted_by" db:"deleted_by" updatable:"true"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrEmailRequired
	}

	if u.Password == "" {
		return ErrPasswordRequired
	}

	return nil
}
