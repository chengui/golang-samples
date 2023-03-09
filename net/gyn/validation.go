package gyn

import (
	"errors"
	"reflect"
	"strings"
)

func Validate(o interface{}) (err error) {
	typ, val := reflect.TypeOf(o), reflect.ValueOf(o)
	if typ.Kind() == reflect.Ptr {
		typ, val = typ.Elem(), val.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if strings.Contains(field.Tag.Get("binding"), "required") {
			fval := val.Field(i).Interface()
			zero := reflect.Zero(field.Type).Interface()
			if reflect.DeepEqual(zero, fval) {
				name := field.Name
				if n := field.Tag.Get("json"); n != "" {
					name = n
				}
				err = errors.New("Required " + name)
				return
			}
		}
	}
	return
}
