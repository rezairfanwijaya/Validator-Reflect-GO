package main

import (
	"errors"
	"fmt"
	"reflect"
)

func IsValid(data interface{}) (bool, error) {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			required := reflect.ValueOf(data).Field(i).Interface()
			if required == "" || required == 0 {
				msg := fmt.Sprintf("Field %v, cannot be empty", t.Field(i).Name)
				return false, errors.New(msg)
			}
		}
	}

	return true, nil
}
