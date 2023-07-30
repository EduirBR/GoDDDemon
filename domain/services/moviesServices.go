package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
)

func CreateTableMovies() {
	obj := entities.Movie{}
	repositories.CreateTable(obj.GetDbName(), obj)
}
