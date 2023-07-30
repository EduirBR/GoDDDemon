package services

import (
	"ddd2/domain/entities"
	"ddd2/infra/repositories"
	"log"
)

var main_obj = entities.Movie{}

func CreateTableMovies() {
	repositories.CreateTable(main_obj.GetDbName(), main_obj)
}

func CreateMovie(obj_json []byte) {
	repositories.Insert(main_obj.GetDbName(), main_obj, obj_json)
}

func GetMovies() ([]entities.Movie, error) {
	rows, _ := repositories.GetAll(main_obj.GetDbName()) //Estrae las SQLRows
	var movies []entities.Movie

	for rows.Next() {
		movie := entities.Movie{}
		err := rows.Scan(&movie.ID, &movie.Isbn, &movie.Title)
		if err != nil {
			log.Println("Error ", err)
			return nil, err
		}
		movies = append(movies, movie)
	}
	// jsonData, err := json.Marshal(movies)
	// if err != nil {
	// 	return nil, err
	// }
	return movies, nil
}
