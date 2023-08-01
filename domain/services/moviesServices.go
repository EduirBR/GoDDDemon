package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.CreateTable(main_obj.GetDbName(), main_obj)
}

func CreateMovie(obj_json []byte) {
	repositories.Insert(main_obj.GetDbName(), main_obj, obj_json)
}

func ListMovies() ([]entities.Movie, error) {
	data, _ := repositories.GetAll(main_obj.GetDbName(), main_obj) //Estrae las SQLRows
	var movies []entities.Movie

	for _, d := range data {
		movie := entities.Movie{}
		fmt.Println(d)
		//mainDecoder(movie, d)
		//err := mapstructure.Decode(d, &movie) // mapstructure adapta lo que este en el mapa a la estructura siempre
		/*y cuando el tag de json sea el mismo nombre que en la base de datos.
		if err != nil {
			log.Println("Error ", err)
			return nil, err
		}*/
		movies = append(movies, movie)
	}
	return movies, nil
}

func mainDecoder(obj interface{}, data map[string]interface{}) {
	objType := reflect.TypeOf(obj)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		columnName := field.Tag.Get("json")
		columnType := field.Type.Kind().String()
		h := data[columnName].([]uint8)
		var b interface{}
		//fmt.Println(string(h))
		json.Unmarshal(h, &b)
		//fmt.Println(b, " ", reflect.TypeOf(b))
		fmt.Println(columnName, " ", columnType, " ", b)

	}
	for _, vl := range data {
		var a any
		fmt.Println(vl.([]byte))
		err := json.Unmarshal(vl.([]byte), &a)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(a)
	}
}
