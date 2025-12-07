package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/r0kyi/glua/core"
)

type DataBase struct {
	driverName     string
	dataSourceName string

	db *sql.DB
}

func (db *DataBase) open() error {
	db_, err := sql.Open(db.driverName, db.dataSourceName)
	if err != nil {
		return err
	}
	db.db = db_

	return nil
}

func (db *DataBase) exec(query string, args []any) error {
	_, err := db.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) query(query string, args []any) ([]map[string]any, error) {
	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any
	for rows.Next() {
		values := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range values {
			ptrs[i] = &values[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		row := make(map[string]any, len(cols))
		for i, col := range cols {
			v := values[i]
			if b, ok := v.([]byte); ok {
				row[col] = core.B2S(b)
			} else {
				row[col] = v
			}
		}
		results = append(results, row)
	}

	return results, nil
}
