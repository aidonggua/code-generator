package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};
{{""}}
{{- range imports}}
{{.}}
{{end -}}
import com.baomidou.mybatisplus.annotation.tableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

/**
* {{table "name"}}
*
* @Author {{config "author"}}
* @Date {{now}}
*/
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@TableName("{{table "name"}}")
@ApiModel(value="{{table "comment"}}实体类", description="{{table "name"}}")
public class {{table "name" | camelCase | title}} {
{{- range columns}}
    @ApiModelProperty(value = "{{.comment}}")
    @TableField("{{.name}}")
    private {{dbToJava .type}} {{camelCase .name}}{{";"}}
{{end -}}
}