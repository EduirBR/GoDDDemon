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

	fmt.Println(query)
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

func GetAll(tableName string) ([]map[string]interface{}, error) {
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

	columns, err := rows.Columns() //extrae las columnas de las filas dadas por la base de datos en un array de strings
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns { // para la cantidad de columnas
			valuePtrs[i] = &values[i] //crea una cantidad de variables igual a las columnas que hay
		}
		err := rows.Scan(valuePtrs...) // guarda los valores de las filas en cada una de las variables creadas
		if err != nil {
			log.Println("Error: ", err)
			return nil, err
		}
		rowData := make(map[string]interface{}) //fila de datos
		for i, col := range columns {           //itera la cantidad de columnas
			val := values[i]      // para guardar los valores de la base de datos en var
			b, ok := val.([]byte) //verifica si estan en bytes
			if ok {
				rowData[col] = string(b) //si? los convierte en string
			} else {
				rowData[col] = val //no? los guarda como estan
			}
		}
		data = append(data, rowData) // guarda en un arreglo los mapas
	}

	return data, nil
}
