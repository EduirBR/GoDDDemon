package repositories

import (
	"ddd2/infra/database"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// db *sql.DB,
func CreateTable(tableName string, obj interface{}) error {

	// Obtener el tipo de la estructura
	objType := reflect.TypeOf(obj)

	// Crear la consulta SQL
	var columns []string
	var unique []string
	primaryKey := ""
	uniqueSql := ""
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		columnName := field.Tag.Get("json")
		columnType := getColumnType(field.Type)
		extras := field.Tag.Get("extra")
		columns = append(columns, fmt.Sprintf("%s %s %s", columnName, columnType, extras))

		if field.Tag.Get("pk") == "true" {
			if primaryKey == "" {
				primaryKey = fmt.Sprintf(", PRIMARY KEY (%s)", columnName)
			} else {
				fmt.Println("Error solo puede haber una llave primaria")
			}
		}
		if field.Tag.Get("unique") == "true" {
			unique = append(unique, columnName)
		}
		if len(unique) > 0 {
			uniqueSql = fmt.Sprintf(", UNIQUE (%s)", strings.Join(unique, ", "))
		}
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s%s%s);", tableName, strings.Join(columns, ", "), primaryKey, uniqueSql)

	//conexion con la base de datos

	db := database.DbConnect()
	defer db.Close()
	// Ejecutar la consulta SQL
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error", err)
		return err
	}
	return nil
}

func getColumnType(field reflect.Type) string {
	switch field.Kind() {
	case reflect.String:
		return "varchar(255)"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Float32, reflect.Float64:
		return "float"
	default:
		return "varchar(255)"
	}
}

func Insert(tableName string, obj interface{}, jsonD []byte) error {
	// Obtener el tipo de la estructura
	objType := reflect.TypeOf(obj)

	// Convertir el objeto JSON a un mapa
	var data map[string]interface{}
	err := json.Unmarshal(jsonD, &data)
	if err != nil {
		log.Println("Error", err)
		return err
	}

	// Crear la consulta SQL
	var columns []string
	var values []interface{}
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

	n_values := strings.Repeat("?, ", len(values)-1) + "?"
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), n_values)

	//conexion con la base de datos

	db := database.DbConnect()
	defer db.Close()
	// Ejecutar la consulta SQL
	_, err = db.Exec(query, values...)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	return nil
}

func GetAll(tableName string, dest interface{}) ([]map[string]interface{}, error) {

	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	//conexion con la base de datos
	db := database.DbConnect()
	defer db.Close()
	//consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error", err)
		return nil, err
	}
	//revision de errores
	if err = rows.Err(); err != nil {
		log.Println("Error", err)
		return nil, err
	}
	var data []map[string]interface{} // crea un objeto mapa con cuyos valores seran strings

	for rows.Next() {

		destType := reflect.TypeOf(dest)

		// Crear un slice de interfaces para almacenar los valores escaneados
		values := make([]interface{}, destType.NumField())

		// Asignar cada elemento del slice de interfaces a un campo de la interfaz de destino
		for i := 0; i < destType.NumField(); i++ {
			values[i] = reflect.New(destType.Field(i).Type).Interface()
		}

		// Escanear los valores de la fila en el slice de interfaces
		err = rows.Scan(values...)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(values)
		// Asignar los valores escaneados a los campos de la interfaz de destino
		for i := 0; i < destType.NumField(); i++ {
			reflect.ValueOf(dest).Field(i).Set(reflect.ValueOf(values[i]))
		}
		fmt.Println(destType)
	}
	return data, nil
}
