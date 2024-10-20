package task

import (
	"context"
	"to-do-app/internal/app/entities"
)

func (r *repo) Create(ctx context.Context, task entities.Task) error {

	query := `
		INSERT INTO tasks (uuid, title, deadline, description, status, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	// INSERT TO DB
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
