package entities

type Movie struct {
	Models
	ID    int    `json:"id" pk:"true" extra:"auto_increment"`
	Code  int    `json:"code" unique:"true"`
	Isbn  string `json:"isbn" extra:"not null"`
	Title string `json:"title" unique:"true"`
	Good  bool   `json:"good"`
}

/*
pk:"true" para especificar la llave primaria
unique:"true" para unicos
extra: "todos los extras como 'not null' separados por un espacio"
*/
func (*Movie) DbSchema() string {
	SQLScheema := `
	CREATE TABLE IF NOT EXISTS Movies2 (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		code INT UNIQUE,
		isbn VARCHAR(255) NOT NULL,
		title VARCHAR(255) UNIQUE,
		good BOOL
	);`
	return SQLScheema
}

// Funcion a añadir al modelo para optener el nombre de la tabla
func (*Movie) GetDbName() string {
	return "Movies2"
}
