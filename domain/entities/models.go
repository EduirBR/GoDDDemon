package entities

import (
	"ddd2/domain/extras"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// set the model name
type Models struct {
	// Atributes Atribute.Type then json tag
}

func (*Models) DbSchema() string {
	SQLScheema := `
	CREATE TABLE IF NOT EXISTS Models (

	);`
	return SQLScheema
}

// Funcion a añadir al modelo para optener el nombre de la tabla
func (m *Models) GetDbName() string {
	return "Models"
}

func (Mv *Models) Sqline(command, dbname string, obj interface{}, args ...any) (sql string, values []interface{}) {

	switch command {
	case "selectAll":
		sql = fmt.Sprintf("SELECT * FROM %s", dbname)
		values = nil
	case "selectById":
		_, id := getargs(args...)
		sql = fmt.Sprintf("SELECT * FROM %s WHERE id=%s", dbname, id)
	case "insert":
		jsons, _ := getargs(args...)
		vars, valuesAmount, val := getItems(jsons, obj)
		values = val
		sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", dbname, vars, valuesAmount)
	case "update":
		jsons, id := getargs(args...)
		vars, val := getUpdateSql(jsons, obj, id)
		values = val
		sql = fmt.Sprintf("UPDATE %s SET %s WHERE id=?", dbname, vars)
	case "delete":
		_, id := getargs(args...)
		sql = fmt.Sprintf("DELETE FROM %s WHERE id=%s", dbname, id)
	}
	return
}

func getItems(jsonD []byte, obj interface{}) (items, valuesAmount string, values []interface{}) {
	objType := reflect.TypeOf(obj)
	//Parse Json into map
	var data map[string]interface{}
	err := json.Unmarshal(jsonD, &data)
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	//build the queryset
	var columns []string
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		columnName := field.Tag.Get("json")
		columnValue, ok := data[columnName]
		if !ok {
			continue
		}
		columns = append(columns, columnName)
		values = append(values, columnValue)
	}

	valuesAmount = strings.Repeat("?, ", len(values)-1) + "?"
	items = strings.Join(columns, ", ")
	return
}

func getUpdateSql(jsonD []byte, obj interface{}, id string) (query string, values []interface{}) {

	objType := reflect.TypeOf(obj)
	//Parse Json into map
	var data map[string]interface{}
	err := json.Unmarshal(jsonD, &data)
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	var querySet []string
	//build the queryset
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		columnName := field.Tag.Get("json")
		columnValue, ok := data[columnName]
		if !ok {
			continue
		}
		sqline := fmt.Sprintf("%s=?", columnName)
		querySet = append(querySet, sqline)
		values = append(values, columnValue)

	}
	values = append(values, id)
	query = strings.Join(querySet, ", ")
	return
}

func getargs(args ...any) (jsonObj []byte, id string) {
	for _, arg := range args {
		if reflect.TypeOf(arg).Kind() == reflect.Slice && reflect.TypeOf(arg).Elem().Kind() == reflect.Uint8 {
			jsonObj = arg.([]byte)
		}
		if reflect.TypeOf(arg).Kind() == reflect.Int {
			id = strconv.Itoa(arg.(int))

		}
	}
	return
}
