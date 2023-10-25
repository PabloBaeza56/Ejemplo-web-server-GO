package main

//https://codea.app/blog/centrar-div-boostrap

import (
	"lis_connected/funciones_rutas_paginas"
	pseudo_database "lis_connected/pseudo_database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", funciones_rutas_paginas.Home)
	router.HandleFunc("/notas_java", funciones_rutas_paginas.Notas_Java)
	router.HandleFunc("/cursos_java", funciones_rutas_paginas.Cursos_Java)

	/// TEST AREA
	router.HandleFunc("/todos_cursos", pseudo_database.Mostrar_Cursos)
	router.HandleFunc("/ver", pseudo_database.Ver)

	router.NotFoundHandler = http.HandlerFunc(funciones_rutas_paginas.RedirectToHomePage)

	log.Println("Servidor corriendo ...")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":3002", router))
}
