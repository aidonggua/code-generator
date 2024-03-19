package transform

import "strings"

type TypeTransformer struct {
}

func (t *TypeTransformer) DbToJava(dbType string) string {
	if dbType == "tinyint" || strings.Contains(dbType, "tinyint") {
		return "Byte"
	}
	if dbType == "bigint" || strings.Contains(dbType, "bigint") {
		return "Long"
	}
	if dbType == "int" || strings.Contains(dbType, "bigint") {
		return "Integer"
	}
	if strings.Contains(dbType, "varchar") {
		return "String"
	}
	if dbType == "text" || strings.Contains(dbType, "text") {
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

func (t *TypeTransformer) DbToJDBC(dbType string) string {
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
