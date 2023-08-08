package services

import (
	"ddd2/domain/entities"
	"ddd2/domain/extras"
	"ddd2/infra/repositories"
	"encoding/json"
)

var model = entities.Models{}

func CreateTableModels() {
	repositories.RunSql(model.DbSchema())
}

func CreateModel(obj_json []byte) {
	//repositories.Insert(model.GetDbName(), model, obj_json)
}

func ListModels() ([]entities.Movie, error) {
	data, _ := repositories.Select(model.GetDbName()) //Get Bytes
	var Models []entities.Movie
	if err := json.Unmarshal(data, &Models); err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	return Models, nil
}
