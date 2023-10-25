package estructuras_globales

type Datos_Basicos_curso struct {
	Name        string
	Description string
}

type Arreglo_cursos struct {
	Nombre      string
	Description string
	Link_Imagen string
	Link_URL    string
}

type Encapasulacion_para_templates_CURSOS struct {
	Personal_Info_Cursos  Datos_Basicos_curso
	Personal_Lista_Cursos []Arreglo_cursos
}
