package generator

import (
	"bytes"
	"code-generator/helper"
	"code-generator/task"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"strings"
	"text/template"
)

type DefaultGenerator struct {
	configMap   map[string]any
	tasks       []*task.Task
	conn        *MysqlConnector
	refs        map[string]*task.Task
	currentTask *task.Task
}

func (g *DefaultGenerator) Generate() string {
	defer func() {
		if g.conn != nil {
			g.conn.close()
		}
	}()

	h := helper.Helper{}
	funcMap := template.FuncMap{
		"titleCamelCase": h.TitleCamelCase,
		"camelCase":      h.CamelCase,
		"snakeCase":      h.SnakeCase,
		"title":          strings.Title,
		"upperCase":      strings.ToUpper,
		"lowerCase":      strings.ToLower,
		"dbToJava":       h.DbToJava,
		"dbToJDBC":       h.DbToJDBC,
		"dbToGo":         h.DbToGo,
		"now":            h.Now,
		"date":           h.Date,
		"time":           h.Time,
	}

	for _, t := range g.tasks {
		g.currentTask = t

		var table *Table
		if t.SourceType == "mysql" {
			if g.conn == nil {
				g.conn = &MysqlConnector{
					DatabaseName: g.configMap["mysql.database"].(string),
					Username:     g.configMap["mysql.username"].(string),
					Password:     g.configMap["mysql.password"].(string),
					Host:         g.configMap["mysql.host"].(string),
					Port:         g.configMap["mysql.port"].(int),
				}
				g.conn.connect()
			}

			table = TableInfo(g.conn.db, t.Table)
		}

		var imports []string
		for _, v := range table.Columns {
			importStr := h.DbToJavaImport(v.Type)
			if importStr != "" && !slices.Contains(imports, importStr) {
				imports = append(imports, importStr)
			}
		}

		tpl, err := template.New(t.Template).Funcs(funcMap).ParseFiles(".cg/templates/" + t.Template)
		var buffer bytes.Buffer
		data := struct {
			Task    *task.Task
			Table   *Table
			Refs    map[string]*task.Task
			Imports []string
		}{
			Task:    t,
			Table:   table,
			Refs:    g.refs,
			Imports: imports,
		}
		err = tpl.Execute(&buffer, &data)
		if err != nil {
			panic(err)
		}

		FileWriter{}.Write(buffer.String(), fmt.Sprintf("./.cg/output/%s", t.Output))
	}
	return ""
}

func (g *DefaultGenerator) LoadConfig() error {
	keys := viper.AllKeys()
	if len(keys) == 0 {
		return fmt.Errorf("no keys found in config")
	}

	g.configMap = make(map[string]any)
	g.refs = make(map[string]*task.Task)
	g.tasks = []*task.Task{}
	for _, key := range keys {
		if key == "tasks" {
			tasksConf := viper.Get("tasks")
			for _, taskConf := range tasksConf.([]interface{}) {
				taskConfMap := taskConf.(map[string]interface{})
				structuredTask := &task.Task{
					Name:       taskConfMap["name"].(string),
					Template:   taskConfMap["template"].(string),
					SourceType: taskConfMap["source-type"].(string),
					Table:      taskConfMap["table"].(string),
					Output:     taskConfMap["output"].(string),
					Enable:     taskConfMap["enable"].(bool),
					Imports:    make([]string, 0),
				}
				if taskConfMap["variables"] == nil {
					structuredTask.Variables = make(map[string]interface{})
				} else {
					structuredTask.Variables = taskConfMap["variables"].(map[string]interface{})
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

func (g *DefaultGenerator) CurrentTask() *task.Task {
	return g.currentTask
}
