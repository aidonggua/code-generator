package initialize

import (
	"code-generator/generator"
	"fmt"
	"os"
)

type Initializer struct {
}

const configYaml = `project:
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
    enable: true                                        # enable or disable the task`

const entityTpl = `package {{project.BasePackage}}.{{task.SubPackage}};

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
{% endfor %}}`

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
	fmt.Println("Initialized successfully.")
}
