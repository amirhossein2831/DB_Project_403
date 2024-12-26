package utils

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"reflect"
	"time"
)

var ShouldBePointerToStruct = errors.New("model must be a pointer to a struct")

// FillStructFromRows use when scan a list of record
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

// FillStructFromRowsWithJoin use when want ot scan list of record with one to one join
func FillStructFromRowsWithJoin(rows pgx.Rows, model interface{}) error {
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
			if field.Type() == reflect.TypeOf((*time.Time)(nil)) {
				// Directly add the *time.Time field to args
				args = append(args, new(*time.Time))
			} else {
				if tag := v.Type().Field(i).Tag.Get("sql"); tag != "-" {
					// If it's a pointer to another struct, initialize it and process subfields
					field.Set(reflect.New(field.Type().Elem()))
					subFields := field.Elem()
					for j := 0; j < subFields.NumField(); j++ {
						if sqlTag := subFields.Type().Field(j).Tag.Get("sql"); sqlTag != "" && sqlTag != "-" {
							args = append(args, subFields.Field(j).Addr().Interface())
						}
					}
				}
			}
		} else {
			if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "" && sqlTag != "-" {
				args = append(args, v.Field(i).Addr().Interface())
			}
		}
	}

	// Scan the row into the struct using the arguments
	if err := rows.Scan(args...); err != nil {
		return err
	}

	return nil
}

// FillStructFromRowsWithJoinMToM use when want ot scan a record or list of record with at list a one to many join
func FillStructFromRowsWithJoinMToM(rows pgx.Rows, model interface{}) error {
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
			if field.Type() == reflect.TypeOf((*time.Time)(nil)) {
				// Directly add the *time.Time field to args
				args = append(args, new(*time.Time))
			} else {
				// If the field is a pointer to a struct, initialize it and add its fields to args
				if tag := v.Type().Field(i).Tag.Get("sql"); tag != "-" {
					field.Set(reflect.New(field.Type().Elem()))
					subFields := field.Elem()
					for j := 0; j < subFields.NumField(); j++ {
						if sqlTag := subFields.Type().Field(j).Tag.Get("sql"); sqlTag != "" {
							args = append(args, subFields.Field(j).Addr().Interface())
						}
					}
				}
			}
		} else if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Ptr && field.Type().Elem().Elem().Kind() == reflect.Struct {
			// If the field is a slice of pointers to structs, initialize a new slice and add its fields to args
			if tag := v.Type().Field(i).Tag.Get("sql"); tag != "-" {
				sliceType := field.Type().Elem().Elem()
				newElem := reflect.New(sliceType).Elem()
				for j := 0; j < newElem.NumField(); j++ {
					if sqlTag := newElem.Type().Field(j).Tag.Get("sql"); sqlTag != "" {
						args = append(args, newElem.Field(j).Addr().Interface())
					}
				}
				field.Set(reflect.Append(field, newElem.Addr()))
			}
		} else {
			if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "" {
				args = append(args, v.Field(i).Addr().Interface())
			}
		}
	}

	// Scan the row into the struct using the arguments
	if err := rows.Scan(args...); err != nil {
		return err
	}

	return nil
}

// FillStructFromRow use when scan a record
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

// FillStructFromRowWithJoin use when scan a record with one to one join
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
			if field.Type() == reflect.TypeOf((*time.Time)(nil)) {
				// Directly add the *time.Time field to args
				args = append(args, new(*time.Time))
			} else {
				if tag := v.Type().Field(i).Tag.Get("sql"); tag != "-" {

					field.Set(reflect.New(field.Type().Elem()))
					subFields := field.Elem()
					for j := 0; j < subFields.NumField(); j++ {
						if sqlTag := subFields.Type().Field(j).Tag.Get("sql"); sqlTag != "" && sqlTag != "-" {
							args = append(args, subFields.Field(j).Addr().Interface())
						}
					}
				}
			}
		} else {
			if sqlTag := v.Type().Field(i).Tag.Get("sql"); sqlTag != "-" && sqlTag != "" {
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
