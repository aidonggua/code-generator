# code-generator
一个有意思的代码生成器
- [x] 常用模版
- [x] 任务依赖+引用
- [x] 字符串转换指令
- [x] 类型转换指令
- [x] 日期指令
- [ ] 导入上下文

## 使用方法
### 1. 编译（或者直接下载）code-generator，并将可执行文件放到系统环境变量中
```shell
go build -o cg main.go
mv ./cg ~/bin/cg
cg version
```

### 2. 初始化.cg工作环境
在需要生成代码的项目根目录里使用 `cg init` 命令，会在项目根目录生成 `.cg` 文件夹，用于存放模板文件、配置文件以及输出的文件。
```shell
cd <your_project_directory>
cg init
```

.cg文件夹结构如下：
```text
tree .cg
├── config.yaml
├── output
│   ├── User.java
│   ├── UserMapper.java
│   ├── UserMapper.xml
│   └── user.go
└── templates
    ├── go
    │   └── go.entity.tpl
    └── java
        ├── entity.tpl
        ├── mapper.tpl
        └── mapper.xml.tpl
```
### 3. 自定义文件模板
init命令会默认生成常用模版文件，可以根据不同的项目需求，在 `.cg/templates` 文件夹下创建或者修改模板文件  

例如自动生成的`entity.tpl` 文件，可用于生成java entity类
```text
package {{ task.Variables.package }};

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class {{ transformer.Case.Title(transformer.Case.CamelCase(task.Table)) }} {
{% for column in table.Columns %}    /** {{ column.Comment }} */
    private {{ transformer.Type.DbToJava(column.Type) }} {{ transformer.Case.CamelCase(column.Name) }};
{% endfor %}}
```
### 4. 配置task
在 `.cg/config.yaml` 文件中配置task，例如生成entity文件的任务。
```yaml
# config file of code generator
mysql:
  username: root
  password: root
  host: 127.0.0.1
  port: 3306
  database: test
tasks:
  # 生成java实体
  - name: java_entity                                   # task name
    template: java/entity.tpl                           # template file from .cg/templates folder
    source-type: mysql                                  # table to entity
    table: user                                         # table name
    output: User.java                                   # output file name
    variables:                                          # variables for template
      package: com.example.dao.domain
    enable: true
  # 生成java mapper类
  - name: java_mapper
    template: java/mapper.tpl
    source-type: mysql
    table: user
    output: UserMapper.java
    variables:
      package: com.example.dao.mapper
    enable: true
  # 生成java mybatis 的xml文件
  - name: java_mapper_xml
    template: java/mapper.xml.tpl
    source-type: mysql
    table: user
    output: UserMapper.xml
    enable: true
  # 生成go实体
  - name: go_entity
    template: go/entity.tpl
    source-type: mysql
    table: user
    output: user.go
    variables:
      package: domain
    enable: true
```
### 5. 生成代码
在需要生成代码的项目的根目录下执行 `cg run` 命令，会根据配置文件生成代码到`.cg/output`目录。
```shell
cg run
```