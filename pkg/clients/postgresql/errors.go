package postgresql

import (
	"errors"
	"fmt"
	"main/pkg/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ! Error:
// *Class 23 — Integrity Constraint Violation
// * 23000	integrity_constraint_violation
// * 23001	restrict_violation
// * 23502	not_null_violation
// * 23503	foreign_key_violation
// * 23505	unique_violation
// * 23514	check_violation
// * 23P01	exclusion_violation
// * 02000  no_data
func Error(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		pgErr = err.(*pgconn.PgError)
		if pgErr.Code == "23505" {
			return utils.ErrUniqueField
		}
		if pgErr.Code == "23503" {
			return utils.ErrInvalidForeinKey
		}
	} else if err == pgx.ErrNoRows {
		return utils.ErrNoEntity
	}
	return err
}

func Trace(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		pgErr = err.(*pgconn.PgError)
		return fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
	}
	return err
}
