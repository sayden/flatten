package flatten

import (
	"fmt"
	"reflect"
	"strings"
)

type InterfaceIterator interface {
	Next() (string, interface{}, bool)
	NextBytes() (string, []byte, bool)
}

func NewIterator(i interface{}, name ...string) InterfaceIterator {

	switch reflect.TypeOf(i).Kind() {
	case reflect.Struct:
		return newStructIterator(i, name...)
	case reflect.Map:
		return newMapIterator(i, name...)
	}

	return nil
}

func UnflattenBytes(result map[string][]byte, i interface{}) {
	for k, v := range result {
		unflattenBytes(k, v, i)
	}
}

func unflattenBytes(inK string, v []byte, i interface{}) {
	rootValue := reflect.ValueOf(i)

	if strings.Contains(inK, ".") {
		//Struct or map type
		keys := strings.Split(inK, ".")
		var field reflect.Value = rootValue

		for _, k := range keys {
			if strings.Contains(inK, "InnerMap") {
				fmt.Println(k)
			}

			switch field.Kind() {
			case reflect.Ptr:
				field = field.Elem().FieldByName(k)

			case reflect.Struct:
				field = field.FieldByName(k)

			case reflect.Map:
				if field.IsNil() && field.CanSet() {
					m := reflect.MakeMap(field.Type())
					field.Set(m)
				}

				var value reflect.Value
				switch field.Type().Elem().Kind() {
				case reflect.Interface:
					value = reflect.ValueOf(v)
				case reflect.Struct:
					//fmt.Println(field.Elem().FieldByName())
					fmt.Println("STRUCT", inK)
					//value = getValue(field.Type().Elem().Kind(), v)
				case reflect.Map:
					fmt.Println("MAP", k)
					if field.IsNil() && field.CanSet() {
						m := reflect.MakeMap(field.Type())
						field.Set(m)
					}
					value = getValue(field.Type().Elem().Kind(), v)

					field.SetMapIndex(reflect.ValueOf(k), value)
				default:
					value = getValue(field.Type().Elem().Kind(), v)
				}

				field.SetMapIndex(reflect.ValueOf(k), value)
			}

			var value reflect.Value = getValue(field.Kind(), v)
			if value.IsValid() {
				field.Set(value)
			}
		}
	}

	field := rootValue.
		Elem().
		FieldByName(inK)

	if !field.CanSet() {
		return
	}

	var value reflect.Value
	value = getValue(field.Kind(), v)

	if value.IsValid() {
		field.Set(value)
	}
}

func getValue(k reflect.Kind, v []byte) (value reflect.Value) {
	switch k {
	case reflect.Float32:
		value = reflect.ValueOf(float32FromBytes(v))

	case reflect.Float64:
		value = reflect.ValueOf(float64FromBytes(v))

	case reflect.String:
		value = reflect.ValueOf(string(v))

	case reflect.Bool:
		value = reflect.ValueOf(v[0] == 1)

	case reflect.Uint:
		value = reflect.ValueOf(uintFromBytes(v))
	case reflect.Uint8:
		value = reflect.ValueOf(uint8FromBytes(v))
	case reflect.Uint16:
		value = reflect.ValueOf(uint16FromBytes(v))
	case reflect.Uint32:
		value = reflect.ValueOf(uint32FromBytes(v))
	case reflect.Uint64:
		value = reflect.ValueOf(uint64FromBytes(v))

	case reflect.Int:
		value = reflect.ValueOf(intFromBytes(v))
	case reflect.Int8:
		value = reflect.ValueOf(int8FromBytes(v))
	case reflect.Int16:
		value = reflect.ValueOf(int16FromBytes(v))
	case reflect.Int32:
		value = reflect.ValueOf(int32FromBytes(v))
	case reflect.Int64:
		value = reflect.ValueOf(int64FromBytes(v))
	}

	return
}
