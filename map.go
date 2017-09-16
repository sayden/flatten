package flatten

import (
	"fmt"
	"reflect"
)

func newMapIterator(i interface{}, name ...string) InterfaceIterator {
	var n string
	if len(name) == 1 {
		n = name[0]
	}

	mapI := &mapIter{
		prefix: n,
		s:      i,
		t:      reflect.TypeOf(i),
		v:      reflect.ValueOf(i),
	}

	mapI.keys = mapI.v.MapKeys()

	return mapI
}

type mapIter struct {
	s             interface{}
	cur           int
	t             reflect.Type
	v             reflect.Value
	prefix        string
	interfaceIter InterfaceIterator
	iterating     bool
	keys          []reflect.Value
}

func (st *mapIter) Next() (s string, i interface{}, finished bool) {
next:
	if st.cur >= len(st.keys) {
		return "", nil, true
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

	key := st.keys[st.cur]
	v := st.v.MapIndex(key)

	if st.prefix != "" {
		s = fmt.Sprintf("%s.%s", st.prefix, key.String())
	} else {
		s = key.String()
	}

	switch v.Kind() {
	case reflect.Map:
		st.interfaceIter = newMapIterator(v.Interface(), s)
		st.iterating = true
		goto iterating
	case reflect.Struct:
		numFields := v.NumField()
		var valid bool
		for i := 0; i < numFields; i++ {
			if v.Field(i).IsValid() {
				valid = true
			}
		}

		if !valid {
			i = v.Interface()
			st.cur++
			return
		}

		st.interfaceIter = newStructIterator(v.Interface(), s)
		st.iterating = true
		goto iterating

	default:
		if !v.IsValid() {
			st.cur++
			goto next
		}

		i = v.Interface()
		if i == nil {
			st.cur++
			goto next
		}

		if st.prefix != "" {
			s = fmt.Sprintf("%s.%s", st.prefix, key.String())
		} else {
			s = key.String()
		}
	}

	st.cur++

	return
}

func (st *mapIter) NextBytes() (s string, i []byte, finished bool) {
next:
	if st.cur >= len(st.keys) {
		return "", nil, true
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

	key := st.keys[st.cur]
	v := st.v.MapIndex(key)

	if st.prefix != "" {
		s = fmt.Sprintf("%s.%s", st.prefix, key.String())
	} else {
		s = key.String()
	}

	switch v.Kind() {
	case reflect.Map:
		st.interfaceIter = newMapIterator(v.Interface(), s)
		st.iterating = true
		goto iterating
	case reflect.Struct:
		numFields := v.NumField()
		var valid bool
		for i := 0; i < numFields; i++ {
			if v.Field(i).IsValid() {
				valid = true
			}
		}

		if !valid {
			i = bytes(v.Interface())
			st.cur++
			return
		}

		st.interfaceIter = newStructIterator(v.Interface(), s)
		st.iterating = true
		goto iterating

	default:
		if !v.IsValid() {
			st.cur++
			goto next
		}

		i = bytes(v.Interface())
		if i == nil {
			st.cur++
			goto next
		}

		if st.prefix != "" {
			s = fmt.Sprintf("%s.%s", st.prefix, key.String())
		} else {
			s = key.String()
		}
	}

	st.cur++

	return
}
