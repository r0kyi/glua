package database

import (
	"fmt"

	. "github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (db *DataBase) String() string {
	return fmt.Sprintf("glua.database: %p", db)
}

func (db *DataBase) AssertFunction() lua.LGFunction {
	return nil
}

func (db *DataBase) MetatableName() string {
	return "lua.table.database"
}

func (db *DataBase) openL(L *lua.LState) int {
	driverName := L.CheckString(1)
	dataSourceName := L.CheckString(2)
	db.driverName = driverName
	db.dataSourceName = dataSourceName

	err := db.open()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	ud := NewUserData(L, db)
	L.Push(ud)
	L.Push(lua.LNil)
	return 2
}

func (db *DataBase) execL(L *lua.LState) int {
	var args []any
	query := L.CheckString(1)
	for i := 2; i <= L.GetTop(); i++ {
		args = append(args, L.CheckString(i))
	}

	err := db.exec(query, args)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	L.Push(lua.LNil)
	return 1
}

func (db *DataBase) queryL(L *lua.LState) int {
	var args []any
	query := L.CheckString(1)
	for i := 2; i <= L.GetTop(); i++ {
		args = append(args, L.CheckString(i))
	}

	rows, err := db.query(query, args)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	t := SliceMapToLTable(L, rows)
	L.Push(t)
	L.Push(lua.LNil)
	return 2
}

func (db *DataBase) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "open":
		return L.NewFunction(db.openL)
	case "exec":
		return L.NewFunction(db.execL)
	case "query":
		return L.NewFunction(db.queryL)
	default:
		return lua.LNil
	}
}

func Preload(L *lua.LState) lua.LValue {
	db := &DataBase{}
	ud := NewUserData(L, db)

	return ud
}
