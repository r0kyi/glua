package core

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

func MapToLTable(L *lua.LState, m interface{}) *lua.LTable {
	tbl := L.NewTable()
	convertMapToLTable(L, tbl, m)
	return tbl
}

func convertMapToLTable(L *lua.LState, tbl *lua.LTable, m interface{}) {
	if m == nil {
		return
	}

	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return
	}

	for _, key := range v.MapKeys() {
		k := fmt.Sprintf("%v", key.Interface())
		val := v.MapIndex(key).Interface()
		tbl.RawSetString(k, toLValue(L, val))
	}
}

func toLValue(L *lua.LState, val interface{}) lua.LValue {
	if val == nil {
		return lua.LNil
	}

	rv := reflect.ValueOf(val)

	switch rv.Kind() {
	case reflect.String:
		return lua.LString(rv.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return lua.LNumber(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return lua.LNumber(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return lua.LNumber(rv.Float())
	case reflect.Bool:
		return lua.LBool(rv.Bool())
	case reflect.Slice, reflect.Array:
		arr := L.NewTable()
		for i := 0; i < rv.Len(); i++ {
			arr.Append(toLValue(L, rv.Index(i).Interface()))
		}
		return arr
	case reflect.Map:
		tbl := L.NewTable()
		convertMapToLTable(L, tbl, val)
		return tbl
	case reflect.Ptr, reflect.Interface:
		if !rv.IsNil() {
			return toLValue(L, rv.Elem().Interface())
		}
		return lua.LNil
	default:
		return lua.LString(fmt.Sprintf("%v", val))
	}
}
