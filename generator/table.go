package generator

import "database/sql"

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name       string
	Type       string
	Collation  sql.NullString
	Null       string
	Key        string
	Default    sql.NullString
	Extra      string
	Privileges string
	Comment    string
}

func TableInfo(db *sql.DB, tableName string) *Table {
	rows, err := db.Query("SHOW FULL COLUMNS FROM " + tableName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	table := &Table{Name: tableName}
	for rows.Next() {
		column := Column{}
		err := rows.Scan(&column.Name, &column.Type, &column.Collation, &column.Null, &column.Key, &column.Default, &column.Extra, &column.Privileges, &column.Comment)
		if err != nil {
			panic(err)
		}
		table.Columns = append(table.Columns, column)
	}
	return table
}
