package services

import (
	"ddd2/app/extras"
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"encoding/json"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.CreateTable(main_obj.DbSchema())
}

func CreateMovie(obj_json []byte) {
	repositories.Insert(main_obj.GetDbName(), main_obj, obj_json)
}

func ListMovies() ([]entities.Movie, error) {
	data, _ := repositories.GetAll(main_obj.GetDbName()) //Get Bytes
	var movies []entities.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	return movies, nil
}
