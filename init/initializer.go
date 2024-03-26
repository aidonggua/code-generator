package initialize

import (
	"code-generator/generator"
	"fmt"
	"os"
)

type Initializer struct {
}

const configYaml = `# config file of code generator
global:
  author: melon
mysql:
  username: root
  password: root
  host: 127.0.0.1
  port: 3306
  database: test
tasks:
  # 生成java实体
  - name: JavaEntity                                    # task name
    template: java_entity.tpl                           # template file from .cg/templates folder
    source-type: mysql                                  # table to entity
    table: user                                         # table name
    output: User.java                                   # output file name
    variables:                                          # variables for template
      package: com.example.dao.domain
    enable: true
  # 生成java mapper类
  - name: JavaMapper
    template: java_mapper.tpl
    source-type: mysql
    table: user
    output: UserMapper.java
    variables:
      package: com.example.dao.mapper
    enable: true
  # 生成java mybatis 的xml文件
  - name: JavaMapperXml
    template: java_mapper_xml.tpl
    source-type: mysql
    table: user
    output: UserMapper.xml
    enable: true
`

const javaEntityTpl = `package {{.Task.Variables.package}};

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;
{{""}}
{{- range .Imports}}
{{.}}
{{end -}}
{{""}}
/**
* {{.Table.Name}}
*
* @Author {{""}}
* @Date {{now}}
*/
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class {{titleCamelCase .Task.Table}} {
{{- range .Table.Columns}}
    /** {{.Comment}} */
    private {{dbToJava .Type}} {{camelCase .Name}}{{";" -}}
{{end}}
}`

const javaMapperTpl = `package {{.Task.Variables.package}};

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.Refs.JavaEntity.Variables.package}}.{{titleCamelCase .Table.Name}};

public interface {{titleCamelCase .Table.Name}}Mapper extends BaseMapper<{{titleCamelCase .Table.Name}}> {

}`

const javaMapperXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{.Refs.JavaMapper.Variables.package}}.{{titleCamelCase .Table.Name}}Mapper">
    <resultMap id="BaseResultMap" type="{{.Refs.JavaEntity.Variables.package}}.{{titleCamelCase .Table.Name}}">
    {{range .Table.Columns -}}
        {{"    "}}<id column="{{.Name}}" jdbcType="{{dbToJDBC .Type}}" property="{{camelCase .Name}}" />
    {{end -}}
    </resultMap>

    <sql id="Base_Column_List">
    {{"    " -}}
    {{range $i,$v := .Table.Columns -}}
        {{if ne $i 0}},{{end}}{{.Name -}}
    {{end}}
    </sql>
</mapper>`

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
	fmt.Println("Initialized successfully.")
}
