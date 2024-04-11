package {{package "."}}

import {{fullClassName "entity"}};
import {{fullClassName "mapper"}};
import {{fullClassName "service"}};
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;

/**
 * {{table "comment"}} 业务实现类
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
@Slf4j
@AllArgsConstructor
@Service
public class {{className "."}} extends ServiceImpl<{{className "mapper"}}, {{className "entity"}}> implements {{className "service"}} {
}
