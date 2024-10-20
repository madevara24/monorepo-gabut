package entities

import (
	"net/http"
	"time"

	"github.com/guregu/null"
	commonError "github.com/madevara24/go-common/errors"
)

type Task struct {
	UUID        string      `json:"uuid"`
	Title       string      `json:"title"`
	Description null.String `json:"description"`
	Deadline    null.Time   `json:"deadline"`
	Status      TaskStatus  `json:"status"`
	CreatedBy   string      `json:"created_by"`
	UpdatedBy   string      `json:"updated_by"`
	DeletedBy   null.String `json:"deleted_by"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   null.Time   `json:"deleted_at"`
}

const (
	// TABLE NAME
	TASK_TABLE_NAME = "tasks"

	// TASK STATUS
	TASK_STATUS_PENDING     TaskStatus = "PENDING"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN_PROGRESS"
	TASK_STATUS_COMPLETED   TaskStatus = "COMPLETED"

	// ERROR CODE
	ERROR_CODE_TASK_TITLE_REQUIRED   = "TASK_001"
	ERROR_CODE_TASK_STATUS_NOT_VALID = "TASK_002"
)

// ERRORS
var (
	ERROR_TASK_TITLE_REQUIRED   = commonError.NewErr(http.StatusBadRequest, ERROR_CODE_TASK_TITLE_REQUIRED, "title is required")
	ERROR_TASK_STATUS_NOT_VALID = commonError.NewErr(http.StatusBadRequest, ERROR_CODE_TASK_STATUS_NOT_VALID, "status is not valid")
)

type TaskStatus string

func (t *Task) Validate() error {
	if t.Title == "" {
		return ERROR_TASK_TITLE_REQUIRED
	}

	if t.Status != TASK_STATUS_PENDING && t.Status != TASK_STATUS_IN_PROGRESS && t.Status != TASK_STATUS_COMPLETED {
		return ERROR_TASK_STATUS_NOT_VALID
	}

	return nil
}
