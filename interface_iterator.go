package flatten

import "reflect"

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
