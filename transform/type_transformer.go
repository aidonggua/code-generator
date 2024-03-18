package transform

import "strings"

type TypeTransformer struct {
}

func (t *TypeTransformer) DbToJava(dbType string) string {
	if dbType == "tinyint(1)" {
		return "Boolean"
	}
	if dbType == "tinyint" || strings.Contains(dbType, "tinyint") {
		return "Byte"
	}
	if dbType == "int" {
		return "Integer"
	}
	if strings.Contains(dbType, "varchar") {
		return "String"
	}
	if dbType == "text" || strings.Contains(dbType, "text") {
		return "String"
	}
	if dbType == "decimal" {
		return "BigDecimal"
	}
	if dbType == "bigint" || strings.Contains(dbType, "bigint") {
		return "Long"
	}
	if dbType == "timestamp" || dbType == "datetime" || dbType == "date" || dbType == "time" {
		return "Date"
	}
	return "Unknown"
}
