package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.CreateTable(main_obj.GetDbName(), main_obj)
}

func CreateMovie(obj_json []byte) {
	repositories.Insert(main_obj.GetDbName(), main_obj, obj_json)
}
