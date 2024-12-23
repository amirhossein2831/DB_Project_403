package utils

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"reflect"
)

var ShouldBePointerToStruct = errors.New("model must be a pointer to a struct")

func FillStructFromRows(row pgx.Rows, model interface{}) error {
	// Ensure the model is a pointer to a struct
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ShouldBePointerToStruct
	}

	// Prepare a slice to hold pointers to the struct fields
	v = v.Elem()
	args := make([]interface{}, 0)

	// Iterate over each field of the struct and create a pointer to the field
	for i := 0; i < v.NumField(); i++ {
		if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
			args = append(args, v.Field(i).Addr().Interface())
		}
	}

	// Scan the row into the struct using the arguments
	if err := row.Scan(args...); err != nil {
		return err
	}

	return nil
}

func FillStructFromRowsWithJoin(row pgx.Rows, model interface{}) error {
	// Ensure the model is a pointer to a struct
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ShouldBePointerToStruct
	}

	// Prepare a slice to hold pointers to the struct fields
	v = v.Elem()
	args := make([]interface{}, 0)

	// Iterate over each field of the struct and create a pointer to the field
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			// If the field is a pointer to a struct, initialize it and add its fields to args
			field.Set(reflect.New(field.Type().Elem()))
			subFields := field.Elem()
			for j := 0; j < subFields.NumField(); j++ {
				if sqlTag := subFields.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
					args = append(args, subFields.Field(j).Addr().Interface())
				}
			}
		} else {
			if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
				args = append(args, v.Field(i).Addr().Interface())
			}
		}
	}

	// Scan the row into the struct using the arguments
	if err := row.Scan(args...); err != nil {
		return err
	}

	return nil
}

func FillStructFromRow(row pgx.Row, model interface{}) error {
	// Ensure the model is a pointer to a struct
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ShouldBePointerToStruct
	}

	// Prepare a slice to hold pointers to the struct fields
	v = v.Elem()
	args := make([]interface{}, 0)

	// Iterate over each field of the struct and create a pointer to the field
	for i := 0; i < v.NumField(); i++ {
		if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
			args = append(args, v.Field(i).Addr().Interface())
		}
	}

	// Scan the row into the struct using the arguments
	if err := row.Scan(args...); err != nil {
		return err
	}

	return nil
}

func FillStructFromRowWithJoin(row pgx.Row, model interface{}) error {
	// Ensure the model is a pointer to a struct
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ShouldBePointerToStruct
	}

	// Prepare a slice to hold pointers to the struct fields
	v = v.Elem()
	args := make([]interface{}, 0)

	// Iterate over each field of the struct and create a pointer to the field
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			// If the field is a pointer to a struct, initialize it and add its fields to args
			field.Set(reflect.New(field.Type().Elem()))
			subFields := field.Elem()
			for j := 0; j < subFields.NumField(); j++ {
				if sqlTag := subFields.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
					args = append(args, subFields.Field(j).Addr().Interface())
				}
			}
		} else {
			if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
				args = append(args, v.Field(i).Addr().Interface())
			}
		}
	}

	// Scan the row into the struct using the arguments
	if err := row.Scan(args...); err != nil {
		return err
	}

	return nil
}
