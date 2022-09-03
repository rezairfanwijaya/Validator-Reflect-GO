package validator_reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func panicHandling() {
	msg := recover()
	if msg != nil {
		fmt.Printf("%v [Params Must Be Single Struct]\n", msg)
	}
}

func Validate(data interface{}) (bool, error) {
	defer panicHandling()

	t := reflect.TypeOf(data)
	max, err := checkData(t.NumField())

	if err != nil {
		panic("params must be struct")
	}

	for i := 0; i < max; i++ {
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

func checkData(data int) (int, error) {
	numberOfFiled := data
	if numberOfFiled <= 0 {
		return numberOfFiled, errors.New("params must be struct")
	}
	return numberOfFiled, nil
}
