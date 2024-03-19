package generator

import (
	"code-generator/transform"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/spf13/viper"
)

type DefaultGenerator struct {
	configMap map[string]any
	tasks     []*Task
	conn      *MysqlConnector
	refs      map[string]*Task
}

func (g *DefaultGenerator) Generate() string {
	defer func() {
		if g.conn != nil {
			g.conn.close()
		}
	}()

	for _, task := range g.tasks {
		tpl, err := pongo2.FromFile(".cg/templates/" + task.Template)
		if err != nil {
			panic(err)
		}

		if task.SourceType == "mysql" {
			var table *Table
			if g.conn == nil {
				g.conn = &MysqlConnector{DatabaseName: g.configMap["mysql.database"].(string), Username: g.configMap["mysql.username"].(string), Password: g.configMap["mysql.password"].(string), Host: g.configMap["mysql.host"].(string), Port: g.configMap["mysql.port"].(int)}
				g.conn.connect()
			}

			table = TableInfo(g.conn.db, task.Table)

			out, err := tpl.Execute(pongo2.Context{
				"task":        task,
				"table":       table,
				"transformer": &transform.Transformer{},
				"refs":        g.refs,
			})
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

	g.configMap = make(map[string]any)
	g.refs = make(map[string]*Task)
	g.tasks = []*Task{}
	for _, key := range keys {
		if key == "tasks" {
			tasks := viper.Get("tasks")
			for _, task := range tasks.([]interface{}) {
				taskDict := task.(map[string]interface{})
				structuredTask := &Task{
					Name:       taskDict["name"].(string),
					Template:   taskDict["template"].(string),
					SourceType: taskDict["source-type"].(string),
					Table:      taskDict["table"].(string),
					Output:     taskDict["output"].(string),
					Enable:     taskDict["enable"].(bool),
				}
				if taskDict["variables"] == nil {
					structuredTask.Variables = make(map[string]interface{})
				} else {
					structuredTask.Variables = taskDict["variables"].(map[string]interface{})
				}

				g.refs[structuredTask.Name] = structuredTask

				if structuredTask.Enable {
					g.tasks = append(g.tasks, structuredTask)
				}
			}
		} else {
			g.configMap[key] = viper.Get(key)
		}
	}

	return nil
}
