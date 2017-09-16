package flatten

import (
	"fmt"
	"reflect"
)

type structIter struct {
	s             interface{}
	cur           int
	t             reflect.Type
	v             reflect.Value
	prefix        string
	interfaceIter InterfaceIterator
	iterating     bool
}

func newStructIterator(i interface{}, name ...string) InterfaceIterator {
	var n string
	if len(name) == 1 {
		n = name[0]
	}

	return &structIter{
		prefix: n,
		s:      i,
		t:      reflect.TypeOf(i),
		v:      reflect.ValueOf(i),
	}
}

func (st *structIter) Next() (s string, i interface{}, finished bool) {
next:
	if st.cur == st.t.NumField() {
		finished = true
		return
	}

	//Iterating an inner struct
iterating:
	if st.iterating {
		var hasFinished bool
		s, i, hasFinished = st.interfaceIter.Next()

		if !hasFinished {
			return s, i, false
		} else {
			st.iterating = false
			st.cur++
			goto next
		}
	}

	//Maybe we can't extract the field
	val := st.v.Field(st.cur)
	if !val.CanInterface() {
		st.cur++
		goto next
	}
	ty := st.t.Field(st.cur)

	if st.prefix != "" {
		s = fmt.Sprintf("%s.%s", st.prefix, ty.Name)
	} else {
		s = ty.Name
	}

	//v could be an embedded type
	switch val.Kind() {
	case reflect.Map:
		st.interfaceIter = newMapIterator(val.Interface(), s)
		st.iterating = true
		goto iterating
	case reflect.Struct:
		numFields := val.NumField()
		var valid bool
		for i := 0; i < numFields; i++ {
			if val.Field(i).CanInterface() {
				valid = true
			}
		}

		if !valid {
			i = val.Interface()
			st.cur++
			return
		}

		st.interfaceIter = newStructIterator(val.Interface(), s)
		st.iterating = true
		goto iterating
	case reflect.Array:
		st.cur++
		goto next
	default:
		i = val.Interface()
	}

	st.cur++

	return
}


func (st *structIter) NextBytes() (s string, i []byte, finished bool) {
next:
	if st.cur == st.t.NumField() {
		finished = true
		return
	}

	//Iterating an inner struct
iterating:
	if st.iterating {
		var hasFinished bool
		s, i, hasFinished = st.interfaceIter.NextBytes()

		if !hasFinished {
			return s, i, false
		} else {
			st.iterating = false
			st.cur++
			goto next
		}
	}

	//Maybe we can't extract the field
	val := st.v.Field(st.cur)
	if !val.CanInterface() {
		st.cur++
		goto next
	}
	ty := st.t.Field(st.cur)

	if st.prefix != "" {
		s = fmt.Sprintf("%s.%s", st.prefix, ty.Name)
	} else {
		s = ty.Name
	}

	//v could be an embedded type
	switch val.Kind() {
	case reflect.Map:
		st.interfaceIter = newMapIterator(val.Interface(), s)
		st.iterating = true
		goto iterating
	case reflect.Struct:
		numFields := val.NumField()
		var valid bool
		for i := 0; i < numFields; i++ {
			if val.Field(i).CanInterface() {
				valid = true
			}
		}

		if !valid {
			i = bytes(val.Interface())
			st.cur++
			return
		}

		st.interfaceIter = newStructIterator(val.Interface(), s)
		st.iterating = true
		goto iterating
	case reflect.Array:
		st.cur++
		goto next
	default:
		i = bytes(val.Interface())
	}

	st.cur++

	return
}