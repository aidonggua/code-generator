package {{.Task.Variables.package}};

type {{title .Table.Name}} struct {
{{- range .Table.Columns}}
    {{titleCamelCase .Name}} {{dbToGo .Type}} `json:"{{.Name}}"` // {{.Comment -}}
{{end}}
}

func (*{{title .Table.Name}}) TableName() string {
	return "{{.Table.Name}}"
}