package estructuras_globales

type Datos_Basicos_notas struct {
	Name           string
	Description    string
	Autor_Original string
}

type Arreglo_notas struct {
	Titulo    string
	PseudoId  string
	Link_Gist string
}

type Encapasulacion_para_templates_NOTAS struct {
	Personal_Info_Notas  Datos_Basicos_notas
	Personal_Lista_Notas []Arreglo_notas
}
