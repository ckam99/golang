package postgresql

import (
	"errors"
	"fmt"
	"log"
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
		log.Fatalf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
		if pgErr.Code == "23505" {
			return utils.ErrUniqueField
		}
	} else if err == pgx.ErrNoRows {
		return utils.ErrNoEntity
	}
	return err
}

func ErrorCode(err error) (string, error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		pgErr = err.(*pgconn.PgError)
		return pgErr.Code, err
	}
	return "", err
}
