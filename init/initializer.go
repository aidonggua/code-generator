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

const entityTpl = `package {{ task.Variables.package }};

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
{% endfor %}}`

const mapperTpl = `package {{ task.Variables.package }};

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{ refs.entity.Variables.package }}.{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }};

public interface {{ transformer.Case.Title(transformer.Case.CamelCase(task.Table)) }}Mapper extends BaseMapper<{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }}>{
}`

const mapperXmlTpl = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{ refs.mapper.Variables.package }}.{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }}Mapper">
  <resultMap id="BaseResultMap" type="{{ refs.entity.Variables.package }}.{{ transformer.Case.Title(transformer.Case.CamelCase(refs.entity.Table)) }}">
{% for column in table.Columns %}    <id column="{{ column.Name }}" jdbcType="{{ transformer.Type.DbToJDBC(column.Type) }}" property="{{ transformer.Case.CamelCase(column.Name) }}" />
{% endfor %}  </resultMap>
  <sql id="Base_Column_List">
    {% for column in table.Columns %}{{ column.Name }},{% endfor %}
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
