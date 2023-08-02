package repositories

import (
	"database/sql"
	"ddd2/infra/database"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// db *sql.DB,
func CreateTable(query string) error {

	db := database.DbConnect()
	defer db.Close()
	// Ejecutar la consulta SQL
	a, err := db.Exec(query)
	if err != nil {
		database.DbError(err)
	}
	fmt.Println(a)
	return nil
}

func Insert(tableName string, obj interface{}, jsonD []byte) error {
	// Obtener el tipo de la estructura
	objType := reflect.TypeOf(obj)

	// Convertir el objeto JSON a un mapa
	var data map[string]interface{}
	err := json.Unmarshal(jsonD, &data)
	if err != nil {
		database.DbError(err)
	}
	fmt.Println(data)
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
		database.DbError(err)
	}

	return nil
}

func GetAll(tableName string, obj interface{}) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	objFields := ObjecType(obj)
	//conexion con la base de datos
	db := database.DbConnect()
	defer db.Close()
	//consulta
	rows, err := db.Query(query)
	if err != nil {
		database.DbError(err)
	}
	//revision de errores
	if err = rows.Err(); err != nil {
		database.DbError(err)
	}
	data := fixSql(rows, objFields)
	return data, nil
}

func fixSql(rows *sql.Rows, objFields []map[string]interface{}) []map[string]interface{} {
	columns, err := rows.Columns() //extrae las columnas de las filas dadas por la base de datos en un array de strings
	if err != nil {
		database.DbError(err)
	}
	// ty, _ := rows.ColumnTypes()
	// for _, u := range ty {
	// 	fmt.Println(u.ScanType())
	// 	fmt.Println(u.DatabaseTypeName())
	// }
	var data []map[string]interface{} // crea un objeto mapa con cuyos valores seran strings
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns { // para la cantidad de columnas
			valuePtrs[i] = &values[i] //crea una cantidad de variables igual a las columnas que hay
		}
		err := rows.Scan(valuePtrs...) // guarda los valores de las filas en cada una de las variables creadas
		if err != nil {
			database.DbError(err)
		}
		rowData := make(map[string]interface{}) //fila de datos
		for i, col := range columns {           //itera la cantidad de columnas
			types := objFields[i][col]
			val := values[i]      // para guardar los valores de la base de datos en var
			b, ok := val.([]byte) //verifica si estan en bytes
			if ok {
				rowData[col] = VarType(b, types) //si? los convierte en string
			} else {
				rowData[col] = val //no? los guarda como estan
			}
		}
		data = append(data, rowData) // guarda en un arreglo los mapas
	}
	return data
}

func VarType(value []uint8, ty interface{}) any {
	// fmt.Println(value, " ", ty)
	tyR := reflect.TypeOf(ty)
	tya := reflect.ValueOf(ty)
	fmt.Println(reflect.TypeOf(tya.Type()))
	fmt.Println("----------------")
	fmt.Println(tya.Interface())
	fmt.Println("----------------")
	hR := reflect.New(tyR)
	a := reflect.ValueOf(value)
	hR.Elem().Set(a.Convert(tya.Type()))
	h := hR.Elem().Interface()
	fmt.Println(h)

	return "hola"
}
func ObjecType(obj interface{}) (objFields []map[string]interface{}) {
	objType := reflect.TypeOf(obj)
	//var campos []map[string]interface{}
	// Crear la consulta SQL
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		columnName := field.Tag.Get("json")
		columnType := field.Type.Kind()
		item := map[string]any{
			columnName: columnType,
		}
		objFields = append(objFields, item)
	}
	return
}
