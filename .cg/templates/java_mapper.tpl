package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}};
import com.baomidou.mybatisplus.core.mapper.BaseMapper;

public interface {{table "name" | camelCase | title}}Mapper extends BaseMapper<{{table "name" | camelCase | title}}> {

}