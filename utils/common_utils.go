package utils

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	AdAccount string
	WorkCode  string
	Nickname  string
}

func (*UserInfo) getWorkCode(info *UserInfo) string {
	work_code := info.WorkCode
	return work_code
}

func ArrayUnit() {
	var aa = []int{1, 2, 4}

	for key, v := range aa {
		fmt.Println("----", key, v)
	}
	fmt.Println(aa)
}

// 深度拷贝
func DeepCopy(source interface{}) interface{} {
	sourceValue := reflect.ValueOf(source)
	if sourceValue.Kind() != reflect.Struct {
		return nil
	}
	newValue := reflect.New(sourceValue.Type()).Elem()
	for i := 0; i < sourceValue.NumField(); i++ {
		field := newValue.Field(i)
		fieldValue := sourceValue.Field(i)
		if fieldValue.Kind() == reflect.Ptr {
			if !field.CanSet() || fieldValue.IsNil() {
				continue
			}
			field.Set(reflect.New(field.Type().Elem()))
			field.Elem().Set(reflect.Indirect(fieldValue).Addr())
		} else if fieldValue.Kind() == reflect.Struct {
			subStructCopy := DeepCopy(fieldValue.Interface())
			if subStructCopy != nil {
				field.Set(reflect.ValueOf(subStructCopy))
			}
		} else {
			if field.CanSet() && fieldValue.IsValid() {
				field.Set(fieldValue)
			}
		}
	}

	return newValue.Interface()
}
