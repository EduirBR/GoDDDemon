package entities

import (
	"ddd2/app/extras"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Movie struct {
	ID    int    `json:"id" pk:"true" extra:"auto_increment"`
	Code  int    `json:"code" unique:"true"`
	Isbn  string `json:"isbn" extra:"not null"`
	Title string `json:"title" unique:"true"`
	Good  bool   `json:"good"`
}

/*
pk:"true" para especificar la llave primaria
unique:"true" para unicos
extra: "todos los extras como 'not null' separados por un espacio"
*/
func (*Movie) DbSchema() string {
	SQLScheema := `
	CREATE TABLE IF NOT EXISTS Movies2 (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		code INT UNIQUE,
		isbn VARCHAR(255) NOT NULL,
		title VARCHAR(255) UNIQUE,
		good BOOL
	);`
	return SQLScheema
}

// Funcion a añadir al modelo para optener el nombre de la tabla
func (*Movie) GetDbName() string {
	return "Movies2"
}

func (Mv *Movie) Sqline(command string, args ...any) (sql string, values []interface{}) {

	switch command {
	case "getAll":
		sql = fmt.Sprintf("SELECT * FROM %s", Mv.GetDbName())
		values = nil
	case "getBy":
		_, id := getargs(args...)
		sql = fmt.Sprintf("SELECT * FROM %s WHERE id=%s", Mv.GetDbName(), id)
	case "insert":
		jsons, _ := getargs(args...)
		vars, valuesAmount, val := getItems(jsons)
		values = val
		sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", Mv.GetDbName(), vars, valuesAmount)
	case "update":
		jsons, id := getargs(args...)
		vars, val := getUpdateSql(jsons, id)
		values = val
		sql = fmt.Sprintf("UPDATE %s SET %s WHERE id=?", Mv.GetDbName(), vars)
	case "delete":
		_, id := getargs(args...)
		sql = fmt.Sprintf("DELETE FROM %s WHERE id=%s", Mv.GetDbName(), id)
	}
	return
}

func getItems(jsonD []byte) (items, valuesAmount string, values []interface{}) {
	movie := Movie{}
	objType := reflect.TypeOf(movie)
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

func getUpdateSql(jsonD []byte, id string) (query string, values []interface{}) {
	movie := Movie{}
	objType := reflect.TypeOf(movie)
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
