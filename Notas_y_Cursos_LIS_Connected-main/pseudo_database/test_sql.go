package pseudo_database

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// NO OLVIDES PONER ESTO EN VARIABLE DE ENTORNO
func conexionDB() (conexion *sql.DB) {
	conexion, err := sql.Open("mysql", ("clave_secreta"))
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

type Estructura_Cursos struct {
	Id            string
	Logo_URL      string
	Nombre        string
	Tema_Lenguaje string
	URL           string
	Idioma        string
	Precio        string
}

type Estructura_Cursos_Reducido struct {
	Id            string
	Logo_URL      string
	Nombre        string
	Tema_Lenguaje string
	//URL           string
	//Idioma        string
	//Precio        string

}

var Plantillas = template.Must(template.ParseGlob("templates/*"))

func Mostrar_Cursos(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionDB()
	registros, err := conexionEstablecida.Query("SELECT id, nombre, tema_Lenguaje, logo_URL FROM cursos_generales")
	if err != nil {
		panic(err.Error())
	}
	curso := Estructura_Cursos_Reducido{}
	arregloCursos := []Estructura_Cursos_Reducido{}

	for registros.Next() {
		var id string
		var nombre string
		var tema_Lenguaje string
		var logo_URL string

		err = registros.Scan(&id, &nombre, &tema_Lenguaje, &logo_URL)
		if err != nil {
			panic(err.Error())
		}

		curso.Id = id
		curso.Nombre = nombre
		curso.Tema_Lenguaje = tema_Lenguaje
		curso.Logo_URL = logo_URL

		arregloCursos = append(arregloCursos, curso)
	}
	fmt.Println(arregloCursos)

	Plantillas.ExecuteTemplate(w, "Plantilla_Global_Cursos", arregloCursos)
}

func Ver(w http.ResponseWriter, r *http.Request) {
	curso := Estructura_Cursos{}
	idCurso := r.URL.Query().Get("id")
	conexionEstablecida := conexionDB()
	registro, err := conexionEstablecida.Query("SELECT * FROM cursos_generales WHERE id=?", idCurso)
	if err != nil {
		panic(err.Error())
	}

	for registro.Next() {
		var id string
		var nombre string
		var uRL string
		var tema_Lenguaje string
		var idioma string
		var logo_URL string
		var precio string

		err = registro.Scan(&id, &nombre, &uRL, &tema_Lenguaje, &idioma, &logo_URL, &precio)
		if err != nil {
			panic(err.Error())
		}
		curso.Nombre = nombre
		curso.Tema_Lenguaje = tema_Lenguaje
		curso.Idioma = idioma
		curso.Logo_URL = logo_URL
		curso.URL = uRL
		curso.Precio = precio

	}
	Plantillas.ExecuteTemplate(w, "ver", curso)

}
