package gref

import (
	"errors"
	"github.com/zy1024/gref/utils"
	"reflect"
	"unicode"
)

// GenerateUpdateFields generates a map of fields to update based on the provided request object.
// Notice : This req must be pointers to structs
// It only includes fields that are not zero values and are exported (i.e., start with an uppercase letter).
// The field names are converted from camelCase to snake_case for use in the database.
// example: dataBase : {"name": "John", "age": 30} req : {"Age": 31} return : {"age": 31}
// gorm ues Updates func will update the age field to 31
func GenerateUpdateFields(req interface{}) (map[string]interface{}, error) {
	updateFields := make(map[string]interface{})

	reqValue := reflect.ValueOf(req)

	// check if req is a pointer
	if reqValue.Kind() != reflect.Ptr {
		return nil, errors.New("CopyFields : req must be pointers")
	}

	reqValue = reqValue.Elem()

	// check if the value is a struct
	if reqValue.Kind() != reflect.Struct {
		return nil, errors.New("CopyFields : req must be pointers to structs")
	}

	v := reflect.ValueOf(req).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Check if the field is exported (i.e., starts with an uppercase letter)
		if unicode.IsUpper(rune(fieldName[0])) {
			fieldValue := field.Interface()

			// Check if the field is not a zero value
			if !utils.IsZero(field) {
				// Convert the field name from camelCase to snake_case
				dbFieldName := utils.CamelToSnakeCase(fieldName)
				updateFields[dbFieldName] = fieldValue
			}
		}
	}

	return updateFields, nil
}
