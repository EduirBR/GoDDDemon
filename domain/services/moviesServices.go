package services

import (
	"ddd2/domain/entities"
	"ddd2/domain/extras"
	"ddd2/infra/repositories"
	"encoding/json"
)

var main_obj = entities.Movie{}
var movide_name = main_obj.GetDbName()

func CreateTableMovies() {
	repositories.RunSql(main_obj.DbSchema())
}

func CreateMovie(obj_json []byte) error {
	query, values := main_obj.Sqline("insert", movide_name, obj_json)
	err := repositories.InsertUpdate(query, values)
	if err != nil {
		return err
	}
	return nil
}

func EditMovie(obj_json []byte, id int) error {
	query, values := main_obj.Sqline("update", movide_name, obj_json, id)
	err := repositories.InsertUpdate(query, values)
	if err != nil {
		return err
	}
	return nil
}

func ListMovies() ([]entities.Movie, error) {
	query, _ := main_obj.Sqline("selectAll", movide_name, nil)
	data, _ := repositories.Select(query) //Get Bytes
	var movies []entities.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
		return nil, err
	}
	return movies, nil
}

func RetreiveMovie(id int) ([]entities.Movie, error) {
	query, _ := main_obj.Sqline("selectById", movide_name, nil, id)
	data, _ := repositories.Select(query) //Get Bytes
	var movies []entities.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
		return nil, err
	}
	return movies, nil
}

func DeleteMovie(id int) error {
	query, _ := main_obj.Sqline("delete", movide_name, nil, id)
	err := repositories.RunSql(query) //Get Bytes
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
		return err
	}
	return nil
}
