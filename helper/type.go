package helper

import "strings"

type TypeHelper struct {
}

func (t *TypeHelper) DbToJava(dbType string) string {
	if dbType == "tinyint" || strings.Contains(dbType, "tinyint") {
		return "Byte"
	}
	if dbType == "bigint" || strings.Contains(dbType, "bigint") {
		return "Long"
	}
	if dbType == "int" || strings.Contains(dbType, "bigint") {
		return "Integer"
	}
	if strings.Contains(dbType, "varchar") || dbType == "text" || strings.Contains(dbType, "text") {
		return "String"
	}
	if dbType == "decimal" || strings.Contains(dbType, "decimal") {
		return "BigDecimal"
	}
	if dbType == "timestamp" || dbType == "datetime" || dbType == "date" || dbType == "time" {
		return "Date"
	}
	return "Unknown"
}

func (t *TypeHelper) DbToJDBC(dbType string) string {
	if dbType == "tinyint" || strings.Contains(dbType, "tinyint") {
		return "TINYINT"
	}
	if dbType == "bigint" || strings.Contains(dbType, "bigint") {
		return "BIGINT"
	}
	if dbType == "int" || strings.Contains(dbType, "int") {
		return "INTEGER"
	}
	if strings.Contains(dbType, "varchar") {
		return "VARCHAR"
	}
	if dbType == "text" || strings.Contains(dbType, "text") {
		return "VARCHAR"
	}
	if dbType == "decimal" || strings.Contains(dbType, "decimal") {
		return "DECIMAL"
	}
	if dbType == "timestamp" || dbType == "datetime" {
		return "TIMESTAMP"
	}
	if dbType == "date" {
		return "DATE"
	}
	if dbType == "time" {
		return "TIME"
	}
	return "Unknown"
}

func (t *TypeHelper) DbToGo(dbType string) string {
	if dbType == "tinyint" || strings.Contains(dbType, "tinyint") || dbType == "bigint" || strings.Contains(dbType, "bigint") {
		return "int64"
	}
	if dbType == "int" || strings.Contains(dbType, "int") {
		return "int"
	}
	if strings.Contains(dbType, "varchar") || dbType == "text" || strings.Contains(dbType, "text") {
		return "string"
	}
	if dbType == "decimal" || strings.Contains(dbType, "decimal") {
		return "float64"
	}
	if dbType == "timestamp" || dbType == "datetime" || dbType == "date" || dbType == "time" {
		return "time.Time"
	}
	return "Unknown"
}
