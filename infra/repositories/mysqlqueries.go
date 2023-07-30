package repositories

import (
	"ddd2/infra/database"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// db *sql.DB,
func CreateTable(tableName string, sample interface{}) error {
	//conexion con la base de datos

	db := database.DbConnect()
	defer db.Close()
	// Obtener el tipo de la estructura
	sampleType := reflect.TypeOf(sample)

	// Crear la consulta SQL
	var columns []string
	var unique []string
	primaryKey := ""
	uniqueSql := ""
	for i := 0; i < sampleType.NumField(); i++ {
		field := sampleType.Field(i)
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

	fmt.Println(query)
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
