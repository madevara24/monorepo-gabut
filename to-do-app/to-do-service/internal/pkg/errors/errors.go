package errors

import (
	"net/http"

	commonError "github.com/madevara24/go-common/errors"
)

var (
	// TODO: move to common error
	ERROR_INVALID_UUID = commonError.NewErr(http.StatusBadRequest, commonError.ErrorInvalidParameter, "invalid uuid format")
)
