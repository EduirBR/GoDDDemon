package entities

type Movie struct {
	ID    int    `json:"id" pk:"true" extra:"auto_increment"`
	Code  int    `json:"code" unique:"true"`
	Isbn  string `json:"isbn" extra:"not null"`
	Title string `json:"title" unique:"true"`
}

/*
pk:"true" para especificar la llave primaria
unique:"true" para unicos
extra: "todos los extras como 'not null' separados por un espacio"
*/

// Funcion a añadir al modelo para optener el nombre de la tabla
func (*Movie) GetDbName() string {
	return "Movies2"
}
