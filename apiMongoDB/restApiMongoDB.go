package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//GetDeportesHandler Muestra todos los datos referentes a deportes en formato JSON
func GetDeportesHandler(writer http.ResponseWriter, request *http.Request) {

}

//GetSportHandler Muestra los datos referentes al deporte con el id del parametro
func GetSportHandler(writer http.ResponseWriter, request *http.Request) {

}

//CreateDeporteHandler Crea un nuevo deporte con el id del parametro. Se deberia no pasar ningun parametro de id y generarlo automaticamente
func CreateDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	defer r.Body.Close()
	var deporte Sport
	if err := json.NewDecoder(r.Body).Decode(&deporte); err != nil {
		respondWithError(w, http.StatusBadRequest, "Respusta invalida")
		return
	}
	deporte.ID = bson.NewObjectId()
	if err := dao.Insert(deporte); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

//DeleteDeporteHandler Elimina el deporte con el id del parametro
func DeleteDeporteHandler(writer http.ResponseWriter, request *http.Request) {

}

//UpdateDeporteHandler Actualiza los datos del id del parametro
func UpdateDeporteHandler(writer http.ResponseWriter, request *http.Request) {

}
func main() {
	//Iniciamos el array de prueba

	router := mux.NewRouter() //Creamos un nuevo router
	router.HandleFunc("/deporte", GetDeportesHandler).Methods("GET")
	router.HandleFunc("/deporte/{id}", GetSportHandler).Methods("GET")
	router.HandleFunc("/deporte/", CreateDeporteHandler).Methods("POST")
	router.HandleFunc("/deporte/", DeleteDeporteHandler).Methods("DELETE")
	router.HandleFunc("/deporte/", UpdateDeporteHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router)) //Iniciamos el server y esperamos a la escucha en el puerto 3000

}
