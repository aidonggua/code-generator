package {{.Task.Variables.package}};

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.Refs.JavaEntity.Variables.package}}.{{titleCamelCase .Table.Name}};

public interface {{titleCamelCase .Table.Name}}Mapper extends BaseMapper<{{titleCamelCase .Table.Name}}> {

}