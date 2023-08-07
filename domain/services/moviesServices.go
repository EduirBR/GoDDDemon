package services

import (
	"ddd2/app/extras"
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"encoding/json"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.RunSql(main_obj.DbSchema())
}

func CreateMovie(obj_json []byte) {
	query, values := main_obj.Sqline("insert", obj_json)
	repositories.InsertUpdate(query, values)
}

func UpdateMovie(obj_json []byte, id int) {
	query, values := main_obj.Sqline("update", obj_json, id)
	repositories.InsertUpdate(query, values)
}

func ListMovies() ([]entities.Movie, error) {
	query, _ := main_obj.Sqline("getAll")
	data, _ := repositories.GetAll(query) //Get Bytes
	var movies []entities.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	return movies, nil
}

func RetreiveMovie(id int) ([]entities.Movie, error) {
	query, _ := main_obj.Sqline("getBy", id)
	data, _ := repositories.GetAll(query) //Get Bytes
	var movies []entities.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	return movies, nil
}

func DeleteMovie(id int) error {
	query, _ := main_obj.Sqline("delete", id)
	err := repositories.RunSql(query) //Get Bytes
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
		return err
	}
	return nil
}
