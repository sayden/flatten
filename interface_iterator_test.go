package flatten

import (
	"testing"

	"time"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"fmt"
)

type crazyType struct {
	CreatedAt     time.Time
	Name          string
	Surname       string
	Age           int
	IsTrue        bool
	ComposedType  composedType
	AnotherMap    map[string]string
	MyMap         map[string]interface{}
	Val           float64
	DifficultType map[string]struct {
		AKey   string
		AValue string
	}
	InnerMap map[string]map[string]interface{}
}

type composedType struct {
	Surname string
	//AnotherType   anotherType
	//DifficultType map[string]mapStruct
}

type mapStruct struct {
	AKey   string
	AValue string
}

type anotherType struct {
	Times int
}

var ct crazyType = crazyType{
	Name:      "Mario",
	Surname:   "Caster",
	CreatedAt: time.Now(),
	Age:       11,
	IsTrue:    true,
	AnotherMap: map[string]string{
		"Hello": "World",
	},
	MyMap: map[string]interface{}{
		"hello": "world",
		"nis":   "nas",
	},
	ComposedType: composedType{
		//AnotherType: anotherType{
		//	Times: 4,
		//},
		Surname: "QWE",
		//DifficultType: map[string]mapStruct{
		//	"qwer": {AKey: "qwer", AValue: "tyui"},
		//},
	},
	Val: 123.2,
	//DifficultType: map[string]struct {
	//	AKey   string
	//	AValue string
	//}{
	//	"DifficultTypeMapKey": {AKey: "a_key", AValue: "a_value"},
	//	"DifficultTypeMapKey2": {AKey: "a_key2", AValue: "a_value2"},
	//},
	//InnerMap: map[string]map[string]interface{}{
	//	"inner":  {"masdad": 1, "qwerqer": "hello"},
	//	"inner2": {"masdad": 1, "qwerqer": "hello"},
	//},
}

type testCrazy struct {
	CreatedAt string
	crazyType
}

func Test_InterfaceIterator(t *testing.T) {

	si := NewIterator(ct)

	name, val, finished := si.Next()

	result := make(map[string]interface{})

	for !finished {
		result[name] = val
		name, val, finished = si.Next()
	}

	assert.Equal(t, "Mario", result["Name"])
	assert.Equal(t, 11, result["Age"])
	assert.Equal(t, "world", result["MyMap.hello"])
	assert.Equal(t, "nas", result["MyMap.nis"])
	//assert.Equal(t, 4, result["ComposedType.AnotherType.Times"])
	assert.Equal(t, "QWE", result["ComposedType.Surname"])
	//assert.Equal(t, "qwer", result["ComposedType.DifficultType.qwer.AKey"])
	//assert.Equal(t, "tyui", result["ComposedType.DifficultType.qwer.AValue"])
	//assert.Equal(t, "a_key", result["DifficultType.DifficultTypeMapKey.AKey"])
	//assert.Equal(t, "a_key2", result["DifficultType.DifficultTypeMapKey2.AKey"])
	//assert.Equal(t, "a_value", result["DifficultType.DifficultTypeMapKey.AValue"])
	//assert.Equal(t, "a_value2", result["DifficultType.DifficultTypeMapKey2.AValue"])
	//assert.Equal(t, 1, result["InnerMap.inner.masdad"])
	//assert.Equal(t, "hello", result["InnerMap.inner.qwerqer"])
}

func Test_Unflatten(t *testing.T) {
	si := NewIterator(ct)

	name, val, finished := si.NextBytes()
	result := make(map[string][]byte)

	for !finished {
		result[name] = val
		name, val, finished = si.NextBytes()
	}

	newCt := testCrazy{}
	UnflattenBytes(result, &newCt)

	pretty.Println(newCt)
}
