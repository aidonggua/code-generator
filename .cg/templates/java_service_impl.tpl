package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}};
import {{config "base-package"}}.{{config "module"}}.{{refs "JavaMapper" "sub-package"}}.{{table "name" | camelCase | title}}{{refs "JavaMapper" "class-postfix"}};
import {{config "base-package"}}.{{config "module"}}.{{refs "JavaService" "sub-package"}}.{{table "name" | camelCase | title}}{{refs "JavaService" "class-postfix"}};
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;

/**
 * {{table "comment"}} 业务实现类
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
@Service
public class {{table "name" | camelCase | title}}{{var "class-postfix"}} extends ServiceImpl<{{table "name" | camelCase | title}}{{refs "JavaMapper" "class-postfix"}}, {{table "name" | camelCase | title}}> implements {{table "name" | camelCase | title}}{{refs "JavaService" "class-postfix"}} {
}
