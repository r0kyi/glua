package database

import (
	"database/sql"
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (db *DataBase) String() string {
	return fmt.Sprintf("glua.database: %p", db)
}

func (db *DataBase) Type() lua.LValueType {
	return lua.LTObject
}

func (db *DataBase) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newDataBaseL), true
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

	tbl := core.MapToLTable(L, rows)
	L.Push(tbl)
	L.Push(lua.LNil)

	return 2
}

func (db *DataBase) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "exec":
		return L.NewFunction(db.execL)
	case "query":
		return L.NewFunction(db.queryL)
	default:
		return lua.LNil
	}
}

func newDataBaseL(L *lua.LState) int {
	db := &DataBase{}

	tbl := L.CheckTable(1)
	_ = core.LTableToStrut(tbl, db)

	db_, err := sql.Open(db.DriverName, db.DataSourceName)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	db.db = db_

	L.Push(db)
	L.Push(lua.LNil)

	return 2
}

func Preload() lua.LValue {
	db := &DataBase{}

	return db
}
