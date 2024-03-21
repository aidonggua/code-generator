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
  - name: entity                                        # task name
    template: entity.tpl                                # template file from .cg/templates folder
    source-type: mysql                                  # table to entity
    table: user                                         # table name
    output: User.java                                   # output file name
    variables:                                          # variables for template
      package: com.example.dao.domain
    enable: true                                        # enable or disable the task
  - name: mapper
    template: mapper.tpl
    source-type: mysql
    table: user
    output: UserMapper.java
    variables:
      package: com.example.dao.mapper
    enable: true
  - name: mapper_xml
    template: mapper.xml.tpl
    source-type: mysql
    table: user
    output: UserMapper.xml
    enable: true
`

const entityTpl = `package {{.Task.Variables.package}};

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

const mapperTpl = `package {{.Task.Variables.package}};

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.Refs.entity.Variables.package}}.{{titleCamelCase .Table.Name}};

public interface {{titleCamelCase .Table.Name}}Mapper extends BaseMapper<{{titleCamelCase .Table.Name}}> {

}`

const mapperXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
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

func (i *Initializer) Init() {
	// check if .cg folder exists
	if _, err := os.Stat(".cg"); err == nil {
		fmt.Print("The .cg folder already exists.")
		return
	}

	generator.FileWriter{}.CreateFolder(".cg/output")
	generator.FileWriter{}.CreateFolder(".cg/templates")
	generator.FileWriter{}.Write(configYaml, ".cg/config.yaml")
	generator.FileWriter{}.Write(entityTpl, ".cg/templates/entity.tpl")
	generator.FileWriter{}.Write(mapperTpl, ".cg/templates/mapper.tpl")
	generator.FileWriter{}.Write(mapperXmlTpl, ".cg/templates/mapper.xml.tpl")
	fmt.Println("Initialized successfully.")
}
