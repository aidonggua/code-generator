package initialize

import (
	"code-generator/generator"
	"fmt"
	"os"
)

type Initializer struct {
}

const configYaml = `mysql:
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
  # 生成java实体类
  - name: JavaEntity
    template: java_entity.tpl
    output: .cg/output
    file-postfix: .java
    variables:
      sub-package: dao.domain
    enable: true
  # 生成java mapper接口
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
  # 生成java service接口
  - name: JavaService
    template: java_service.tpl
    output: .cg/output
    file-postfix: Service.java
    variables:
      sub-package: service
      class-postfix: Service
    enable: true
  # 生成java services实现类
  - name: JavaServiceImpl
    template: java_service_impl.tpl
    output: .cg/output
    file-postfix: ServiceImpl.java
    variables:
      sub-package: service.impl
      class-postfix: ServiceImpl
    enable: true
  # 生成java controller类
  - name: JavaController
    template: java_controller.tpl
    output: .cg/output
    file-postfix: Controller.java
    variables:
      sub-package: controller
      class-postfix: Controller
    enable: true
`

const javaEntityTpl = `package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};
{{""}}
{{- range imports}}
{{.}}
{{end -}}
import com.baomidou.mybatisplus.annotation.tableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

/**
 * {{table "comment"}} 实体类
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
}`

const javaMapperTpl = `package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

import {{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}};
import com.baomidou.mybatisplus.core.mapper.BaseMapper;

/**
 * {{table "comment"}} Mapper
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
public interface {{table "name" | camelCase | title}}{{var "class-postfix"}} extends BaseMapper<{{table "name" | camelCase | title}}> {

}`

const javaMapperXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{config "base-package"}}.{{config "module"}}.{{refs "JavaMapper" "sub-package"}}.{{table "name" | camelCase | title}}{{refs "JavaMapper" "class-postfix"}}">
    <resultMap id="BaseResultMap" type="{{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}}">
    {{range columns -}}
        {{"    "}}<id column="{{.name}}" jdbcType="{{dbToJDBC .type}}" property="{{camelCase .name}}" />
    {{end -}}
    </resultMap>

    <sql id="Base_Column_List">
    {{"    " -}}
    {{range $i,$v := columns -}}
        {{if ne $i 0}},{{end}}{{.name -}}
    {{end}}
    </sql>
</mapper>`

const javaServiceTpl = `package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

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
`

const javaServiceImplTpl = `package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

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
`

const javaControllerTpl = `package {{config "base-package"}}.{{config "module"}}.{{var "sub-package"}};

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
`

func (i *Initializer) Init() {
	// check if .cg folder exists
	if _, err := os.Stat(".cg"); err == nil {
		fmt.Print("The .cg folder already exists.")
		return
	}

	generator.FileWriter{}.CreateFolder(".cg/output")
	generator.FileWriter{}.CreateFolder(".cg/templates")
	generator.FileWriter{}.Write(configYaml, ".cg/config.yaml")
	generator.FileWriter{}.Write(javaEntityTpl, ".cg/templates/java_entity.tpl")
	generator.FileWriter{}.Write(javaMapperTpl, ".cg/templates/java_mapper.tpl")
	generator.FileWriter{}.Write(javaMapperXmlTpl, ".cg/templates/java_mapper_xml.tpl")
	generator.FileWriter{}.Write(javaServiceTpl, ".cg/templates/java_service.tpl")
	generator.FileWriter{}.Write(javaServiceImplTpl, ".cg/templates/java_service_impl.tpl")
	generator.FileWriter{}.Write(javaControllerTpl, ".cg/templates/java_controller.tpl")
	fmt.Println("Initialized successfully.")
}
