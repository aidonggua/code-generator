package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaService" "sub-package"}}.{{table "name" | camelCase | title}}{{refs "JavaService" "class-postfix"}};
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;

/**
 * {{table "comment"}} 控制器
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
@Slf4j
@AllArgsConstructor
@RestController
@RequestMapping("/{{table "name" | kebabCase}}")
public class {{table "name" | camelCase | title}}{{var "class-postfix"}} {

    private {{table "name" | camelCase | title}}{{refs "JavaService" "class-postfix"}} {{table "name" | camelCase}}Service;
}
