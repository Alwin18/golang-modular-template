package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// function ini jalan secara recursive
func GetJSONTagForField[T any](i T, fieldName string) string {
	val := reflect.ValueOf(i)
	typ := val.Type()

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	if val.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous {
			if tag := GetJSONTagForField(val.Field(i).Interface(), fieldName); tag != "" {
				return tag
			}
		} else {
			if field.Name == fieldName {
				return field.Tag.Get("json")
			}
		}

		if field.Type.Kind() == reflect.Struct {
			nestedTag := GetJSONTagForField(val.Field(i).Interface(), fieldName)
			if nestedTag != "" {
				return nestedTag
			}
		}

		if field.Type.Kind() == reflect.Slice || field.Type.Kind() == reflect.Array {
			for j := 0; j < val.Field(i).Len(); j++ {
				tag := GetJSONTagForField(val.Field(i).Index(j).Interface(), fieldName)
				if tag != "" {
					return tag
				}
			}
		}
	}

	return ""
}

// function ini jalan secara generic

func FormatValidationErrors[T any](err error, obj T) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			structNamespace := fieldErr.StructNamespace() // e.g., "KaryaTulisIlmiah[0].Judul"
			parts := strings.Split(structNamespace, ".")

			var jsonPath []string
			current := reflect.ValueOf(obj)
			if current.Kind() == reflect.Ptr {
				current = current.Elem()
			}

			for _, part := range parts {
				// Tangani index array: "KaryaTulisIlmiah[0]" => "KaryaTulisIlmiah"
				fieldName := part
				index := ""
				if strings.Contains(part, "[") {
					fieldName = part[:strings.Index(part, "[")]
					index = part[strings.Index(part, "["):]
				}

				if current.Kind() == reflect.Struct {
					t := current.Type()
					for i := 0; i < t.NumField(); i++ {
						f := t.Field(i)
						if f.Name == fieldName {
							jsonTag := f.Tag.Get("json")
							if jsonTag == "" || jsonTag == "-" {
								jsonTag = f.Name
							}

							// Tambahkan index jika ada
							jsonPath = append(jsonPath, jsonTag+index)

							current = current.Field(i)
							if current.Kind() == reflect.Ptr {
								current = current.Elem()
							}
							break
						}
					}
				}

				// Tangani jika current adalah slice
				if current.Kind() == reflect.Slice && index != "" {
					idx := 0
					fmt.Sscanf(index, "[%d]", &idx)
					if idx < current.Len() {
						current = current.Index(idx)
						if current.Kind() == reflect.Ptr {
							current = current.Elem()
						}
					}
				}
			}

			// Gabungkan path
			fieldDisplay := strings.Join(jsonPath, " -> ")

			if fieldErr.Tag() == "oneof" {
				errors = append(errors, fmt.Sprintf(
					"%s tidak sesuai, harus salah satu dari [%s]",
					fieldDisplay,
					fieldErr.Param(),
				))
			} else {
				errors = append(errors, fmt.Sprintf("%s tidak boleh kosong", fieldDisplay))
			}
		}
	}

	return errors
}
