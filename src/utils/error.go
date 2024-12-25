package utils

import (
	"errors"
	"github.com/jackc/pgconn"
)

func IsErrorCode(err error, code string) bool {
	if err == nil {
		return false
	}
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return false
	}

	return pgErr.Code == code
}

func GetErrorConstraintName(err error) string {
	if err == nil {
		return ""
	}
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return ""
	}

	return pgErr.ConstraintName
}
