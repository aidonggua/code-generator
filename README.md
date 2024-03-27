# code-generator
一个有意思的命令行代码生成器
- [x] 常用模版
- [x] 任务依赖+引用
- [x] 字符串转换指令
- [x] 类型转换指令
- [x] 日期指令
- [x] 导入上下文

## 使用方法
### 1. 编译（或者直接下载）code-generator，并将可执行文件放到系统环境变量PATH路径里
```shell
go build -o cg main.go
mv ./cg ~/bin/cg # 改成你的path路径
cg version # 测试并查看版本
```

### 2. 初始化.cg工作环境
在需要生成代码的项目根目录里使用 `cg init` 命令，会在项目根目录生成 `.cg` 文件夹，用于存放模板文件、配置文件以及输出的文件。
```shell
cd <your_project_directory>
cg init
```

.cg文件夹结构如下：
```text
.cg
├── config.yaml
├── output
│   ├── User.java
│   ├── UserMapper.java
│   └── UserMapper.xml
└── templates
    ├── java_entity.tpl
    ├── java_mapper.tpl
    └── java_mapper_xml.tpl
```
### 3. 自定义文件模板
init命令会默认生成常用模版文件，可以根据不同的项目需求，在 `.cg/templates` 文件夹下创建或者修改模板文件  

例如初始化命令自动生成的`java_entity.tpl` 文件，可用于生成java entity类
```text
package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};
{{""}}
{{- range imports}}
{{.}}
{{end -}}
{{""}}
import com.baomidou.mybatisplus.annotation.tableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

/**
* {{table "name"}}
*
* @Author {{config "author"}}
* @Date {{now}}
*/
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@TableName("{{table "name"}}")
@ApiModel(value="{{table "comment"}}实体类", description="{{table "name"}}")
public class {{table "name" | camelCase | title}} {
{{- range columns}}
    @ApiModelProperty(value = "{{.comment}}")
    @TableField("{{.name}}")
    private {{dbToJava .type}} {{camelCase .name}}{{";"}}
{{end -}}
}
```
### 4. 配置task
在 `.cg/config.yaml` 文件中配置task，例如生成entity文件的任务。
```yaml
mysql:
  username: root
  password: root
  host: 127.0.0.1
  port: 3306
  database: test
  table: user

author: melon
base-package: com.example
module: cg

tasks:
  # 生成java实体
  - name: JavaEntity
    template: java_entity.tpl
    output: .cg/output
    file-postfix: .java
    variables:
      sub-package: dao.domain
    enable: true
  # 生成java mapper类
  - name: JavaMapper
    template: java_mapper.tpl
    output: .cg/output
    file-postfix: Mapper.java
    variables:
      sub-package: dao.mapper
      class-postfix: Mapper
    enable: true
  # 生成java mybatis 的xml文件
  - name: JavaMapperXml
    template: java_mapper_xml.tpl
    output: .cg/output
    file-postfix: Mapper.xml
    enable: true
```
### 5. 生成代码
在需要生成代码的项目的根目录下执行 `cg run` 命令，会根据配置文件生成代码到`.cg/output`目录。
```shell
cg run
```