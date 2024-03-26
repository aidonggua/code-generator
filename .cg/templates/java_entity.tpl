package {{.Task.Variables.package}};

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;
{{""}}
{{- range .Imports}}
{{.}}
{{end -}}
{{""}}
/**
* {{.Table.Name}}
*
* @Author {{""}}
* @Date {{now}}
*/
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class {{titleCamelCase .Task.Table}} {
{{- range .Table.Columns}}
    /** {{.Comment}} */
    private {{dbToJava .Type}} {{camelCase .Name}}{{";" -}}
{{end}}
}