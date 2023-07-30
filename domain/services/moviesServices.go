package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"fmt"
	"reflect"
)

func CreateTableMovies() {
	obj := entities.Movie{}
	fmt.Println(reflect.TypeOf(obj))

	repositories.CreateTable(obj.GetDbName(), obj)
}
