package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}};
import com.baomidou.mybatisplus.core.mapper.BaseMapper;

/**
 * {{table "comment"}} Mapper
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
public interface {{table "name" | camelCase | title}}{{var "class-postfix"}} extends BaseMapper<{{table "name" | camelCase | title}}> {

}