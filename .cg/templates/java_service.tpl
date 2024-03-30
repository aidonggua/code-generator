package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}};
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * {{table "comment"}} 业务接口
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
public interface {{table "name" | camelCase | title}}{{var "class-postfix"}} extends IService<{{table "name" | camelCase | title}}> {
}
