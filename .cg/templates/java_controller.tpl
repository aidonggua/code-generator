package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaService" "sub-package"}}.{{table "name" | camelCase | title}}{{refs "JavaService" "class-postfix"}};

/**
 * {{table "comment"}} 控制器
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
@RestController
@RequestMapping("/{{table "name" | kebabCase}}")
public class {{table "name" | camelCase | title}}{{var "class-postfix"}} {
    private {{table "name" | camelCase | title}}{{refs "JavaService" "class-postfix"}} {{table "name" | camelCase}}Service;
}
