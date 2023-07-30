package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"log"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.CreateTable(main_obj.GetDbName(), main_obj)
}

func CreateMovie(obj_json []byte) {
	repositories.Insert(main_obj.GetDbName(), main_obj, obj_json)
}

func ListMovies() ([]entities.Movie, error) {
	data, _ := repositories.GetAll(main_obj.GetDbName()) //Estrae las SQLRows
	var movies []entities.Movie

	for _, d := range data {
		movie := entities.Movie{}
		id, _ := strconv.Atoi(d["id"].(string)) // comvierte a numero de str
		delete(d, "id")                         // elimina el elemento id del mapa para evitar problemas al parsear
		movie.ID = id
		err := mapstructure.Decode(d, &movie) // mapstructure adapta lo que este en el mapa a la estructura siempre
		//y cuando el tag de json sea el mismo nombre que en la base de datos.
		if err != nil {
			log.Println("Error ", err)
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
