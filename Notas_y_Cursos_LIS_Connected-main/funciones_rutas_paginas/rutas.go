package funciones_rutas_paginas

//https://stackoverflow.com/questions/38686583/golang-parse-all-templates-in-directory-and-subdirectories
import (
	estructuras_globales "lis_connected/estructuras_globales"
	pseudo_database "lis_connected/pseudo_database"
	"net/http"
	"text/template"
)

var Plantillas = template.Must(template.ParseGlob("templates/*"))

func Home(w http.ResponseWriter, r *http.Request) {

	Plantillas.ExecuteTemplate(w, "Main", nil)
}

func RedirectToHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Notas_Java(w http.ResponseWriter, r *http.Request) {

	data := estructuras_globales.Encapasulacion_para_templates_NOTAS{
		Personal_Info_Notas:  pseudo_database.Info_Notas_Java,
		Personal_Lista_Notas: pseudo_database.Lista_Notas_Java,
	}

	Plantillas.ExecuteTemplate(w, "Notas_Plantilla_Generica", data)

}

func Cursos_Java(w http.ResponseWriter, r *http.Request) {

	data := estructuras_globales.Encapasulacion_para_templates_CURSOS{
		Personal_Info_Cursos:  pseudo_database.Info_Cursos_Java,
		Personal_Lista_Cursos: pseudo_database.Lista_Cursos_Java,
	}

	Plantillas.ExecuteTemplate(w, "Cursos_Plantilla_Generica", data)
}
