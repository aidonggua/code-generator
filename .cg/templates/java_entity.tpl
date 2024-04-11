package {{package "."}};
{{""}}
{{- range imports}}
{{.}}
{{end -}}
import com.baomidou.mybatisplus.annotation.TableField;
import com.baomidou.mybatisplus.annotation.TableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.*;

/**
 * {{table "comment"}} 实体类
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@TableName("{{table "name"}}")
@ApiModel(value="{{table "comment"}}实体类", description="{{table "comment"}}实体类")
public class {{className "."}} {
{{- range columns}}
    @ApiModelProperty(value = "{{.comment}}")
    @TableField("{{.name}}")
    private {{dbToJava .type}} {{camelCase .name}}{{";"}}
{{end -}}
}