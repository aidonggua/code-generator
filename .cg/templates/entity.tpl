package {{project.BasePackage}}.{{task.SubPackage}};

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class {{ transformer.Case.Title(transformer.Case.CamelCase(task.Source)) }} {
{% for column in table.Columns %}    /** {{column.Comment}} */
    private {{transformer.Type.DbToJava(column.Type)}} {{transformer.Case.CamelCase(column.Name)}};
{% endfor %}}