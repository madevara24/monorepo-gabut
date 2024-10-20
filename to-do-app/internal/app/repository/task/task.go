package task

import (
	"context"
	"database/sql"
	"to-do-app/internal/app/entities"
	internalError "to-do-app/internal/pkg/errors"

	"github.com/google/uuid"
)

func (r *repo) Create(ctx context.Context, task entities.Task) error {

	query := `
		INSERT INTO tasks (uuid, title, deadline, description, status, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	stmt, err := r.datasource.Postgre.PrepareContext(ctx, r.datasource.Postgre.Rebind(query))
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, task.UUID, task.Title, task.Deadline, task.Description, task.Status, task.CreatedBy, task.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetByUUID(ctx context.Context, taskUUID string) (entities.Task, error) {
	// Validate UUID
	_, err := uuid.Parse(taskUUID)
	if err != nil {
		return entities.Task{}, internalError.ERROR_INVALID_UUID
	}
	query := `
		SELECT t."uuid", t."title", t."deadline", t."description", t."status", t."created_by", t."updated_by", t."deleted_by", t."created_at", t."updated_at", t."deleted_at"
		FROM tasks t
		WHERE t."uuid" = $1::uuid
	`

	var task entities.Task
	err = r.datasource.Postgre.QueryRowxContext(ctx, query, taskUUID).StructScan(&task)
	if err == sql.ErrNoRows {
		return entities.Task{}, entities.ERROR_TASK_NOT_FOUND
	}

	if err != nil {
		return entities.Task{}, err
	}
	return task, nil
}
