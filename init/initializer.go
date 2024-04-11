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
  - name: entity
    template: java_entity.tpl
    prefix: ""
    postfix: ""
    file-type: .java
    properties:
      sub-package: dao.entity
    enable: true
  # 生成java mapper接口
  - name: mapper
    template: java_mapper.tpl
    prefix: ""
    postfix: Mapper
    file-type: .java
    properties:
      sub-package: dao.mapper
    enable: true
  # 生成java mybatis 的xml文件
  - name: xml
    template: java_mapper_xml.tpl
    prefix: ""
    postfix: Mapper
    file-type: .xml
    enable: true
  # 生成java service接口
  - name: service
    template: java_service.tpl
    prefix: ""
    postfix: Service
    file-type: .java
    properties:
      sub-package: service
    enable: true
  # 生成java services实现类
  - name: serviceImpl
    template: java_service_impl.tpl
    prefix: ""
    postfix: ServiceImpl
    file-type: .java
    properties:
      sub-package: service.impl
    enable: true
  # 生成java controller类
  - name: controller
    template: java_controller.tpl
    prefix: ""
    postfix: Controller
    file-type: .java
    properties:
      sub-package: controller
    enable: true
`

const javaEntityTpl = `package {{package "."}};
{{""}}
{{- range imports}}
{{.}}
{{end -}}
import com.baomidou.mybatisplus.annotation.TableField;
import com.baomidou.mybatisplus.annotation.TableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.*;

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
@ApiModel(value="{{table "comment"}}实体类", description="{{table "comment"}}实体类")
public class {{className "."}} {
{{- range columns}}
    @ApiModelProperty(value = "{{.comment}}")
    @TableField("{{.name}}")
    private {{dbToJava .type}} {{camelCase .name}}{{";"}}
{{end -}}
}`

const javaMapperTpl = `package {{package "."}};

import {{fullClassName "entity"}}
import com.baomidou.mybatisplus.core.mapper.BaseMapper;

/**
 * {{table "comment"}}dao接口
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
public interface {{className "."}} extends BaseMapper<{{className "entity"}}> {

}`

const javaMapperXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{fullClassName "mapper"}}">
    <resultMap id="BaseResultMap" type="{{fullClassName "entity"}}">
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

const javaServiceTpl = `package {{package "."}}

import {{fullClassName "entity"}};
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * {{table "comment"}} 业务接口
 *
 * @Author {{config "author"}}
 * @Date {{now}}
 */
public interface {{className "."}} extends IService<{{className "entity"}}> {
}`

const javaServiceImplTpl = `package {{package "."}}

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
}`

const javaControllerTpl = `package {{package "."}};

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
`

func (i *Initializer) Init() {
	// check if .cg folder exists
	if _, err := os.Stat(".cg"); err == nil {
		fmt.Print("The .cg folder already exists.")
		return
	}

	generator.FileWriter{}.CreateFolder(".cg/out")
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
