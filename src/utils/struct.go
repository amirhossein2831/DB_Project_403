package utils

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"reflect"
)

var ShouldBePointerToStruct = errors.New("model must be a pointer to a struct")

func FillStructFromRow(row pgx.Rows, model interface{}) error {
	// Ensure the model is a pointer to a struct
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ShouldBePointerToStruct
	}

	// Prepare a slice to hold pointers to the struct fields
	v = v.Elem()
	args := make([]interface{}, v.NumField())

	// Iterate over each field of the struct and create a pointer to the field
	for i := 0; i < v.NumField(); i++ {
		args[i] = v.Field(i).Addr().Interface()
	}

	// Scan the row into the struct using the arguments
	if err := row.Scan(args...); err != nil {
		return err
	}

	return nil
}
