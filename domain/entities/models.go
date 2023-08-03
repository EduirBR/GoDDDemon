package entities

//set the model name
type Models struct {
	// Atributes Atribute.Type then json tag
}

func (*Models) DbSchema() string {
	SQLScheema := `
	CREATE TABLE IF NOT EXISTS Models (

	);`
	return SQLScheema
}

// Funcion a añadir al modelo para optener el nombre de la tabla
func (m *Models) GetDbName() string {
	return "Models"
}

func (*Models) Sqline(command string) (sql string) {
	//Set different SQLines what u need here
	return
}
