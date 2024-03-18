# code-generator
一个有意思的代码生成器
- [x] 表 ---> java entity
- [ ] 表 ---> mybatis mapper、mapper.xml
- [ ] 表 ---> service、serviceImpl
- [ ] 表 ---> controller
- [ ] 表 ---> vue页面

## 使用方法
### 1. 编译code-generator并将可执行文件放到系统环境变量中
```shell
go build -o cg main.go
```

### 2. 初始化.cg工作环境
在需要生成代码的项目根目录里使用 `cg init` 命令，会在项目根目录生成 `.cg` 文件夹，用于存放模板文件、配置文件以及输出的文件。
```shell
cg init
```

.cg文件夹结构如下：
```text
.cg
├── config.yaml
├── output
│   └── User.java
└── templates
    └── entity.tpl

```
### 3. 自定义文件模板
在 `.cg/templates` 文件夹下创建模板文件，例如创建一个 `entity.tpl` 文件，用于生成java entity的模板文件。
```text
package {{project.BasePackage}}.{{task.SubPackage}};

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class {{ transformer.Case.Title(transformer.Case.CamelCase(task.Source)) }} {
{% for column in table.Columns %}    /** {{column.Comment}} */
    private {{transformer.Type.DbToJava(column.Type)}} {{transformer.Case.CamelCase(column.Name)}};
{% endfor %}}
```
### 4. 配置task
在 `.cg/config.yaml` 文件中配置task，例如生成entity文件的任务。
```yaml
project:
  src-path: src/main/java
  base-package: com.sample
mysql:
  username: root
  password: root
  host: 127.0.0.1
  port: 3306
  database: test
tasks:
  - name: generate-entity                               # task name
    template: entity.tpl                                # template file from .cg/templates folder
    type: table_to_entity                               # table to entity
    sub-package: domain                                 # sub package
    source: user                                        # table name
    output: User.java                                   # output file name
    enable: true                                        # enable or disable the task
```
### 5. 生成代码
在需要生成代码的项目的根目录下执行 `cg run` 命令，会根据配置文件生成代码到`.cg/output`目录。
```shell
cg run
```