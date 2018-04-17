package main

import (
	"apiRestMongoDB/config"
	"apiRestMongoDB/dao"
	"apiRestMongoDB/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var conf = config.Config{}
var dataAccessO = dao.SportsDAO{}

//GetDeportesHandler Muestra todos los datos referentes a deportes en formato JSON
func GetDeportesHandler(writer http.ResponseWriter, request *http.Request) {
	deportes, errores := dataAccessO.FindAll()
	if errores != nil {

		respondWithError(writer, http.StatusInternalServerError, errores.Error())
		return
	}
	respondWithJson(writer, http.StatusOK, deportes)
}

//GetSportHandler Muestra los datos referentes al deporte con el id del parametro
func GetSportHandler(writer http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)
	deporte, errores := dataAccessO.FindById(parametros["id"])
	if errores != nil {
		respondWithError(writer, http.StatusBadRequest, "invalid movie ID")
		return
	}
	respondWithJson(writer, http.StatusOK, deporte)
}

//CreateDeporteHandler Crea un nuevo deporte con el id del parametro. Se deberia no pasar ningun parametro de id y generarlo automaticamente
func CreateDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var deporte model.Sport
	if err := json.NewDecoder(request.Body).Decode(&deporte); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Respusta invalida")
		return
	}
	deporte.ID = bson.NewObjectId()
	if err := dataAccessO.Insert(deporte); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(writer, http.StatusCreated, deporte)
}

//DeleteDeporteHandler Elimina el deporte con el id del parametro
func DeleteDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var deporte model.Sport
	err := json.NewDecoder(request.Body).Decode(&deporte)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = dataAccessO.Delete(deporte)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(writer, http.StatusOK, map[string]string{"resultado": "existoso"})
}

//UpdateDeporteHandler Actualiza los datos del id del parametro
func UpdateDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var deporte model.Sport
	err := json.NewDecoder(request.Body).Decode(&deporte)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = dataAccessO.Update(deporte)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(writer, http.StatusOK, map[string]string{"Resultado": "Existoso"})
}

func respondWithError(writer http.ResponseWriter, code int, msg string) {
	respondWithJson(writer, code, map[string]string{"error": msg})
}

//Respuesta en Json
func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	respuesta, _ := json.Marshal(payload)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(respuesta)
}

func init() {
	conf.Read()

	dataAccessO.Server = conf.Server
	dataAccessO.Database = conf.Database
	dataAccessO.Connect()
}
func main() {
	//Iniciamos el array de prueba

	router := mux.NewRouter() //Creamos un nuevo router
	router.HandleFunc("/deporte", GetDeportesHandler).Methods("GET")
	router.HandleFunc("/deporte/{id}", GetSportHandler).Methods("GET")
	router.HandleFunc("/deporte/", CreateDeporteHandler).Methods("POST")
	router.HandleFunc("/deporte/", DeleteDeporteHandler).Methods("DELETE")
	router.HandleFunc("/deporte/", UpdateDeporteHandler).Methods("PUT")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	} //Iniciamos el server y esperamos a la escucha en el puerto 3000

}
