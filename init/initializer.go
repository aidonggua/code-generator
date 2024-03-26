package initialize

import (
	"code-generator/generator"
	"fmt"
	"os"
)

type Initializer struct {
}

const configYaml = `# config file of code generator
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
`

const javaEntityTpl = `package {{.Task.Variables.package}};

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

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
import {{.Refs.entity.Variables.package}}.{{titleCamelCase .Table.Name}};

public interface {{titleCamelCase .Table.Name}}Mapper extends BaseMapper<{{titleCamelCase .Table.Name}}> {

}`

const javaMapperXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{.Refs.mapper.Variables.package}}.{{titleCamelCase .Table.Name}}Mapper">
    <resultMap id="BaseResultMap" type="{{.Refs.entity.Variables.package}}.{{titleCamelCase .Table.Name}}">
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

const goEntityTpl = `package {{.Task.Variables.package}};

type {{title .Table.Name}} struct {
{{- range .Table.Columns}}
    {{titleCamelCase .Name}} {{dbToGo .Type}}` + " `json" + "\"{{.Name}}\"` " + `// {{.Comment -}}
{{end}}
}

func (*{{title .Table.Name}}) TableName() string {
	return "{{.Table.Name}}"
}`

func (i *Initializer) Init() {
	// check if .cg folder exists
	if _, err := os.Stat(".cg"); err == nil {
		fmt.Print("The .cg folder already exists.")
		return
	}

	generator.FileWriter{}.CreateFolder(".cg/output")
	generator.FileWriter{}.CreateFolder(".cg/templates")
	generator.FileWriter{}.Write(configYaml, ".cg/config.yaml")
	generator.FileWriter{}.Write(javaEntityTpl, ".cg/templates/java/entity.tpl")
	generator.FileWriter{}.Write(javaMapperTpl, ".cg/templates/java/mapper.tpl")
	generator.FileWriter{}.Write(javaMapperXmlTpl, ".cg/templates/java/mapper.xml.tpl")
	generator.FileWriter{}.Write(goEntityTpl, ".cg/templates/go/entity.tpl")
	fmt.Println("Initialized successfully.")
}
