package errors

import (
	"net/http"

	goCommonError "github.com/madevara24/go-common/errors"
)

const (
	// hash
	ERR_CODE_PASSWORD_HASH_FAILED = "HASH_001"
)

var (
	ErrPasswordHashFailed = goCommonError.NewErr(http.StatusInternalServerError, ERR_CODE_PASSWORD_HASH_FAILED, "password hash failed")
)
