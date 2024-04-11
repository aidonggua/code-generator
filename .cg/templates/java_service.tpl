package {{package "."}}

import {{fullClassName "entity"}};
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * {{table "comment"}} 业务接口
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
public interface {{className "."}} extends IService<{{className "entity"}}> {
}
