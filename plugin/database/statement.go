package database

import (
	"database/sql"

	"github.com/r0kyi/glua/core"
)

type Statement struct {
	stmt *sql.Stmt
}

func (s *Statement) query(args []any) ([]map[string]any, error) {
	rows, err := s.stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any
	for rows.Next() {
		values := make([]any, len(columns))
		ptrs := make([]any, len(columns))
		for i := range values {
			ptrs[i] = &values[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		row := make(map[string]any, len(columns))
		for i, col := range columns {
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

func (s *Statement) close() error {
	return s.stmt.Close()
}
