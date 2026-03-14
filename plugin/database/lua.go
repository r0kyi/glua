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
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)

	return 1
}

func (db *DataBase) queryL(L *lua.LState) int {
	var args []any
	query := L.CheckString(1)
	for i := 2; i <= L.GetTop(); i++ {
		args = append(args, L.CheckAny(i))
	}

	rows, err := db.query(query, args)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tbl := core.MapToLTable(L, rows)
	L.Push(lua.LTrue)
	L.Push(tbl)

	return 2
}

func (db *DataBase) prepareL(L *lua.LState) int {
	query := L.CheckString(1)
	s, err := db.prepare(query)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(s)

	return 2
}

func (db *DataBase) closeL(L *lua.LState) int {
	err := db.close()
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)

	return 1
}

func (db *DataBase) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "exec":
		return L.NewFunction(db.execL)
	case "query":
		return L.NewFunction(db.queryL)
	case "prepare":
		return L.NewFunction(db.prepareL)
	case "close":
		return L.NewFunction(db.closeL)
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
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	db.db = db_

	L.Push(lua.LTrue)
	L.Push(db)

	return 2
}

func Preload() lua.LValue {
	db := &DataBase{}

	return db
}
