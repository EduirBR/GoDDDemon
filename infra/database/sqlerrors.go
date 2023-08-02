package database

import "log"

func DbError(err error) {
	log.Println("Error en Repositorios: ", err)
}
