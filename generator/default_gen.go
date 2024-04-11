package generator

import (
	"bytes"
	"code-generator/helper"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
	"text/template"
)

type DefaultGenerator struct {
	configMap   map[string]string
	tasks       []*Task
	conn        *MysqlConnector
	refs        map[string]*Task
	currentTask *Task
}

func (g *DefaultGenerator) Generate() string {
	g.conn = &MysqlConnector{
		DatabaseName: g.configMap["mysql.database"],
		Username:     g.configMap["mysql.username"],
		Password:     g.configMap["mysql.password"],
		Host:         g.configMap["mysql.host"],
		Port:         g.configMap["mysql.port"],
	}
	g.conn.connect()

	table := TableInfo(g.conn.db, g.configMap["mysql.database"], g.configMap["mysql.table"])
	defer func() {
		if g.conn != nil {
			g.conn.close()
		}
	}()

	h := helper.Helper{}
	funcMap := template.FuncMap{
		"camelCase": h.CamelCase,
		"snakeCase": h.SnakeCase,
		"kebabCase": h.KebabCase,
		"title":     strings.Title,
		"upperCase": strings.ToUpper,
		"lowerCase": strings.ToLower,
		"dbToJava":  h.DbToJava,
		"dbToJDBC":  h.DbToJDBC,
		"dbToGo":    h.DbToGo,
		"now":       h.Now,
		"date":      h.Date,
		"time":      h.Time,
		"config": func(key string) string {
			return g.configMap[key]
		},
		"string": func(i any) string {
			switch v := i.(type) {
			case string:
				return v
			case int:
				return strconv.Itoa(v)
			default:
				return ""
			}
		},
	}

	for _, t := range g.tasks {
		g.currentTask = t

		var imports []string
		for _, v := range table.Columns {
			importStr := h.DbToJavaImport(v.Type)
			if importStr != "" && !slices.Contains(imports, importStr) {
				imports = append(imports, importStr)
			}
		}

		funcMap["prop"] = func(key string) string {
			return t.Properties[key].(string)
		}
		refsFunc := func(task, key string) string {
			return g.refs[task].Properties[key].(string)
		}
		funcMap["refs"] = refsFunc
		classNameFunc := func(task string) string {
			if task == "." {
				task = t.Name
			}
			return strings.Title(fmt.Sprintf("%s%s%s", g.refs[task].Prefix, h.CamelCase(table.Name), g.refs[task].Postfix))
		}
		funcMap["className"] = classNameFunc
		packageFunc := func(task string) string {
			if task == "." {
				task = t.Name
			}
			return fmt.Sprintf("%s.%s.%s", g.configMap["base-package"], g.configMap["module"], refsFunc(task, "sub-package"))
		}
		funcMap["package"] = packageFunc
		fullClassNameFunc := func(task string) string {
			if task == "." {
				task = t.Name
			}
			return fmt.Sprintf("%s.%s", packageFunc(task), classNameFunc(task))
		}
		funcMap["fullClassName"] = fullClassNameFunc

		funcMap["imports"] = func() []string {
			return imports
		}
		tableMap := make(map[string]string)
		tableMap["name"] = table.Name
		tableMap["comment"] = table.Comment
		funcMap["table"] = func(key string) string {
			return tableMap[key]
		}
		var columnsMap []map[string]string
		for _, v := range table.Columns {
			columnMap := make(map[string]string)
			columnMap["name"] = v.Name
			columnMap["type"] = v.Type
			columnMap["collation"] = v.Collation.String
			columnMap["null"] = v.Null
			columnMap["key"] = v.Key
			columnMap["default"] = v.Default.String
			columnMap["extra"] = v.Extra
			columnMap["privileges"] = v.Privileges
			columnMap["comment"] = v.Comment
			columnsMap = append(columnsMap, columnMap)
		}
		funcMap["columns"] = func() []map[string]string {
			return columnsMap
		}

		tpl, err := template.New(t.Template).Funcs(funcMap).ParseFiles(".cg/templates/" + t.Template)
		var buffer bytes.Buffer
		err = tpl.Execute(&buffer, nil)
		if err != nil {
			panic(err)
		}

		FileWriter{}.Write(buffer.String(), fmt.Sprintf(".cg/out/%s%s%s%s", t.Prefix, strings.Title(h.CamelCase(table.Name)), t.Postfix, t.FileType))
	}
	return ""
}

func (g *DefaultGenerator) LoadConfig() error {
	keys := viper.AllKeys()
	if len(keys) == 0 {
		return fmt.Errorf("no keys found in config")
	}

	g.configMap = make(map[string]string)
	g.refs = make(map[string]*Task)
	g.tasks = []*Task{}
	for _, key := range keys {
		if key == "tasks" {
			tasksConf := viper.Get("tasks")
			for _, taskConf := range tasksConf.([]interface{}) {
				taskConfMap := taskConf.(map[string]interface{})
				structuredTask := &Task{
					Name:     taskConfMap["name"].(string),
					Template: taskConfMap["template"].(string),
					FileType: taskConfMap["file-type"].(string),
					Prefix:   taskConfMap["prefix"].(string),
					Postfix:  taskConfMap["postfix"].(string),
					Enable:   taskConfMap["enable"].(bool),
				}
				if taskConfMap["properties"] == nil {
					structuredTask.Properties = make(map[string]interface{})
				} else {
					structuredTask.Properties = taskConfMap["properties"].(map[string]interface{})
				}
				structuredTask.Properties["name"] = structuredTask.Name
				structuredTask.Properties["template"] = structuredTask.Template
				structuredTask.Properties["file-type"] = structuredTask.FileType
				structuredTask.Properties["prefix"] = structuredTask.Prefix
				structuredTask.Properties["postfix"] = structuredTask.Postfix
				structuredTask.Properties["enable"] = structuredTask.Enable

				g.refs[structuredTask.Name] = structuredTask

				if structuredTask.Enable {
					g.tasks = append(g.tasks, structuredTask)
				}
			}
		} else {
			g.configMap[key] = viper.GetString(key)
		}
	}

	return nil
}

func (g *DefaultGenerator) CurrentTask() *Task {
	return g.currentTask
}
