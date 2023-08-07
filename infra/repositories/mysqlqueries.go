package repositories

import (
	"database/sql"
	"ddd2/app/extras"
	"ddd2/infra/database"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// db *sql.DB,
func RunSql(query string) error {

	db := database.DbConnect()
	defer db.Close()
	// Ejecutar la consulta SQL
	_, err := db.Exec(query)
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	return nil
}

func InsertUpdate(query string, values []interface{}) error {

	//db conect
	db := database.DbConnect()
	defer db.Close()
	//sql request
	stmt, err := db.Prepare(query)
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(values...)
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
		return err
	}
	return nil
}

func GetAll(query string) ([]byte, error) {

	//db conext
	db := database.DbConnect()
	defer db.Close()
	//Query
	rows, err := db.Query(query)
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	//checking errors
	if err = rows.Err(); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	data := getBytes(rows)
	return data, nil
}

func getBytes(rows *sql.Rows) []byte {
	columns, err := rows.Columns() //Get The Colums to know the number of them
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	data := []byte{}
	startkey := []byte("[")
	data = append(data, startkey...)
	newlap := rows.Next() // var made to control the loops
	for newlap {
		start := []byte("{")
		data = append(data, start...)
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns { // para la cantidad de columnas
			valuePtrs[i] = &values[i] //make an pointer for values at i
		}
		err := rows.Scan(valuePtrs...) //set the items in Values
		if err != nil {
			extras.Errors(extras.GetFunctionName(), err)
		}

		for i, col := range columns {
			val := values[i] // set in val the gotten from rows
			colstring := fmt.Sprintf("\"%s\":", col)
			byteCol := []byte(colstring) //setting the key as []byte
			//fmt.Println(col, " ", val)
			b, ok := val.([]byte) //setting the value to []byte
			if ok {
				data = append(data, byteCol...)
				//controling the varchar
				if strings.ToLower(columnTypes[i].DatabaseTypeName()) == "varchar" {
					p1 := []byte("\"")
					p2 := []byte("\"")
					data = append(data, p1...)
					data = append(data, b...)
					data = append(data, p2...)
				} else if strings.ToLower(columnTypes[i].DatabaseTypeName()) == "tinyint" {
					//controlling bools
					if string(b) == "1" {
						booldata, _ := json.Marshal(true)
						data = append(data, booldata...)
					} else {
						booldata, _ := json.Marshal(false)
						data = append(data, booldata...)
					}
				} else {

					data = append(data, b...)
				}

			} else {
				data = append(data, byteCol...)
				if strings.ToLower(columnTypes[i].DatabaseTypeName()) == "tinyint" {
					booldata, _ := json.Marshal(false)
					data = append(data, booldata...)
				} else {
					unit8Data := []byte("0")
					data = append(data, unit8Data...)

				}
			}
			//if to avoid to set the comma at last item
			if i < len(columns)-1 {
				comma := []byte(", ")
				data = append(data, comma...)
			}
		}
		endkey := []byte("}")
		data = append(data, endkey...)
		//controller to dont set the comma at last lap
		newlap = rows.Next()
		if newlap {
			comma := []byte(",")
			data = append(data, comma...)
		}
	}
	end := []byte("]")
	data = append(data, end...)
	return data
}

func ObjecType(obj interface{}) map[int]interface{} {
	//get the object type
	objType := reflect.TypeOf(obj)
	//build a map with object type num of fields
	ValuesFields := make(map[int]interface{}, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		//get object field on i
		field := objType.Field(i)
		//get the type
		columnType := field.Type
		//make a default value with that column type
		zerovalue := reflect.Zero(columnType).Interface()
		//set the value on the map with key i
		ValuesFields[i] = zerovalue
	}
	return ValuesFields
}
