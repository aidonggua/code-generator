package {{package "."}};

import {{fullClassName "service"}}
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
public class {{className "."}} {

    private {{className "service"}};
}
