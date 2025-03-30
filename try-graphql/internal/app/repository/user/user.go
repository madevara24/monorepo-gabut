package user

import (
	"context"
	"database/sql"
	"try-graphql/internal/app/entity"

	"github.com/jmoiron/sqlx"

	"github.com/madevara24/go-common/mapper"
	"github.com/madevara24/go-common/txmanager"
)

func (r *repo) Create(ctx context.Context, user entity.User) error {
	var (
		query string
		data  []interface{}
		tx    *sqlx.Tx
		stmt  *sql.Stmt
		err   error
	)
	query, data, err = r.datasource.Mapper.Insert(ctx, user, entity.USERS_TABLE_NAME)
	if err != nil {
		return err
	}

	tx, _ = txmanager.ExtractTx(ctx)

	if tx != nil {
		stmt, err = tx.PrepareContext(ctx, tx.Rebind(query))
	} else {
		stmt, err = r.datasource.Postgre.PrepareContext(ctx, r.datasource.Postgre.Rebind(query))
	}

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, data...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	query, err := r.datasource.Mapper.Search(ctx, user, entity.USERS_TABLE_NAME, mapper.SearchFilter{
		Where: " email = $1 ",
	})
	if err != nil {
		return entity.User{}, err
	}

	err = r.datasource.Postgre.QueryRowxContext(ctx, query, email).StructScan(&user)
	if err == sql.ErrNoRows {
		return entity.User{}, entity.ErrUserNotFound
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
