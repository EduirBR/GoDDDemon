package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"encoding/json"
	"log"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.CreateTable(main_obj.DbSchema())
}

func CreateMovie(obj_json []byte) {
	repositories.Insert(main_obj.GetDbName(), main_obj, obj_json)
}

func ListMovies() ([]entities.Movie, error) {
	data, _ := repositories.GetAll(main_obj.GetDbName(), main_obj) //Estrae las SQLRows
	var movies []entities.Movie
	jsonStr, errs := json.Marshal(data)
	if errs != nil {
		log.Println(errs)
	}
	json.Unmarshal(jsonStr, &movies)
	return movies, nil
}
