package flatten

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type crazyType struct {
	CreatedAt     time.Time
	Name          string
	Age           int
	IsTrue        bool
	ComposedType  composedType
	MyMap         map[string]interface{}
	Val           float64
	DifficultType map[string]struct {
		AKey   string
		AValue string
	}
	InnerMap map[string]map[string]interface{}
}

type composedType struct {
	Surname       string
	AnotherType   anotherType
	DifficultType map[string]mapStruct
}

type mapStruct struct {
	AKey   string
	AValue string
}

type anotherType struct {
	Times int
}

func Test_InterfaceIterator(t *testing.T) {
	ct := crazyType{
		Name:      "Mario",
		CreatedAt: time.Now(),
		Age:       11,
		IsTrue:    true,
		MyMap: map[string]interface{}{
			"hello": "world",
			"nis":   "nas",
		},
		ComposedType: composedType{
			AnotherType: anotherType{
				Times: 4,
			},
			Surname: "QWE",
			DifficultType: map[string]mapStruct{
				"qwer": {AKey: "qwer", AValue: "tyui"},
			},
		},
		Val: 123.2,
		DifficultType: map[string]struct {
			AKey   string
			AValue string
		}{
			"Adsf": {AKey: "a_key", AValue: "a_value"},
		},
		InnerMap: map[string]map[string]interface{}{
			"inner":  {"masdad": 1, "qwerqer": "hello"},
			"inner2": {"masdad": 1, "qwerqer": "hello"},
		},
	}

	si := newStructIterator(ct)

	name, val, finished := si.Next()
	i := 0

	result := make(map[string]interface{})

	for !finished {
		i++
		result[name] = val
		name, val, finished = si.Next()
	}

	assert.Equal(t, 17, i)
	assert.Equal(t, "Mario", result["Name"])
	assert.Equal(t, 11, result["Age"])
	assert.Equal(t, "world", result["MyMap.hello"])
	assert.Equal(t, "nas", result["MyMap.nis"])
	assert.Equal(t, 4, result["ComposedType.AnotherType.Times"])
	assert.Equal(t, "QWE", result["ComposedType.Surname"])
	assert.Equal(t, "qwer", result["ComposedType.DifficultType.qwer.AKey"])
	assert.Equal(t, "tyui", result["ComposedType.DifficultType.qwer.AValue"])
	assert.Equal(t, "a_key", result["DifficultType.Adsf.AKey"])
	assert.Equal(t, "a_value", result["DifficultType.Adsf.AValue"])
	//
	//for k, v := range result {
	//	fmt.Printf("%s -> %v\n", k, v)
	//}
}
