package generator

import (
	"code-generator/transform"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/spf13/viper"
)

type DefaultGenerator struct {
	project   *Project
	configMap map[string]any
	taskMap   []*Task
	conn      *MysqlConnector
}

func (g *DefaultGenerator) Generate() string {
	defer func() {
		if g.conn != nil {
			g.conn.close()
		}

	}()
	for _, task := range g.taskMap {
		tpl, err := pongo2.FromFile("./.cg/templates/" + task.Template)
		if err != nil {
			panic(err)
		}

		if task.Type == "table_to_entity" {
			var table *Table
			if g.conn == nil {
				g.conn = &MysqlConnector{DatabaseName: g.configMap["mysql.database"].(string), Username: g.configMap["mysql.username"].(string), Password: g.configMap["mysql.password"].(string), Host: g.configMap["mysql.host"].(string), Port: g.configMap["mysql.port"].(int)}
				g.conn.connect()
			}

			table = TableInfo(g.conn.db, task.Source)

			out, err := tpl.Execute(pongo2.Context{"project": g.project, "task": task, "table": table, "transformer": &transform.Transformer{}})
			if err != nil {
				panic(err)
			}
			FileWriter{}.Write(out, fmt.Sprintf("./.cg/output/%s", task.Output))
		}
	}
	return ""
}

func (g *DefaultGenerator) LoadConfig() error {
	keys := viper.AllKeys()
	if len(keys) == 0 {
		return fmt.Errorf("no keys found in config")
	}

	g.project = &Project{BasePackage: viper.GetString("project.base-package"), SrcPath: viper.GetString("project.src-path")}
	g.configMap = make(map[string]any)
	g.taskMap = []*Task{}
	for _, key := range keys {
		if key == "tasks" {
			tasks := viper.Get("tasks")
			for _, task := range tasks.([]interface{}) {
				taskDict := task.(map[string]interface{})
				structuredTask := &Task{Name: taskDict["name"].(string), Type: taskDict["type"].(string), SubPackage: taskDict["sub-package"].(string), Template: taskDict["template"].(string), Source: taskDict["source"].(string), Output: taskDict["output"].(string), Enable: taskDict["enable"].(bool)}
				if structuredTask.Enable {
					g.taskMap = append(g.taskMap, structuredTask)
				}
			}
		} else {
			g.configMap[key] = viper.Get(key)
		}
	}

	return nil
}
