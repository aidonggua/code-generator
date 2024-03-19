package {{ task.Variables.package }};

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{ refs.entity.Variables.package }}.{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }};

public interface {{ transformer.Case.Title(transformer.Case.CamelCase(task.Table)) }}Mapper extends BaseMapper<{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }}>{
}